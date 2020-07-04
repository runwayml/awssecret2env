package parser

import (
	"fmt"
	"strings"
)

type SecretPath struct {
	SecretName string
	Key        string
}
type EnvKeyToSecretPath map[string]SecretPath

func ParseInput(input string) (EnvKeyToSecretPath, error) {
	output := make(EnvKeyToSecretPath)
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	if len(lines) < 1 {
		return nil, fmt.Errorf("parse input error: input string contains no lines")
	}
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			continue
		}
		if !strings.Contains(line, "=") {
			return nil, fmt.Errorf("parse input error: line %d missing required \"=\" separated env pair", i+1)
		}
		pair := strings.Split(line, "=")
		if len(pair) != 2 {
			return nil, fmt.Errorf("parse input error: line %d must contain a pair of exactly one env key to secret name", i+1)
		}
		envKey := pair[0]
		secret := pair[1]
		if _, exists := output[envKey]; exists {
			return nil, fmt.Errorf("parse input error: duplicate env key %s", envKey)
		}

		if !strings.Contains(secret, "/") {
			return nil, fmt.Errorf("parse input error: secret name on line %d must contain at least one \"/\", in the format secretName/key", i+1)
		}
		secretParts := strings.Split(secret, "/")
		output[envKey] = SecretPath{
			SecretName: strings.Join(secretParts[:len(secretParts)-1], "/"),
			Key:        secretParts[len(secretParts)-1],
		}
	}
	if len(output) < 1 {
		return nil, fmt.Errorf("parse input error: no secrets defined in input file")
	}

	return output, nil
}
