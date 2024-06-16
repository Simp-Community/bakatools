package menu

import (
	"bakatools/internal/help"
	"bufio"
	"fmt"
	"strings"
)

func ShowMenu(reader *bufio.Reader) string {
	options := help.GetMenuOptions()

	for {
		fmt.Println(help.GetHelpMessage(true))
		fmt.Print("请输入选项: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入错误:", err)
			return ""
		}

		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			return options[0]
		case "2":
			return options[1]
		default:
			fmt.Println("无效的选项")
		}
	}
}
