package main

import (
	"github.com/therecipe/qt/core"
)

type QmlBridge struct {
	core.QObject

	_ func(data string) `signal:"nzbSite"`
	_ func(data string) `signal:"nzbKey"`
	_ func(data string) `signal:"sabSite"`
	_ func(data string) `signal:"sabKey"`

	_ func(data string) string                                           `slot:"sendToGo"`
	_ func(searchModel *SearchModel, search string)                      `slot:"resetList"`
	_ func(queueModel *QueueModel)                                       `slot:"queueList"`
	_ func(nzbSite string, nzbKey string, sabSite string, sabKey string) `slot:"saveSettings"`
}

func (qmlBridge *QmlBridge) Init() {
	qmlBridge.ConnectSendToGo(func(data string) string {
		return uploadNZBtoClient(data)
	})
	qmlBridge.ConnectResetList(func(searchModel *SearchModel, search string) {
		go RefreshList(searchModel, search)
	})
	qmlBridge.ConnectQueueList(func(queueModel *QueueModel) {
		go GetQueueDetails(queueModel)
	})
	qmlBridge.ConnectSaveSettings(func(nzbSite string, nzbKey string, sabSite string, sabKey string) {
		go setSettings(nzbSite, nzbKey, sabSite, sabKey)
	})
}

func SendSettingsToQml(q *QmlBridge) {
	q.NzbSite(Settings["nzbsite"])
	q.NzbKey(Settings["nzbkey"])
	q.SabSite(Settings["sabsite"])
	q.SabKey(Settings["sabkey"])
}
