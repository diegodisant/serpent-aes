package serpentcli

import (
	"errors"
	"fmt"
	s "serpent_aes/serpent"

	"github.com/urfave/cli"
)

var EncryptCommand Command = Command{
	Aliases:   []string{"e"},
	Name:      "encrypt",
	UsageText: "encrypt <PLAIN-TEXT> <USER-KEY>",
	Exec: func(context *cli.Context) error {
		if context.NArg() < 2 {
			return errors.New("Error: To encrypt needs to introduce plain text followed by user key")
		}

		var serpent s.Serpent
		var plainText string = context.Args().Get(0)
		var userKey string = context.Args().Get(1)

		fmt.Println("-- PLAIN TEXT --")
		fmt.Println(plainText)
		fmt.Println()

		fmt.Println("-- ENCRYPTED TEXT --")
		fmt.Println(serpent.EncryptText(plainText, userKey))
		fmt.Println()

		return nil
	},
}
