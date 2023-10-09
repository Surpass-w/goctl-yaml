package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/zeromicro/goctl-yaml/action"
	"os"
	"runtime"
)

var (
	version  = "20231009"
	commands = []*cli.Command{
		{
			Name:   "authenticate",
			Usage:  "generate authenticate yaml file",
			Action: action.Generator,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "filename",
					Usage: "authenticate save file name",
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Usage = "a plugin of goctl to generate authenticate yaml file for vmp_gateway"
	app.Version = fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("goctl-yaml: %+v\n", err)
	}
}
