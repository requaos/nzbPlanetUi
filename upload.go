package main

import (
	"fmt"
	"net/url"

	"github.com/labstack/gommon/log"
	sabnzbd "github.com/michaeltrobinson/go-sabnzbd"
)

func uploadNZBtoClient(dlID string) string {
	if _, ok := Settings["nzbsite"]; !ok {
		return "Check Settings"
	}
	if _, ok := Settings["nzbkey"]; !ok {
		return "Check Settings"
	}
	fmt.Printf("Sending NZB to sabNZB: %s\n", dlID)
	u, _ := url.ParseRequestURI(Settings["nzbsite"])
	u.Path = "/api"
	restpost := u.Query()
	restpost.Add("id", dlID)
	restpost.Add("apikey", Settings["nzbkey"])
	restpost.Set("t", "get")
	u.RawQuery = restpost.Encode()
	resturl := fmt.Sprintf("%v", u)
	dlIDs, err := SABnzbd.AddURL(sabnzbd.AddNzbUrl(resturl))
	if err != nil {
		log.Fatalf("failed to upload nzb: %s", err.Error())
		return "Error!"
	}
	if len(dlIDs) < 1 {
		log.Fatalf("failed to upload nzb: %s", "SABnzbd failed to return the download ID")
		return "Error!"
	}
	return "Downloading..."
}
