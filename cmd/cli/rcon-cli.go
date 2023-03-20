package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/Sch8ill/rcon"
	"github.com/Sch8ill/rcon/config"
)

func main() {
	app := &cli.App{
		Flags:     declareFlags(),
		Action:    cliAction,
		Usage:     "A rcon CLI for excecuting commands on a remote rcon server",
		Copyright: "Copyright (c) 2023 Sch8ill",
		Version:   config.Version,
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func cliAction(c *cli.Context) error {
	rconClient, err := rcon.Dial(c.String("address"), c.String("password"), c.Duration("timeout"))

	if err != nil {
		fmt.Println(fmt.Errorf("error while trying to connect: %w", err))
		return nil
	}
	return serveInteractive(rconClient)
}

func declareFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "address",
			Aliases: []string{"a"},
			Value:   config.DefaultAddress,
			Usage:   "address of the server you want to connect to (localhost:25575 for example)",
		},
		&cli.StringFlag{
			Name:    "password",
			Aliases: []string{"p"},
			Value:   config.DefaulftPassword,
			Usage:   "password of the RCON server you want to connect to",
		},
		&cli.DurationFlag{
			Name:    "timeout",
			Aliases: []string{"t"},
			Value:   config.DefaultTimeout,
			Usage:   "timeout for the connection to the server",
		},
	}
}

func serveInteractive(rconClient *rcon.RconClient) error {
	cliPrefix := generateCliPrefix(rconClient.Address)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Conencted to %s. Type ':help' to see a list of available commands provided by this shell.\n", rconClient.Address)

	run:
	for {
		fmt.Print(cliPrefix)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		switch cmd {
		case "":
			continue

		case ":help":
			fmt.Println("Not implemented yet ):")

		case ":exit":
			close(rconClient)
			break run

		default:
			response, err := executeCmd(cmd, rconClient)
			if err != nil {
				panic(err)
			}
			fmt.Println(response)
		}
	}
	return nil
}

func executeCmd(cmd string, rconClient *rcon.RconClient) (string, error) {
	res, err := rconClient.ExecuteCmd(cmd)
	if err != nil {
		return "", err
	}
	return res, nil
}

func close(rconClient *rcon.RconClient) {
	fmt.Println("Closing rcon connection...")
	rconClient.Close()
}

func generateCliPrefix(addr string) string {
	return fmt.Sprintf("rcon@%s:~$ ", addr)
}
