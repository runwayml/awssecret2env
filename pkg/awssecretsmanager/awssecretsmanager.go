package awssecretsmanager

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"

	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/runwayml/awssecret2env/pkg/parser"
)

const MAX_CONCURRENT_NETWORK_REQUESTS = 100

type Secret map[string]interface{}

var region = "us-east-1"

func SetAWSRegion(name string) {
	region = name
}

func GetAWSRegion() string {
	return region
}

// GetAWSSecret returns the value of a secret as a string. This string is encoded as base64 if the secret contained binary data.
func GetAWSSecret(name string) (Secret, error) {

	//Create a Secrets Manager client
	svc := secretsmanager.New(session.New(), aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(name),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(input)
	err = getSecretResultError(name, err)
	if err != nil {
		return nil, err
	}

	var rawSecret string
	if result.SecretString != nil {
		rawSecret = *result.SecretString
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			return nil, getSecretResultError(name, fmt.Errorf("Base64 Decode Error: %s", err))
		}
		rawSecret = string(decodedBinarySecretBytes[:len])
	}

	secret, err := stringToSecret(rawSecret)
	if err != nil {
		return nil, getSecretResultError(name, err)
	}
	return secret, nil
}

func GetAllSecrets(mappings parser.EnvKeyToSecretPath) (map[string]string, error) {
	output := make(map[string]string)
	concurrencyLimiter := make(chan struct{}, MAX_CONCURRENT_NETWORK_REQUESTS)
	results := make(chan concurrentSecretResult, len(mappings))
	for envName, secretPath := range mappings {
		go getAWSSecretConcurrently(envName, secretPath, results, concurrencyLimiter)
	}
	for i := 0; i < cap(results); i++ {
		result := <-results
		if result.err != nil {
			return nil, result.err
		}
		if _, exists := result.secret[result.secretPath.Key]; !exists {
			return nil, fmt.Errorf("AWS Secret \"%s\" does not contain key \"%s\"", result.secretPath.SecretName, result.secretPath.Key)
		}
		secretValue, err := parseSecretValue(result.secret[result.secretPath.Key])
		if err != nil {
			return nil, getSecretResultError(result.secretPath.SecretName, err)
		}
		output[result.envName] = secretValue
	}
	return output, nil
}

func parseSecretValue(value interface{}) (string, error) {
	if val, ok := value.(string); ok {
		return val, nil
	}
	if val, ok := value.(int); ok {
		return strconv.Itoa(val), nil
	}
	return "", fmt.Errorf("raw secret value is not an int or string: %s", value)
}

type concurrentSecretResult struct {
	envName    string
	secretPath parser.SecretPath
	secret     Secret
	err        error
}

func getAWSSecretConcurrently(envName string, secretPath parser.SecretPath, results chan concurrentSecretResult, concurrencyLimiter chan struct{}) {
	concurrencyLimiter <- struct{}{}
	secret, err := GetAWSSecret(secretPath.SecretName)
	<-concurrencyLimiter
	results <- concurrentSecretResult{
		envName,
		secretPath,
		secret,
		err,
	}
}

func stringToSecret(rawSecret string) (Secret, error) {
	raw := []byte(rawSecret)
	if json.Valid(raw) {
		secret := Secret{}
		err := json.Unmarshal(raw, &secret)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling JSON: %s", err)
		}
		return secret, nil
	} else {
		return nil, fmt.Errorf("invalid JSON")
	}
}

func getSecretResultError(secretName string, err error) error {
	// In this sample we only handle the specific exceptions for the 'GetSecretValue' API.
	// See https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
	if err != nil {
		message := fmt.Sprintf("error fetching secret %s", secretName)
		if awserr, ok := err.(awserr.Error); ok {
			switch awserr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				message += ": AWS Secrets Manager can't decrypt the protected secret text using the provided KMS key."
			case secretsmanager.ErrCodeInternalServiceError:
				message += ": AWS Secrets Manager experienced an internal service error"
			case secretsmanager.ErrCodeInvalidParameterException:
				message += ": An invalid parameter was provided to AWS Secret Manager"
			case secretsmanager.ErrCodeInvalidRequestException:
				message += ": An AWS Secret Manager parameter value was provided that is not valid for the current state of the resource"
			case secretsmanager.ErrCodeResourceNotFoundException:
				message += ": AWS Secret Manager resource not found"
			default:
				message += ": " + awserr.Error()
			}
		} else {
			message += ": " + err.Error()
		}
		return errors.New(message)
	}
	return nil
}
