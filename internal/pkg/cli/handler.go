package cli

import (
	"bufio"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"os"
)

func ReadFromCLI() (string,error) {
	var input string

	app := cli.NewApp()
	reader := bufio.NewReader(os.Stdin)

	app.Flags = []cli.Flag {
		&cli.StringFlag{
			Name: "lang",
			Value: "en",
			Usage: "language for the greeting",
			Required: true,
		},
	}

	app.Action = func(c *cli.Context) error {

		if c.String("lang") == "de" {
			fmt.Print("Trage den Kartennamen ein: ")
			input, _ = reader.ReadString('\n')
		} else {
			fmt.Print("Enter the card name: ")
			input, _ = reader.ReadString('\n')
		}
		return nil
	}



	err := app.Run(os.Args)
	if err != nil {
		log.Error().Err(err)
		return "", err
	}

	return input, err
}
