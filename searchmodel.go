package main

import (
	"github.com/therecipe/qt/core"
)

const (
	Date = int(core.Qt__UserRole) + 1<<iota
	Description
	ID
)

type SearchModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Search                `property:"rows"`

	_ func(*Search)                               `slot:"addSearch"`
	_ func(row int, date, description, id string) `slot:"editSearch"`
	_ func(row int)                               `slot:"removeSearch"`
}

type Search struct {
	core.QObject

	_ string `property:"description"`
	_ string `property:"date"`
	_ string `property:"id"`
}

func init() {
	Search_QRegisterMetaType()
}

func (m *SearchModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Description: core.NewQByteArray2("description", len("description")),
		Date:        core.NewQByteArray2("date", len("date")),
		ID:          core.NewQByteArray2("id", len("id")),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddSearch(m.addSearch)
	m.ConnectEditSearch(m.editSearch)
	m.ConnectRemoveSearch(m.removeSearch)
}

func (m *SearchModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Rows()) {
		return core.NewQVariant()
	}

	var p = m.Rows()[index.Row()]

	switch role {
	case Description:
		{
			return core.NewQVariant14(p.Description())
		}

	case Date:
		{
			return core.NewQVariant14(p.Date())
		}

	case ID:
		{
			return core.NewQVariant14(p.Id())
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *SearchModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Rows())
}

func (m *SearchModel) columnCount(parent *core.QModelIndex) int {
	return 3
}

func (m *SearchModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *SearchModel) addSearch(p *Search) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Rows()), len(m.Rows()))
	m.SetRows(append(m.Rows(), p))
	m.EndInsertRows()
}

func (m *SearchModel) editSearch(row int, date string, description string, id string) {
	var p = m.Rows()[row]

	if date != "" {
		p.SetDate(date)
	}

	if description != "" {
		p.SetDescription(description)
	}

	if id != "" {
		p.SetId(id)
	}

	var pIndex = m.Index(row, 0, core.NewQModelIndex())
	m.DataChanged(pIndex, pIndex, []int{Description, Date, ID})
}

func (m *SearchModel) removeSearch(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetRows(append(m.Rows()[:row], m.Rows()[row+1:]...))
	m.EndRemoveRows()
}
