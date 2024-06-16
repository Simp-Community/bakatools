package help

import (
	"bakatools/internal/constants/core_constant"
	"fmt"
	"sort"
)

const VersionInfo = core_constant.NAME + " 版本 " + core_constant.VERSION

var commands = map[string]string{
	"version": "显示版本信息",
	"help":    "显示帮助信息",
}

func GetHelpMessage(withNumbers bool) string {
	keys := make([]string, 0, len(commands))
	for key := range commands {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var message string
	for i, key := range keys {
		if withNumbers {
			message += fmt.Sprintf("%d. %s - %s\n", i+1, key, commands[key])
		} else {
			message += fmt.Sprintf("%s - %s\n", key, commands[key])
		}
	}

	if len(message) > 0 && message[len(message)-1] == '\n' {
		message = message[:len(message)-1]
	}

	return "可用的命令:\n" + message
}

func GetMenuOptions() []string {
	keys := make([]string, 0, len(commands))
	for key := range commands {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
