package serpentcli

import (
	"errors"
	"fmt"
	s "serpent_aes/serpent"

	"github.com/urfave/cli"
)

var DecryptCommand Command = Command{
	Aliases:   []string{"d"},
	Name:      "decrypt",
	UsageText: "decrypt <ENCRYPTED-TEXT> <USER-KEY>",
	Exec: func(context *cli.Context) error {
		if context.NArg() < 2 {
			return errors.New("Error: To encrypt needs to introduce encrypted text followed by user key")
		}

		var serpent s.Serpent
		var encryptedText string = context.Args().Get(0)
		var userKey string = context.Args().Get(1)

		fmt.Println("-- ENCRYPTED TEXT --")
		fmt.Println(encryptedText)
		fmt.Println()

		fmt.Println("-- DECRYPTED TEXT --")
		fmt.Println(serpent.DecryptText(encryptedText, userKey))
		fmt.Println()

		return nil
	},
}
