package main

import (
	"github.com/labstack/gommon/log"
	sabnzbd "github.com/michaeltrobinson/go-sabnzbd"
)

func SABnzbdSession() *sabnzbd.Sabnzbd {
	if s, ok := Settings["sabsite"]; !ok || s == "" {
		return nil
	}
	if k, ok := Settings["sabkey"]; !ok || k == "" {
		return nil
	}
	s, err := sabnzbd.New(sabnzbd.Addr(Settings["sabsite"]), sabnzbd.ApikeyAuth(Settings["sabkey"]))
	if err != nil {
		log.Fatalf("couldn't create sabnzbd: %s", err.Error())
		return nil
	}
	auth, err := s.Auth()
	if err != nil {
		log.Fatalf("couldn't get auth type: %s", err.Error())
		return nil
	}
	if auth != "apikey" {
		log.Fatalf("sabnzbd instance must be using apikey authentication")
		return nil
	}
	return s

}
