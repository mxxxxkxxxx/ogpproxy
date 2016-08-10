package console

import (
	"github.com/mxxxxkxxxx/ogpproxy/config"
	"fmt"
)

func Debug(msg string) {
	if config.GetConfig().Debug {
		fmt.Println("[DEBUG] " + msg)
	}
}

func Info(msg string) {
	fmt.Println("[INFO] " + msg)
}

func Error(msg string) {
	fmt.Println("[ERROR] " + msg)
}