//source: http://doc.qt.io/qt-5/qtquickcontrols2-material.html

import QtQuick 2.5
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import QtQuick.Controls.Universal 2.0

ApplicationWindow {
    visible: true
    title: "Search"
    height: 800
    width: 600
    footer: TabBar {
            id: tabBar
            width: parent.width
            currentIndex: swipeView.currentIndex
            Component.onCompleted: QmlBridge.resetList(PersonModel, searchInput.text);
            TabButton {
                text: qsTr("Search")
            }
            TabButton {
                text: qsTr("Downloads")
            }
            TabButton {
                text: qsTr("Settings")
            }
        }
    SwipeView{
            id: swipeView
            anchors.fill: parent
            currentIndex: tabBar.currentIndex
            Pane {
                width: swipeView.width
                height: swipeView.height
                    ListView {
                        id: listView
                        preferredHighlightBegin: 0
                        anchors.topMargin: 3
                        anchors.fill: parent
                        model: PersonModel
                        delegate: Component {
                            Item {
                                id: itemBox
                                width: parent.width
                                height: 20
                                Column {
                                    Text {
                                        id: noticeText
                                        text: date
                                        color: (index == listView.currentIndex) ? "lime" : "green"
                                    }
                                }
                                Column {
                                    Text {
                                        text: new Array(Math.round(date.length * 1.8)).join(' ') + description
                                        color: (index == listView.currentIndex) ? "white" : "steelblue"
                                    }
                                }
                                MouseArea {
                                    anchors.fill: parent
                                    acceptedButtons: Qt.LeftButton
                                    onClicked: {
                                        noticeText.text = date + new Array(Math.round((date.length + description.length) * 1.65)).join(' ') + QmlBridge.sendToGo(id);
                                    }
                                    onEntered: listView.currentIndex = index;
                                    hoverEnabled: true
                                }
                            }
                        }
                        highlight: Rectangle {
                            color: "steelblue"
                        }
                        focus: true
                    }

                    Rectangle {
                        id: rectangle
                        anchors.top: parent.top
                        anchors.topMargin: 1
                        anchors.right: parent.right
                        anchors.rightMargin: -3
                        width: 217
                        height: 22
                        color: "#FFFFFF"
                        radius: 8
                        border.color: "steelblue"
                        border.width: 1
                        clip: true
                        TextInput {
                            id: searchInput
                            width: 158
                            color: "steelblue"
                            text: "[HorribleSubs] 720p"
                            anchors.top: parent.top
                            anchors.topMargin: 2
                            anchors.right: parent.right
                            anchors.bottom: parent.bottom
                            anchors.left: parent.left
                            transformOrigin: Item.Center
                            anchors.rightMargin: 48
                            anchors.leftMargin: 5
                            anchors.bottomMargin: 1
                            z: 2
                            cursorVisible: false
                            font.family: "Arial"
                            selectionColor: "steelblue"
                            font.pixelSize: 12
                            onAccepted: QmlBridge.resetList(PersonModel, searchInput.text);
                        }

                        Rectangle {
                            anchors.right: parent.right
                            width: 48
                            color: "steelblue"
                            anchors.rightMargin: 1
                            anchors.bottom: parent.bottom
                            anchors.bottomMargin: 1
                            anchors.top: parent.top
                            anchors.topMargin: 1
                            z: 1
                            Text {
                                id: searchButtonText
                                text: "Search"
                                z: 2
                                anchors.fill: parent
                                verticalAlignment: Text.AlignVCenter
                                font.bold: false
                                horizontalAlignment: Text.AlignHCenter
                                color: "white"
                            }
                            MouseArea {
                                z: 2
                                anchors.fill: parent
                                acceptedButtons: Qt.LeftButton
                                onClicked: {
                                    QmlBridge.resetList(PersonModel, searchInput.text);
                                }
                            }
                        }
                    }

            }
            Pane {
                width: swipeView.width
                height: swipeView.height
                Rectangle {
                    id: settingsRectangle
                    anchors.top: parent.top
                    anchors.topMargin: 1
                    anchors.right: parent.right
                    anchors.rightMargin: -3
                    width: 217
                    height: 22
                    color: "#FFFFFF"
                    radius: 8
                    border.color: "steelblue"
                    border.width: 1
                    clip: true
                    TextInput {
                        id: nzbPlanetKey
                        width: 158
                        color: "steelblue"
                        text: ""
                        anchors.top: parent.top
                        anchors.topMargin: 2
                        anchors.right: parent.right
                        anchors.bottom: parent.bottom
                        anchors.left: parent.left
                        transformOrigin: Item.Center
                        anchors.rightMargin: 48
                        anchors.leftMargin: 5
                        anchors.bottomMargin: 1
                        z: 2
                        cursorVisible: false
                        font.family: "Arial"
                        selectionColor: "steelblue"
                        font.pixelSize: 12
                    }
                    TextInput {
                        id: sabnzbKey
                        width: 158
                        color: "steelblue"
                        text: ""
                        anchors.top: parent.top
                        anchors.topMargin: 20
                        anchors.right: parent.right
                        anchors.bottom: parent.bottom
                        anchors.left: parent.left
                        transformOrigin: Item.Center
                        anchors.rightMargin: 48
                        anchors.leftMargin: 5
                        anchors.bottomMargin: 1
                        z: 2
                        cursorVisible: false
                        font.family: "Arial"
                        selectionColor: "steelblue"
                        font.pixelSize: 12
                    }
                }
                Rectangle {
                    anchors.right: parent.right
                    width: 48
                    color: "steelblue"
                    anchors.rightMargin: 1
                    anchors.bottom: parent.bottom
                    anchors.bottomMargin: 1
                    anchors.top: parent.top
                    anchors.topMargin: 1
                    z: 1
                    Text {
                        id: saveButtonText
                        text: "Save"
                        z: 2
                        anchors.fill: parent
                        verticalAlignment: Text.AlignVCenter
                        font.bold: false
                        horizontalAlignment: Text.AlignHCenter
                        color: "white"
                    }
                    MouseArea {
                        z: 2
                        anchors.fill: parent
                        acceptedButtons: Qt.LeftButton
                        onClicked: {
                            console.log("Save Clicked");
                        }
                    }
                }
                Rectangle {
                    anchors.right: parent.right
                    width: 48
                    color: "crimson"
                    anchors.rightMargin: 1
                    anchors.bottom: parent.bottom
                    anchors.bottomMargin: 1
                    anchors.top: parent.top
                    anchors.topMargin: 1
                    z: 1
                    Text {
                        id: cancelButtonText
                        text: "Cancel"
                        z: 2
                        anchors.fill: parent
                        verticalAlignment: Text.AlignVCenter
                        font.bold: false
                        horizontalAlignment: Text.AlignHCenter
                        color: "white"
                    }
                    MouseArea {
                        z: 2
                        anchors.fill: parent
                        acceptedButtons: Qt.LeftButton
                        onClicked: {
                            console.log("Cancel Clicked");
                        }
                    }
                }
            }
            Pane {
                width: swipeView.width
                height: swipeView.height
                Column {
                    anchors.centerIn: parent
                    RadioButton { text: qsTr("Male") }
                    RadioButton { text: qsTr("Female");  checked: true }
                    RadioButton { text: qsTr("Other") }
                }
            }
        }
    }
