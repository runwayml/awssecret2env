package main

import (
	"fmt"
	"os"

	"github.com/runwayml/awssecret2env/pkg/awssecretsmanager"
	"github.com/runwayml/awssecret2env/pkg/parser"
	"github.com/runwayml/awssecret2env/pkg/utils"
	flag "github.com/spf13/pflag"
)

func main() {

	args := parseArgs()

	input, err := utils.LoadFile(args.inputFilename)
	utils.HandleError(err)

	mappings, err := parser.ParseInput(input)
	utils.HandleError(err)

	awssecretsmanager.SetAWSRegion(args.awsRegion)
	secrets, err := awssecretsmanager.GetAllSecrets(mappings)
	utils.HandleError(err)

	output := utils.SecretsToEnvString(secrets)
	if args.export {
		output = utils.PrependExportStatementsBeforeEachLine(output)
	}

	if args.outputFilename != "" {
		err = utils.SaveFile(args.outputFilename, output, 0600)
		utils.HandleError(err)
	} else {
		fmt.Print(output)
	}
}

type Args struct {
	inputFilename  string
	outputFilename string
	awsRegion      string
	export         bool
}

func parseArgs() Args {

	output := flag.StringP("output", "o", "", "Redirects output to a file instead of stdout")
	awsRegion := flag.StringP("aws-region", "r", "us-east-1", "The name of the AWS region where secrets are stored")
	export := flag.BoolP("export", "e", false, "Prepends \"export\" statements in front of the output env variables")
	help := flag.BoolP("help", "h", false, "Show this screen")

	flag.Parse()
	flag.Usage = func() {
		fmt.Printf("Usage: %s [OPTIONS] <input-file> ...\n", os.Args[0])
		fmt.Println("Note: <input-file> is a required positional argument.")
		flag.PrintDefaults()
	}
	if flag.NArg() != 1 || *help {
		flag.Usage()
		os.Exit(1)
	}
	inputFilename := flag.Args()[0]
	if fileInfo, err := os.Stat(inputFilename); os.IsNotExist(err) || fileInfo.IsDir() {
		if fileInfo != nil && fileInfo.IsDir() {
			utils.PrintErrorAndExit(fmt.Errorf("input file \"%s\" must be a file, not a directory", inputFilename))
		} else {
			utils.PrintErrorAndExit(fmt.Errorf("input file \"%s\" does not exist", inputFilename))
		}
	}
	return Args{
		inputFilename:  inputFilename,
		outputFilename: *output,
		awsRegion:      *awsRegion,
		export:         *export,
	}
}
