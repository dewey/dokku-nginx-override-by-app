package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: nginx-app-template-source <arg1> <arg2>")
	}
	appPluginConfigPath := fmt.Sprintf("/var/lib/dokku/data/nginx-override-by-hostname/%s/nginx.conf.sigil", args[1])
	if _, err := os.Stat(appPluginConfigPath); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("not exists")
			return
		}
		fmt.Println("err", err)
		return
	}
	fmt.Println(appPluginConfigPath)
}
