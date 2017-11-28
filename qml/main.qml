import QtQuick 2.5

Rectangle {
    gradient: Gradient {
        GradientStop {
            position: 0
            color: "steelblue"
        }

        GradientStop {
            position: 0.04
            color: "#ffffff"
        }

        GradientStop {
            position: 0.259
            color: "#ffffff"
        }

        GradientStop {
            position: 0.96
            color: "#ffffff"
        }

        GradientStop {
            position: 1
            color: "steelblue"
        }
    }
    border.width: 3
    border.color: "steelblue"
    ListView {
        id: listView
        preferredHighlightBegin: 0
        anchors.bottomMargin: 13
        anchors.topMargin: 13
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
        anchors.topMargin: 27
        anchors.right: parent.right
        anchors.rightMargin: 27
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
            Component.onCompleted: QmlBridge.resetList(PersonModel, searchInput.text);
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
