package moul

import (
	"math/rand"
	"strings"

	"github.com/moul/as-a-service/code"
)

func init() {
	RegisterAction("github-code", GetGithubCodeAction)
}

func GetGithubCodeAction(args []string) (interface{}, error) {
	return GetGithubCode(2000)
}

func GetGithubCode(length int) (interface{}, error) {
	content, err := moulcode.Asset("dump.txt")
	if err != nil {
		return content, err
	}
	lines := strings.Split(string(content), "\n")
	nbLines := len(lines)

	start := rand.Intn(nbLines - length)
	return lines[start : start+length], nil
}
