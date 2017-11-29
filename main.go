package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

var Settings = getSettings()

func main() {

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var model = NewPersonModel(nil)
	var qmlBridge = NewQmlBridge(nil)
	qmlBridge.ConnectSendToGo(func(data string) string {
		return uploadNZBtoClient(data)
	})
	qmlBridge.ConnectResetList(func(model *PersonModel, search string) {
		go RefreshList(model, search)
	})

	var app = qml.NewQQmlApplicationEngine(nil)
	app.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	app.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	app.RootContext().SetContextProperty("PersonModel", model)

	gui.QGuiApplication_Exec()
}
