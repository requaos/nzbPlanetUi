package main

import (
	"github.com/therecipe/qt/core"
)

const (
	Size = int(core.Qt__UserRole) + 1<<iota
	Name
	Remaining
	ItemStatus
	Storage
)

type QueueModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Queue                 `property:"items"`

	_ func(*Queue)                                                     `slot:"addQueue"`
	_ func(row int, size, name, remaining, itemStatus, storage string) `slot:"editQueue"`
	_ func(row int)                                                    `slot:"removeQueue"`
}

type Queue struct {
	core.QObject

	_ string `property:"name"`
	_ string `property:"size"`
	_ string `property:"remaining"`
	_ string `property:"itemStatus"`
	_ string `property:"storage"`
}

func init() {
	Queue_QRegisterMetaType()
}

func (m *QueueModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Name:       core.NewQByteArray2("name", len("name")),
		Size:       core.NewQByteArray2("size", len("size")),
		Remaining:  core.NewQByteArray2("remaining", len("remaining")),
		ItemStatus: core.NewQByteArray2("itemStatus", len("itemStatus")),
		Storage:    core.NewQByteArray2("storage", len("storage")),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddQueue(m.addQueue)
	m.ConnectEditQueue(m.editQueue)
	m.ConnectRemoveQueue(m.removeQueue)
}

func (m *QueueModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Items()) {
		return core.NewQVariant()
	}

	var p = m.Items()[index.Row()]

	switch role {
	case Name:
		{
			return core.NewQVariant14(p.Name())
		}

	case Size:
		{
			return core.NewQVariant14(p.Size())
		}

	case Remaining:
		{
			return core.NewQVariant14(p.Remaining())
		}

	case ItemStatus:
		{
			return core.NewQVariant14(p.ItemStatus())
		}

	case Storage:
		{
			return core.NewQVariant14(p.Storage())
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *QueueModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Items())
}

func (m *QueueModel) columnCount(parent *core.QModelIndex) int {
	return 4
}

func (m *QueueModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *QueueModel) addQueue(p *Queue) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Items()), len(m.Items()))
	m.SetItems(append(m.Items(), p))
	m.EndInsertRows()
}

func (m *QueueModel) editQueue(row int, size string, name string, remaining string, itemStatus string, storage string) {
	var p = m.Items()[row]

	if size != "" {
		p.SetSize(size)
	}

	if name != "" {
		p.SetName(name)
	}

	if remaining != "" {
		p.SetRemaining(remaining)
	}

	if itemStatus != "" {
		p.SetItemStatus(itemStatus)
	}

	if storage != "" {
		p.SetStorage(storage)
	}

	var pIndex = m.Index(row, 0, core.NewQModelIndex())
	m.DataChanged(pIndex, pIndex, []int{Name, Size, Remaining, ItemStatus, Storage})
}

func (m *QueueModel) removeQueue(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetItems(append(m.Items()[:row], m.Items()[row+1:]...))
	m.EndRemoveRows()
}
