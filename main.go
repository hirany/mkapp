package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

const (
	perm = 755
)

func main() {
	app := cli.NewApp()
	app.Name = "lt"
	app.Usage = "show directory"
	app.Action = func(c *cli.Context) {
		err := mkapp(c.Args()[0])
		if err != nil {
			log.Fatal(err)
		}
	}
	app.Run(os.Args)
}

func mkapp(cliName string) error {
	err := os.Mkdir(cliName, perm)
	if err != nil {
		return err
	}
	err = os.Chdir(cliName)
	if err != nil {
		return err
	}
	file, err := os.Create("main.go")
	if err != nil {
		return err
	}
	defer file.Close()
	inCode := `package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "`
	inCode += cliName
	inCode += `"
	app.Usage = "how to use"
	app.Action = func(c *cli.Context) {
		run(c.Args())
	}
	app.Run(os.Args)
}

func run(args cli.Args) {
}`
	_, err = file.WriteString(inCode)
	if err != nil {
		return err
	}

	return nil
}
