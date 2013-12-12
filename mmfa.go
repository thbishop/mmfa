package main

import (
	"os"

	"./commands"

	"github.com/codegangsta/cli"
)

func main() {
	os.Exit(realMain())
}

func realMain() (exitCode int) {
	exitCode = 0

	app := cli.NewApp()
	app.Name = "mmfa"
	app.Usage = "osx keychain mfa manager"
	app.Version = Version + VersionPrerelease

	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "add a new keychain item",
			Action: func(c *cli.Context) {
				if !c.Args().Present() {
					println("Name not provided\n")
					cli.ShowCommandHelp(c, "add")
					exitCode = 1
					return
				}
				commands.AddKeychainItem(c.Args().First())
			},
		},
		{
			Name:  "get",
			Usage: "get the code for a keychain item",
			Action: func(c *cli.Context) {
				if !c.Args().Present() {
					println("Name not provided\n")
					cli.ShowCommandHelp(c, "get")
					exitCode = 1
					return
				}
				commands.GetCode(c.Args().First())
			},
		},
	}

	app.Run(os.Args)

	return
}
