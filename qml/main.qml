import QtQuick 2.5


Rectangle {
    color: 'black'
    ListView {
        id: listView
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
                        color: "green"
                    }               
                }
                Column {
                    Text {
                        text: new Array(Math.round(date.length * 1.8)).join(' ') + description
                        color: "steelblue"
                    }   
                }
                MouseArea {
                    anchors.fill: parent
                    acceptedButtons: Qt.LeftButton | Qt.RightButton
                    onClicked: {
                        if (mouse.button == Qt.RightButton)
                            listView.
                            QmlBridge.resetList(PersonModel);
                        else
                            noticeText.text = date + new Array(Math.round((date.length + description.length) * 1.65)).join(' ') + QmlBridge.sendToGo(id);
                    }
                    onEntered: listView.currentIndex = index;
                    hoverEnabled: true
                }                             
            }       
        }
        highlight: Rectangle {
            color: "#250025"
        }
        focus: true
    }
}
