package moul

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
)

type KeybaseProof struct {
	HumanURL          string `json:"human_url"`
	Nametag           string `json:"nametag"`
	PresentationGroup string `json:"presentation_group"`
	PresentationTag   string `json:"presentation_tag"`
	ProofID           string `json:"proof_id"`
	ProofType         string `json:"proof_type"`
	ProofURL          string `json:"proof_url"`
	ServiceURL        string `json:"service_url"`
	SigID             string `json:"sig_id"`
	State             int    `json:"state"`
}

type KeybaseUser struct {
	Basics struct {
		Ctime         int    `json:"ctime"`
		IDVersion     int    `json:"id_version"`
		LastIDChange  int    `json:"last_id_change"`
		Mtime         int    `json:"mtime"`
		TrackVersion  int    `json:"track_version"`
		Username      string `json:"username"`
		UsernameCased string `json:"username_cased"`
	} `json:"basics"`
	CryptocurrencyAddresses map[string][]struct {
		Address string `json:"address"`
		SigID   string `json:"sig_id"`
	} `json:"cryptocurrency_addresses"`
	Devices  struct{} `json:"devices"`
	ID       string   `json:"id"`
	Pictures map[string]struct {
		Height int    `json:"height"`
		Source string `json:"source"`
		URL    string `json:"url"`
		Width  int    `json:"width"`
	} `json:"pictures"`
	Profile struct {
		Bio      string `json:"bio"`
		FullName string `json:"full_name"`
		Location string `json:"location"`
		Mtime    int    `json:"mtime"`
	} `json:"profile"`
	ProofsSummary struct {
		All                 []KeybaseProof            `json:"all"`
		ByPresentationGroup map[string][]KeybaseProof `json:"by_presentation_group"`
		ByProofType         map[string][]KeybaseProof `json:"by_proof_type"`
	} `json:"proofs_summary"`
	PublicKeys struct {
		AllBundles           []string            `json:"all_bundles"`
		EldestKeyFingerprint string              `json:"eldest_key_fingerprint"`
		EldestKid            string              `json:"eldest_kid"`
		Families             map[string][]string `json:"families"`
		PgpPublicKeys        []string            `json:"pgp_public_keys"`
		Primary              struct {
			Bundle         string      `json:"bundle"`
			Ctime          int         `json:"ctime"`
			EldestKid      interface{} `json:"eldest_kid"`
			Etime          int         `json:"etime"`
			KeyAlgo        int         `json:"key_algo"`
			KeyBits        int         `json:"key_bits"`
			KeyFingerprint string      `json:"key_fingerprint"`
			KeyLevel       int         `json:"key_level"`
			KeyType        int         `json:"key_type"`
			Kid            string      `json:"kid"`
			Mtime          int         `json:"mtime"`
			SelfSigned     bool        `json:"self_signed"`
			SigningKid     string      `json:"signing_kid"`
			Status         int         `json:"status"`
			Ukbid          string      `json:"ukbid"`
		} `json:"primary"`
		Sibkeys []string      `json:"sibkeys"`
		Subkeys []interface{} `json:"subkeys"`
	} `json:"public_keys"`
	Sigs struct {
		Last struct {
			PayloadHash string `json:"payload_hash"`
			Seqno       int    `json:"seqno"`
			SigID       string `json:"sig_id"`
		} `json:"last"`
	} `json:"sigs"`
}

type KeybaseResponse struct {
	Status struct {
		Code int
		Name string
	}
	Them []KeybaseUser
}

const KeybaseLookupURL = "https://keybase.io/_/api/1.0/user/lookup.json?usernames=moul"

func GetKeybaseProfile() (*KeybaseUser, error) {
	_, body, errs := gorequest.New().Get(KeybaseLookupURL).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	var response KeybaseResponse
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		return nil, err
	}

	user := response.Them[0]
	return &user, nil

}
