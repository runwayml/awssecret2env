package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func LoadFile(path string) (string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error loading file %s: %s", path, err)
	}
	return string(contents), nil
}

func SaveFile(path, contents string, mode os.FileMode) error {
	err := ioutil.WriteFile(path, []byte(contents), mode)
	if err != nil {
		return fmt.Errorf("error writing file %s: %s", path, err)
	}
	return nil
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
	lines := make([]string, len(secrets))
	for key, value := range secrets {
		// Escape single quotes via:
		//	https://stackoverflow.com/questions/1250079/how-to-escape-single-quotes-within-single-quoted-strings
		escaped := strings.ReplaceAll(value, "'", "'\"'\"'")
		lines = append(lines, fmt.Sprintf("%s='%s'\n", key, escaped))
	}
	sort.Strings(lines)
	return strings.Join(lines, "")
}

func PrependExportStatementsBeforeEachLine(input string) string {
	lines := strings.Split(input, "\n")
	output := ""
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			output += "export " + line + "\n"
		}
	}
	return output
}
