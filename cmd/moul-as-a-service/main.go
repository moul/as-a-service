package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	"github.com/moul/as-a-service"
)

type Action func([]string) (interface{}, error)

func GetManfredTouron(args []string) (interface{}, error) {
	return moul.GetManfredTouron(), nil
}

func GetLatestBlogPosts(args []string) (interface{}, error) {
	return moul.GetLatestBlogPosts()
}

var Actions map[string]Action

func init() {
	Actions = make(map[string]Action)
	Actions["manfred-touron"] = GetManfredTouron
	Actions["tumblr"] = GetLatestBlogPosts
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

	app.Commands = append(app.Commands, cli.Command{
		Name:        "server",
		Description: "Run as a webserver",
		Action:      Daemon,
	})

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

func Daemon(c *cli.Context) {
	r := gin.Default()
	for action, fn := range Actions {
		r.GET(fmt.Sprintf("/%s", action), func(c *gin.Context) {
			ret, err := fn(nil)
			if err != nil {
				c.JSON(500, gin.H{
					"err": err,
				})
				return
			}
			c.JSON(200, ret)
		})
	}
	r.Run(":8080")
}
