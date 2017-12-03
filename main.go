package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

var (
	Settings   = getSettings()
	SABnzbd    = SABnzbdSession()
	startingUp = true
)

func main() {
	gui.NewQGuiApplication(len(os.Args), os.Args)

	var searchModel = NewSearchModel(nil)
	var queueModel = NewQueueModel(nil)
	var qmlBridge = NewQmlBridge(nil)
	qmlBridge.ConnectSendToGo(func(data string) string {
		return uploadNZBtoClient(data)
	})
	qmlBridge.ConnectResetList(func(searchModel *SearchModel, search string) {
		go RefreshList(searchModel, search)
	})
	qmlBridge.ConnectQueueList(func(queueModel *QueueModel) {
		go GetQueueDetails(queueModel)
	})
	go LoopLoadQueue(queueModel)

	var app = qml.NewQQmlApplicationEngine(nil)
	app.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	app.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	app.RootContext().SetContextProperty("SearchModel", searchModel)
	app.RootContext().SetContextProperty("QueueModel", queueModel)

	gui.QGuiApplication_Exec()
}
