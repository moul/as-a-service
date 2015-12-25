package moul

import "github.com/moul/manfred-touron"

type ManfredTouron struct {
	Firstname string
	Lastname  string
	Website   string
	GitHub    string
	Twitter   string
	Location  string
	Headline  string
	Emoji     string
	Groups    []string
}

func GetManfredTouronAction(args []string) (interface{}, error) {
	return GetManfredTouron(), nil
}

func GetManfredTouron() ManfredTouron {
	return ManfredTouron{
		Firstname: manfredtouron.Firstname,
		Lastname:  manfredtouron.Lastname,
		Website:   manfredtouron.Website,
		GitHub:    manfredtouron.GitHub,
		Twitter:   manfredtouron.Twitter,
		Location:  manfredtouron.Location,
		Headline:  manfredtouron.Headline,
		Emoji:     manfredtouron.Emoji,
	}
}
