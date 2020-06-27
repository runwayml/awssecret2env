package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func LoadFile(path string) (string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error loading file %s: %s", path, err)
	}
	return string(contents), nil
}

func PrintUsageAndExit() {
	fmt.Println("Usage: awssecret2env <input-file>")
	os.Exit(1)
}

func PrintErrorAndExit(err error) {
	os.Stderr.WriteString("Fatal: " + err.Error() + "\n")
	os.Exit(1)
}

func HandleError(err error) {
	if err != nil {
		PrintErrorAndExit(err)
	}
}

func SecretsToEnvString(secrets map[string]string) string {
	output := ""
	for key, value := range secrets {
		output += fmt.Sprintf("%s=%s\n", key, value)
	}
	return output
}
