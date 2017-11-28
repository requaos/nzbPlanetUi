package main

import (
	"fmt"
	"net/url"

	"github.com/labstack/gommon/log"
	sabnzbd "github.com/michaeltrobinson/go-sabnzbd"
	"github.com/therecipe/qt/core"
)

type QmlBridge struct {
	core.QObject

	_ func(data string) string                `slot:"sendToGo"`
	_ func(model *PersonModel, search string) `slot:"resetList"`
}

func uploadNZBtoClient(dlID string) string {
	fmt.Printf("Sending NZB to sabNZB: %s\n", dlID)
	settings := Settings
	u, _ := url.ParseRequestURI(settings["nzbsite"])
	u.Path = "/api"
	restpost := u.Query()
	restpost.Add("id", dlID)
	restpost.Add("apikey", settings["nzbkey"])
	restpost.Set("t", "get")
	u.RawQuery = restpost.Encode()
	resturl := fmt.Sprintf("%v", u)
	s, err := sabnzbd.New(sabnzbd.Addr(settings["sabsite"]), sabnzbd.ApikeyAuth(settings["sabkey"]))
	if err != nil {
		log.Fatalf("couldn't create sabnzbd: %s", err.Error())
		return "Error!"
	}
	auth, err := s.Auth()
	if err != nil {
		log.Fatalf("couldn't get auth type: %s", err.Error())
		return "Error!"
	}
	if auth != "apikey" {
		log.Fatalf("sabnzbd instance must be using apikey authentication")
		return "Error!"
	}
	_, err = s.AddURL(sabnzbd.AddNzbUrl(resturl))
	if err != nil {
		log.Fatalf("failed to upload nzb: %s", err.Error())
		return "Error!"
	}
	return "Downloading..."
}
