package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: nginx-app-template-source <arg1> <arg2>")
		return
	}
	// We only want to use this custom template for the app-config argument (https://github.com/dokku/dokku/discussions/6991#discussioncomment-10040689)
	if args[2] != "app-config" {
		return
	}
	appPluginConfigPath := fmt.Sprintf("/var/lib/dokku/data/nginx-override-by-app/%s/nginx.conf.sigil", args[1])
	if _, err := os.Stat(appPluginConfigPath); err != nil {
		if os.IsNotExist(err) {
			return
		}
		return
	}
	fmt.Println(appPluginConfigPath)
}
