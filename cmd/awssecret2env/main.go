package main

import (
	"fmt"
	"os"

	"github.com/runwayml/awssecret2env/awssecretsmanager"
	"github.com/runwayml/awssecret2env/parser"
	"github.com/runwayml/awssecret2env/utils"
)

func main() {

	args := parseArgs()

	input, err := utils.LoadFile(args.inputFile)
	utils.HandleError(err)

	mappings, err := parser.ParseInput(input)
	utils.HandleError(err)

	awssecretsmanager.SetAWSRegion("us-east-1")
	secrets, err := awssecretsmanager.GetAllSecrets(mappings)
	utils.HandleError(err)

	fmt.Print(utils.SecretsToEnvString(secrets))
}

type Args struct {
	inputFile string
}

func parseArgs() Args {

	if len(os.Args) != 2 {
		utils.PrintUsageAndExit()
	}

	return Args{
		inputFile: os.Args[1],
	}
}
