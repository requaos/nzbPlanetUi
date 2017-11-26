package main

import (
	"github.com/therecipe/qt/core"
)

const (
	Date = int(core.Qt__UserRole) + 1<<iota
	Description
	ID
)

type PersonModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Person                `property:"people"`

	_ func(*Person)                               `slot:"addPerson"`
	_ func(row int, date, description, id string) `slot:"editPerson"`
	_ func(row int)                               `slot:"removePerson"`
}

type Person struct {
	core.QObject

	_ string `property:"description"`
	_ string `property:"date"`
	_ string `property:"id"`
}

func init() {
	Person_QRegisterMetaType()
}

func (m *PersonModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Description: core.NewQByteArray2("description", len("description")),
		Date:        core.NewQByteArray2("date", len("date")),
		ID:          core.NewQByteArray2("id", len("id")),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddPerson(m.addPerson)
	m.ConnectEditPerson(m.editPerson)
	m.ConnectRemovePerson(m.removePerson)
}

func (m *PersonModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.People()) {
		return core.NewQVariant()
	}

	var p = m.People()[index.Row()]

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

func (m *PersonModel) rowCount(parent *core.QModelIndex) int {
	return len(m.People())
}

func (m *PersonModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *PersonModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *PersonModel) addPerson(p *Person) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.People()), len(m.People()))
	m.SetPeople(append(m.People(), p))
	m.EndInsertRows()
}

func (m *PersonModel) editPerson(row int, date string, description string, id string) {
	var p = m.People()[row]

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

func (m *PersonModel) removePerson(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetPeople(append(m.People()[:row], m.People()[row+1:]...))
	m.EndRemoveRows()
}
