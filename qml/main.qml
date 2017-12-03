//source: http://doc.qt.io/qt-5/qtquickcontrols2-material.html

import QtQml 2.2
import QtQuick 2.7
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
                        model: SearchModel
                        delegate: Component {
                            Item {
                                id: itemBox
                                width: parent.width
                                height: 20
                                Column {
                                    anchors.left: parent.left
                                    Text {
                                        text: date
                                        color: (index == listView.currentIndex) ? "lime" : "green"
                                    }
                                }
                                Column {
                                    anchors.left: parent.left
                                    leftPadding: 3
                                    Text {
                                        text: new Array(Math.round(date.length * 1.8)).join(' ') + description
                                        color: (index == listView.currentIndex) ? "white" : "steelblue"
                                    }
                                }
                                Column {
                                    anchors.right: parent.right
                                    rightPadding: 1
                                    Text {
                                        id: noticeText
                                        text: ""
                                        color: (index == listView.currentIndex) ? "lime" : "green"
                                    }
                                }

                                MouseArea {
                                    anchors.fill: parent
                                    acceptedButtons: Qt.LeftButton
                                    onClicked: {
                                        noticeText.text = QmlBridge.sendToGo(id);
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
                            onAccepted: QmlBridge.resetList(SearchModel, searchInput.text);
                        }

                        Rectangle {
                            anchors.right: parent.right
                            width: 48
                            radius: 8
                            color: "steelblue"
                            anchors.rightMargin: -2
                            anchors.bottom: parent.bottom
                            anchors.bottomMargin: -2
                            anchors.top: parent.top
                            anchors.topMargin: -2
                            z: 1
                            Text {
                                id: searchButtonText
                                text: " Search"
                                z: 2
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
                                    QmlBridge.resetList(SearchModel, searchInput.text);
                                }
                            }
                        }
                    }

            }
            Pane {
                width: swipeView.width
                height: swipeView.height
                ListView {
                    id: queueView
                    anchors.fill: parent
                    model: QueueModel
                    delegate: Component {
                        Item {
                            width: parent.width
                            height: 30
                            Column {
                                anchors.left: parent.left
                                Text {
                                    id: statusText
                                    text: itemStatus
                                    color: "lime"
                                }
                                Text {
                                    id: nameText
                                    text: name
                                    color: "steelblue"
                                }
                            }
                            Column {
                                anchors.right: parent.right
                                Text {
                                    id: remainingText
                                    text: remaining
                                    color: "lime"
                                }
                                Text {
                                    id: sizeText
                                    text: size
                                    color: "steelblue"
                                }
                            }
                            MouseArea {
                                anchors.fill: parent
                                width: parent.width
                                height: parent.height
                                onClicked: Qt.openUrlExternally(storage)
                            }
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
