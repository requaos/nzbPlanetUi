package main

import (
	"fmt"

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
		fmt.Println("Config file not found...")
		panic(err)
	}
	settings := viper.GetStringMapString("settings")
	// for i, x := range settings {
	// 	fmt.Printf("Key/Value: %s/%s", i, x)
	// }
	return settings
}
