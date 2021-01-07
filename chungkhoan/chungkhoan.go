package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "Chứng khoán CLI",
		Usage: "Đốt tiền",
		Action: func(c *cli.Context) error {
			exec.Command("open", "https://banggia.vps.com.vn/").Run()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
