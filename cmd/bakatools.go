package main

import (
	"bakatools/internal/cli"
	"bakatools/internal/menu"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	if len(os.Args) > 1 {
		cli.HandleCLI()
		return
	}

	for {
		command := menu.ShowMenu(reader)
		if command != "" {
			os.Args = []string{"", command}
			cli.HandleCLI()
		}

		if !askForRestart(reader) {
			break
		}

		os.Args = []string{os.Args[0]}
	}
}

func askForRestart(reader *bufio.Reader) bool {
	for {
		fmt.Print("是否重启程序? (y/n): ")
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入错误:", err)
			continue
		}
		response = strings.TrimSpace(response)
		if response != "" && (response == "y" || response == "Y" || response == "n" || response == "N") {
			return response == "y" || response == "Y"
		}
	}
}
