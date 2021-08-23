package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func initCLI() *cli.App {

	app := cli.NewApp()
	app.Name = "log plumber"
	app.Usage = "doing plumber work for logs"
	app.Version = "1.0.0"
	addCLICommands(app)

	app.Action = func(cnx *cli.Context) error {
		cli.ShowAppHelpAndExit(cnx, 0)
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	cli.VersionFlag = &cli.BoolFlag{
		Name:  "print-version",
		Usage: "print version",
	}

	return app
}

func addCLICommands(app *cli.App) {
	app.Commands = []*cli.Command{
		{
			Name:      "setTarget",
			Usage:     "set target folder for log files",
			UsageText: "set target folder for log files",
			Action: func(cnx *cli.Context) error {
				logPath := cnx.Args().Get(0)
				if !isFolderExist(logPath) {
					fmt.Println("folder does not exist, please try again")
					os.Exit(1)
				}
				fmt.Println("it works")
				return nil
			},
		},
	}
}

func main() {

	app := initCLI()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func isFolderExist(path string) bool {

	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}
