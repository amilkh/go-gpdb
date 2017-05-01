package argParser

import (
	"flag"
	"os"
	"fmt"
)

import (
	"../../pkg/core/arguments"
	"../../pkg/core/methods"
)

// OS Argument Parser
func ArgParser() {

	// Download Command Parser
	downloadCmd := flag.NewFlagSet("download", flag.ExitOnError)
	DownloadProductFlag := downloadCmd.String("p", "gpdb", "What product do you want to Install? [OPTIONS: gpdb, gpcc, gpextras]")
	DownloadVersionFlag := downloadCmd.String("v", "", "OPTIONAL: Which GPDB version software do you want to list ?")
	DownloadInstallFlag := downloadCmd.Bool("install", false, "OPTIONAL: Install after downloaded (Only works with \"gpdb\")?")

	// Install Command Parser
	InstallCmd := flag.NewFlagSet("install", flag.ExitOnError)
	InstallProductFlag := InstallCmd.String("p", "gpdb", "What product do you want to Install [OPTIONS: gpdb, gpcc]?")
	InstallVersionFlag := InstallCmd.String("v", "", "Which version do you want to Install ?")

	// Remove Command Parser
	RemoveCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	RemoveVersionFlag := RemoveCmd.String("v", "", "Provide the version from the installed list to remove")

	// Environment Command Parser
	EnvCmd := flag.NewFlagSet("env", flag.ExitOnError)
	EnvListFlag := EnvCmd.String("l", "", "List all the version that is currently installed")
	EnvVersionFlag := EnvCmd.String("v", "", "Provide the version from the installed list to remove")

	// If no COMMAND keyword provided then show the help menu.
	if len(os.Args) == 1 {
		ShowHelp()
	}

	// If there is a command keyword provided then check to what is it and then parse the appropriate options
	switch os.Args[1] {
		case "download":                                                // Download command parser
			arguments.ArgOption = "download"
			downloadCmd.Parse(os.Args[2:])
		case "install":                                                 // Install command parser
			arguments.ArgOption = "install"
			InstallCmd.Parse(os.Args[2:])
		case "env":                                                     // Env command parser
			arguments.ArgOption = "env"
			EnvCmd.Parse(os.Args[2:])
		case "remove":                                                  // Remove command parser
			arguments.ArgOption = "remove"
			RemoveCmd.Parse(os.Args[2:])
		case "version":                                                 // Version of the software
			fmt.Printf("Version: %.1f\n", arguments.Version)
			os.Exit(0)
		case "help":                                                    // Help Menu
			ShowHelp()
		default:                                                        // Error if command is invalid
			fmt.Printf("ERROR: %q is not valid command.\n", os.Args[1])
			ShowHelp()
	}

	// If the command send is download, then parse the commandline arguments
	if downloadCmd.Parsed() {

		// If the product parameter is passed, then check if its valid value.
		if *DownloadProductFlag != "" {

			// If its valid then we are going to store it
			if methods.Contains(arguments.AcceptedDownloadProduct, *DownloadProductFlag) {
				arguments.RequestedDownloadProduct = *DownloadProductFlag
			} else { // Else print error to choose the right value
				fmt.Println("ERROR: Invalid options provided to the argument -p")
				fmt.Print("Usage of download: \n")
				downloadCmd.PrintDefaults()
				os.Exit(2)
			}
		}

		// If the version parameter is passed, then store the value
		if *DownloadVersionFlag != "" {
			arguments.RequestedDownloadVersion = *DownloadVersionFlag
		}

		// If the user request to install the product after download
		if *DownloadInstallFlag && arguments.RequestedDownloadProduct == "gpdb" {
			arguments.InstallAfterDownload = true
		} else if *DownloadInstallFlag {
			fmt.Println("ERROR: -install only works with -p \"gpdb\"")
			fmt.Print("Usage of download: \n")
			downloadCmd.PrintDefaults()
			os.Exit(2)
		}
	}


	// Install command argument parser
	if InstallCmd.Parsed() {
		if *InstallProductFlag != "" {
			// If its valid then we are going to store it
			if methods.Contains(arguments.AcceptedInstallProduct, *InstallProductFlag) {
				arguments.RequestedInstallProduct = *InstallProductFlag
			} else { // Else print error to choose the right value
				fmt.Println("ERROR: Invalid options provided to the argument -p")
				fmt.Print("Usage of install: \n")
				InstallCmd.PrintDefaults()
				os.Exit(2)
			}
		}

		// If the version parameter is passed, then store the value
		if *InstallVersionFlag != "" {
			arguments.RequestedInstallVersion = *InstallVersionFlag
		} else { // On install this is a required parameter, if not provided then terminate and ask to enter the version
			fmt.Println("ERROR: Please provide the version of the product that needs to be installed")
			fmt.Print("Usage of install: \n")
			InstallCmd.PrintDefaults()
			os.Exit(2)
		}
	}

	// All the below command line parse will be updated when the function is written
	if RemoveCmd.Parsed() {

		if *RemoveVersionFlag == "" {
			fmt.Println("Please supply the message using -message option.")
			return
		}

		fmt.Printf("Your message is sent to %q.\n", *RemoveVersionFlag)
	}


	if EnvCmd.Parsed() {

		if *EnvListFlag == "" {
			fmt.Println("Please supply the message using -message option.")
			return
		}
		if *EnvVersionFlag == "" {
			fmt.Println("Please supply the message using -message option.")
			return
		}

		fmt.Printf("Your message is sent to %q.\n", *EnvListFlag)
		fmt.Printf("Message: %q.\n", *EnvVersionFlag)
	}

}