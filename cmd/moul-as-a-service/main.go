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

func main() {
	app := cli.NewApp()
	app.Name = "moul-as-a-service"
	app.Usage = "moul, but as a service"
	app.Commands = []cli.Command{}

	for action := range moul.Actions() {
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
	ret, err := moul.Actions()[action](c.Args())
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

	// Register index
	r.GET("/", func(c *gin.Context) {
		services := []string{}
		for action := range moul.Actions() {
			services = append(services, fmt.Sprintf("/%s", action))
		}
		c.JSON(200, gin.H{
			"services": services,
		})
	})

	// Register actions
	for action, fn := range moul.Actions() {
		fmt.Println(action, fn)
		func(action string, fn moul.Action) {
			callback := func(c *gin.Context) {
				//u, err := url.Parse(c.Request.URL.String())
				//if err != nil {
				//	c.String(500, fmt.Sprintf("failed to poarse url %q: %v", c.Request.URL.String(), err))
				//}

				ret, err := fn(nil)
				// ret, err :- fn(u.RawQuery, c.Request.Body)
				if err != nil {
					c.JSON(500, gin.H{
						"err": err,
					})
					return
				}

				// FIXME: handle content-types
				c.JSON(200, ret)
			}
			r.GET(fmt.Sprintf("/%s", action), callback)
			// POST
		}(action, fn)
	}

	// Start server
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	r.Run(fmt.Sprintf(":%s", port))
}
