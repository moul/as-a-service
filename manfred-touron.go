package moul

import "github.com/moul/manfred-touron"

func init() {
	RegisterAction("manfred-touron", GetManfredTouronAction)
}

func GetManfredTouronAction(args []string) (interface{}, error) {
	return manfredtouron.Manfred, nil
}
