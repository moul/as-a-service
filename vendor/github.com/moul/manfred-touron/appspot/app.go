package manfredtouronapp

import (
	"encoding/json"
	"fmt"
	"net/http"

	manfred "github.com/moul/manfred-touron"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	manfred := map[string]string{
		"firstname": manfred.Firstname,
		"lastname":  manfred.Lastname,
		"headline":  manfred.Headline,
		"location":  manfred.Location,
		"github":    manfred.GitHub,
		"twitter":   manfred.Twitter,
		"website":   manfred.Website,
		"emoji":     manfred.Emoji,
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(manfred); err != nil {
		fmt.Fprintf(w, "Failed to encode Manfred Touron: %v\n", err)
	}
}
