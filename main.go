package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	fnNamePtr := flag.String("f", "", "name of the function to be deployed")
	filePathPtr := flag.String("c", "cf.yaml", "path to yaml config")

	flag.Parse()

	err := CheckInstallCLI()
	if err != nil {
		fmt.Printf("Error: %s\n\nInstall the gcloud CLI, login and choose the project you want to deploy on. \nInstallation guide: https://cloud.google.com/sdk/docs/install \n", err)
		return
	}

	if *fnNamePtr == "" {
		fmt.Printf("function name wasn't provided, be sure to use -f flag\n")
		return
	}

	configFile, err := os.ReadFile(*filePathPtr)
	if err != nil {
		fmt.Printf("%s\n\nfunction deploy aborted \n", err.Error())
		return
	}

	var config Config

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Printf("%s\n .yaml file must be incorrect you can find a full example at:\n https://gist.github.com/lafusew/c4e4ed32b09d7e91e05f29e6dcc164c3", err)
	}

	args := CreateArgs(*fnNamePtr, config)

	Exec(args)
}
