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

	_ func(data string) string                      `slot:"sendToGo"`
	_ func(searchModel *SearchModel, search string) `slot:"resetList"`
	_ func(queueModel *QueueModel)                  `slot:"queueList"`
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
