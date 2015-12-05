package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/moul/as-a-service"
)

type Action func([]string) (interface{}, error)
type Actions map[string]Action

func GetManfredTouron(args []string) (interface{}, error) {
	return moul.GetManfredTouron(), nil
}

func main() {
	actions := map[string]Action{
		"manfred-touron": GetManfredTouron,
	}

	action := "manfred-touron"
	ret, err := actions[action](nil)
	if err != nil {
		logrus.Fatalf("Failed to execute %q: %v", action, err)
	}
	fmt.Printf("%v\n", ret)
}
