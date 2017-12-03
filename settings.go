package main

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"github.com/syndtr/goleveldb/leveldb"
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
		log.Infof("Config file not found...%s", err.Error())
	}
	settingsFile := viper.GetStringMapString("settings")
	db, err := leveldb.OpenFile("my.db", nil)
	if err != nil {
		log.Infof("Error opening DB...%s", err.Error())
	}
	defer db.Close()
	if len(settingsFile) > 0 {
		batch := new(leveldb.Batch)
		for i, x := range settingsFile {
			batch.Put([]byte(i), []byte(x))
		}
		err = db.Write(batch, nil)
		if err != nil {
			log.Infof("Error writing settings batch to DB...%s", err.Error())
		}
	}
	settingsDB := db.NewIterator(nil, nil)
	var settings map[string]string
	for settingsDB.Next() {
		settings[string(settingsDB.Key())] = string(settingsDB.Value())
	}
	settingsDB.Release()
	err = settingsDB.Error()
	if err != nil {
		log.Infof("Error reading settings from the DB...%s", err.Error())
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
	db, err := leveldb.OpenFile("my.db", nil)
	if err != nil {
		log.Infof("Error opening DB...%s", err.Error())
	}
	defer db.Close()
	if len(settingsFile) > 0 {
		batch := new(leveldb.Batch)
		for i, x := range settingsFile {
			batch.Put([]byte(i), []byte(x))
		}
		err = db.Write(batch, nil)
		if err != nil {
			log.Infof("Error writing settings batch to DB...%s", err.Error())
		}
	}
}
