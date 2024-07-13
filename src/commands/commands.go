package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"strconv"
	"strings"

	columnize "github.com/ryanuber/columnize"
)

const (
	helpHeader = `Usage: dokku repo[:COMMAND]

Runs commands that interact with the app's repo

Additional commands:`

	helpContent = `
    nginx-override-by-app:add <app> <path/to/custom/nginx.conf.sigil>, copies custom nginx configuration file to the app's nginx configuration directory
`
)

func main() {
	flag.Usage = usage
	flag.Parse()

	cmd := flag.Arg(0)
	switch cmd {
	case "nginx-override-by-app:add":
		appPluginBasePath := "/var/lib/dokku/data/nginx-override-by-app"
		if len(flag.Args()) != 3 {
			usage()
			return
		}
		appPluginAppBasePath := path.Join(appPluginBasePath, flag.Arg(1))
		err := os.MkdirAll(appPluginAppBasePath, os.ModePerm)
		if err != nil {
			log.Println("error creating plugin data directory", err)
			return
		}
		u, err := user.Lookup("dokku")
		if err != nil {
			fmt.Println("error getting user info for dokku user")
			return
		}
		uid, err := strconv.Atoi(u.Uid)
		if err != nil {
			fmt.Println("error converting uid to int")
			return
		}
		gid, err := strconv.Atoi(u.Gid)
		if err != nil {
			fmt.Println("error converting gid to int")
			return
		}
		if err := os.Chown(appPluginAppBasePath, uid, gid); err != nil {
			fmt.Println("error changing file owner of plugin app base path", err)
			return
		}
		configPath := flag.Arg(2)
		if _, err := os.Stat(configPath); err != nil {
			if os.IsNotExist(err) {
				return
			}
		}
		input, err := os.ReadFile(configPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		destinationFile := path.Join(appPluginAppBasePath, "nginx.conf.sigil")
		err = os.WriteFile(destinationFile, input, 0644)
		if err != nil {
			fmt.Printf("error creating %s: %v\n", destinationFile, err)
			return
		}

		if err := os.Chown(destinationFile, uid, gid); err != nil {
			fmt.Println("error changing file owner", err)
			return
		}

		fmt.Printf("success! the nginx.conf.sigil was copied to: %s, use ps:rebuild to apply.\n", appPluginAppBasePath)
	case "smoke-test-plugin:help":
		usage()
	case "help":
		fmt.Print(helpContent)
	default:
		dokkuNotImplementExitCode, err := strconv.Atoi(os.Getenv("DOKKU_NOT_IMPLEMENTED_EXIT"))
		if err != nil {
			fmt.Println("failed to retrieve DOKKU_NOT_IMPLEMENTED_EXIT environment variable")
			dokkuNotImplementExitCode = 10
		}
		os.Exit(dokkuNotImplementExitCode)
	}
}

func usage() {
	config := columnize.DefaultConfig()
	config.Delim = ","
	config.Prefix = "\t"
	config.Empty = ""
	content := strings.Split(helpContent, "\n")[1:]
	fmt.Println(helpHeader)
	fmt.Println(columnize.Format(content, config))
}
