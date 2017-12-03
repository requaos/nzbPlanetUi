package main

import (
	"fmt"
	"time"

	"github.com/asdine/storm"
	"github.com/boltdb/bolt"
	"github.com/spf13/viper"
)

func getSettings() map[string]string {
	// example file: secrets.toml
	// [settings]
	// nzbSite = "https://api.nzbplanet.net"
	// nzbKey = "157b4974da310d1f834644fe93298171"
	// sabSite = "localhost:8080"
	// sabKey = "6a1c4e43be73e58c2c2617043c72b8de"
	viper.SetConfigName("secrets")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Config file not found...%s", err.Error())
	}
	settingsFile := viper.GetStringMapString("settings")
	db, err := storm.Open("my.db", storm.BoltOptions(0600, &bolt.Options{Timeout: 1 * time.Second}))
	if err != nil {
		fmt.Printf("Error opening DB...%s", err.Error())
	}
	defer db.Close()
	if len(settingsFile) > 0 {
		db.Set("settings", "set", settingsFile)
	}
	var settings map[string]string
	db.Get("settings", "set", settings)
	for i, x := range settings {
		fmt.Printf("Key: % 9s Value: %s", i, x)
	}
	return settings
}

func setSettings(nzbSite string, nzbKey string, sabSite string, sabKey string) {
	settingsFile := map[string]string{
		"nzbsite": nzbSite,
		"nzbkey":  nzbKey,
		"sabsite": sabSite,
		"sabkey":  sabKey,
	}
	db, err := storm.Open("my.db", storm.BoltOptions(0600, &bolt.Options{Timeout: 1 * time.Second}))
	if err != nil {
		fmt.Printf("Error opening DB...%s", err.Error())
	}
	defer db.Close()
	if len(settingsFile) > 0 {
		db.Set("settings", "set", settingsFile)
	}
	Settings = settingsFile
}
