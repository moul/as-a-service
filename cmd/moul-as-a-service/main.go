package main

import (
	"encoding/json"
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

	out, err := json.MarshalIndent(ret, "", "  ")
	if err != nil {
		logrus.Fatalf("Failed to marshal json: %v", err)
	}
	fmt.Printf("%s\n", out)
}
