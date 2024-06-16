package cli

import (
	"bakatools/internal/help"
	"flag"
	"fmt"
	"os"
)

func HandleCLI() {
	versionCmd := flag.NewFlagSet("version", flag.ExitOnError)
	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println(help.GetHelpMessage(true))
		os.Exit(1)
	}

	switch os.Args[1] {
	case "version":
		versionCmd.Parse(os.Args[2:])
		fmt.Println(help.VersionInfo)
	case "help":
		helpCmd.Parse(os.Args[2:])
		fmt.Println(help.GetHelpMessage(false))
	default:
		fmt.Printf("未知的子命令: %s\n", os.Args[1])
		os.Exit(1)
	}
}
