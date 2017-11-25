import QtQuick 2.0

ListView {
    id: listView
    anchors.fill: parent
    model: PersonModel

    delegate: Component {
        Item {
            id: itemBox
            width: parent.width
            height: 40
            Column {
                Text {
                    text: date
                    color: "green"
                    }
                Text {
                    text: description
                    color: "lightsteelblue"
                    }
            }
            MouseArea {
                anchors.fill: parent
                acceptedButtons: Qt.LeftButton | Qt.RightButton
                onClicked: {
                    if (mouse.button == Qt.RightButton)
                        parent.opacity = !parent.opacity;
                    else
                        listView.currentIndex = index;
                }
                onEntered: listView.currentIndex = index;
                hoverEnabled: true
            }
        }
    }
    highlight: Rectangle {
        color: 'darkviolet'
    }
    background: Rectangle {
        color: 'black'
    }
    focus: true
    onCurrentItemChanged: console.log(PersonModel.Id() + ' selected')
}
