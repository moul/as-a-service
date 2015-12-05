package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/moul/as-a-service"
)

type Action func([]string) (interface{}, error)

func GetManfredTouron(args []string) (interface{}, error) {
	return moul.GetManfredTouron(), nil
}

var Actions map[string]Action

func init() {
	Actions = make(map[string]Action)
	Actions["manfred-touron"] = GetManfredTouron
}

func main() {
	app := cli.NewApp()
	app.Name = "moul-as-a-service"
	app.Usage = "moul, but as a service"
	app.Commands = []cli.Command{}

	for action := range Actions {
		command := cli.Command{
			Name:   action,
			Action: CliActionCallback,
		}
		app.Commands = append(app.Commands, command)
	}

	app.Run(os.Args)
}

func CliActionCallback(c *cli.Context) {
	action := c.Command.Name
	ret, err := Actions[action](c.Args())
	if err != nil {
		logrus.Fatalf("Failed to execute %q: %v", action, err)
	}

	out, err := json.MarshalIndent(ret, "", "  ")
	if err != nil {
		logrus.Fatalf("Failed to marshal json: %v", err)
	}
	fmt.Printf("%s\n", out)
}
