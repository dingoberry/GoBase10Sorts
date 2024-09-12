package patterns

import (
	"container/list"
	"strings"
)

type DbOperation interface {
	Query() string
	Insert(data string)
	Delete(data string)
}

type MySql struct {
	DbOperation
	dataStore list.List
}

func (m MySql) Query() string {
	sb := strings.Builder{}
	for e := m.dataStore.Front(); e != nil; e = e.Next() {
		sb.WriteString(e.Value.(string))
	}
	return sb.String()
}
func (m MySql) Insert(data string) {
	m.dataStore.PushBack(data)
}
func (m MySql) Delete(data string) {
	for e := m.dataStore.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == data {
			m.dataStore.Remove(e)
			break
		}
	}
}

type OracleSql struct {
	Operation DbOperation
	dataStore map[string]string
}

func (m OracleSql) Query() string {
	sb := strings.Builder{}
	for s := range m.dataStore {
		sb.WriteString(s)
	}
	return sb.String()
}
func (m OracleSql) Insert(data string) {
	m.dataStore[data] = data
}
func (m OracleSql) Delete(data string) {
	m.Operation.Delete(data)
}

type bridgeDb struct {
	DbOperation
	operation DbOperation
}

func (m bridgeDb) Query() string {
	return m.operation.Query()
}
func (m bridgeDb) Insert(data string) {
	m.operation.Insert(data)
}
func (m bridgeDb) Delete(data string) {
	m.operation.Delete(data)
}

func CreateBridgeDb(op DbOperation) bridgeDb {
	return bridgeDb{
		operation: op,
	}
}
