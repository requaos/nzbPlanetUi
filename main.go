package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

var Settings = getSettings()

func main() {
	gui.NewQGuiApplication(len(os.Args), os.Args)

	var view = quick.NewQQuickView(nil)

	var model = NewPersonModel(nil)
	var qmlBridge = NewQmlBridge(nil)
	qmlBridge.ConnectSendToGo(func(data string) string {
		return uploadNZBtoClient(data)
	})
	qmlBridge.ConnectResetList(func(model *PersonModel, search string) {
		go RefreshList(model, search)
	})

	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	view.RootContext().SetContextProperty("PersonModel", model)

	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.SetTitle("sabNZB Search")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetHeight(800)
	view.SetMinimumHeight(100)
	view.SetWidth(600)
	view.SetMaximumWidth(1024)
	view.SetMinimumWidth(400)
	view.Show()

	gui.QGuiApplication_Exec()
}

func RefreshList(model *PersonModel, search string) {
	model.BeginResetModel()
	model.SetPeople([]*Person{NewPerson(nil)})
	model.EndResetModel()
	model.RemovePerson(0)

	searchList := SearchForHSnzbs(search, Settings)

	//add person
	for i := 0; i < len(searchList.Channel.Item); i++ {
		publishedDate, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", searchList.Channel.Item[i].PubDate)
		if err != nil {
			log.Error("Error parsing time/date stamp on item")
		}
		var p = NewPerson(nil)
		p.SetDescription(searchList.Channel.Item[i].Title)
		p.SetDate(publishedDate.Format("01/02/2006"))
		p.SetId(searchList.Channel.Item[i].GUID[34:])
		model.AddPerson(p)
	}
}

type SearchResponse struct {
	Attributes struct {
		Version string `json:"version"`
	} `json:"@attributes"`
	Channel struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Link        string `json:"link"`
		Language    string `json:"language"`
		WebMaster   string `json:"webMaster"`
		Category    struct {
		} `json:"category"`
		Image struct {
			URL         string `json:"url"`
			Title       string `json:"title"`
			Link        string `json:"link"`
			Description string `json:"description"`
		} `json:"image"`
		Response struct {
			Attributes struct {
				Offset string `json:"offset"`
				Total  string `json:"total"`
			} `json:"@attributes"`
		} `json:"response"`
		Item []struct {
			Title       string `json:"title"`
			GUID        string `json:"guid"`
			Link        string `json:"link"`
			Comments    string `json:"comments"`
			PubDate     string `json:"pubDate"`
			Category    string `json:"category"`
			Description string `json:"description"`
			Enclosure   struct {
				Attributes struct {
					URL    string `json:"url"`
					Length string `json:"length"`
					Type   string `json:"type"`
				} `json:"@attributes"`
			} `json:"enclosure"`
			Attr []struct {
				Attributes struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"@attributes"`
			} `json:"attr"`
		} `json:"item"`
	} `json:"channel"`
}

//SearchForHSnzbs is the initial athentication call
func SearchForHSnzbs(search string, settings map[string]string) SearchResponse {
	client := &http.Client{}
	u, _ := url.ParseRequestURI(settings["nzbsite"])
	u.Path = "/api"
	restpost := u.Query()
	restpost.Add("apikey", settings["nzbkey"])
	restpost.Set("o", "json")
	restpost.Add("q", search)
	restpost.Set("t", "search")
	u.RawQuery = restpost.Encode()
	resturl := fmt.Sprintf("%v", u)
	r, _ := http.NewRequest("GET", resturl, nil)
	fmt.Println(r)
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var b SearchResponse
	err = json.Unmarshal(body, &b)
	if err != nil {
		panic(err)
	}
	return b
}

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
