package main

import (
	"log"
	"os"
	scli "serpent_aes/serpentcli"
	"sort"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func registerAppInfo() {
	app.Name = "Serpent CLI Application"
	app.Usage = "CLI Application used to encrypt and decrypt text with serpent aes"
	app.Author = "captaincode0"
	app.Version = "0.0.1"
}

func registerAppCommands() {
	app.Commands = []cli.Command{
		{
			Name:    scli.EncryptCommand.Name,
			Aliases: scli.EncryptCommand.Aliases,
			Usage:   scli.EncryptCommand.UsageText,
			Action:  scli.EncryptCommand.Exec,
		},
		{
			Name:    scli.DecryptCommand.Name,
			Aliases: scli.DecryptCommand.Aliases,
			Usage:   scli.DecryptCommand.UsageText,
			Action:  scli.DecryptCommand.Exec,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
}

func main() {
	registerAppInfo()
	registerAppCommands()

	errors := app.Run(os.Args)

	if errors != nil {
		log.Fatal(errors)
	}
}
