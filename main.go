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

	qmlBridge.Init()

	go LoopLoadQueue(queueModel)

	var app = qml.NewQQmlApplicationEngine(nil)
	app.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	app.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	app.RootContext().SetContextProperty("SearchModel", searchModel)
	app.RootContext().SetContextProperty("QueueModel", queueModel)

	go SendSettingsToQml(qmlBridge)

	gui.QGuiApplication_Exec()
}
