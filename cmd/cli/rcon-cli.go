package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/Sch8ill/rcon"
	"github.com/Sch8ill/rcon/color"
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

// the entrypoint of the cli
func cliAction(ctx *cli.Context) error {
	rconClient, err := rcon.Dial(ctx.String("address"), ctx.String("password"), ctx.Duration("timeout"))

	if err != nil {
		fmt.Println(fmt.Errorf("error while trying to connect: %w", err))
		return nil
	}

	// check for "single command mode"
	if cmd := ctx.String("command"); cmd != "" {
		if err := x(cmd, rconClient, ctx); err != nil {
			return err
		}
		return nil
	}
	return serveInteractive(rconClient, ctx)
}

// declares the flags used by the cli
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
		&cli.StringFlag{
			Name:    "command",
			Aliases: []string{"c"},
			Value:   "",
			Usage:   "a single command to be executed",
		},
		&cli.DurationFlag{
			Name:    "timeout",
			Aliases: []string{"t"},
			Value:   config.DefaultTimeout,
			Usage:   "timeout for the connection to the server",
		},
		&cli.BoolFlag{
			Name:    "no-colors",
			Aliases: []string{"no-colours"},
			Value:   false,
			Usage:   "if the cli should not output colors",
		},
	}
}

// serves an interactive RCON shell
func serveInteractive(rconClient *rcon.RconClient, ctx *cli.Context) error {
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
			if err := x(cmd, rconClient, ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

// executes the command and formats and prints the response
func x(cmd string, rconClient *rcon.RconClient, ctx *cli.Context) error {
	response, err := rconClient.ExecuteCmd(cmd)
	if err != nil {
		return fmt.Errorf("error while executing command %w", err)
	}
	response = color.ParseColorCodes(response, !ctx.Bool("no-colors"))
	fmt.Println(response)
	return nil
}

// closes the underlying RCON client
func close(rconClient *rcon.RconClient) {
	fmt.Println("Closing rcon connection...")
	rconClient.Close()
}

// generates the cli new line prefix
func generateCliPrefix(addr string) string {
	return fmt.Sprintf("rcon@%s:~$ ", addr)
}
