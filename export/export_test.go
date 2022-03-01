package export

import (
	"fmt"
	"os"
	"testing"

	"service/user"
)

type MockWriteErr struct {
}

func (m MockWriteErr) WriteRecord(records []Records) error {
	return fmt.Errorf("cannot create record")
}

func (m MockWriteErr) WriteFile(name string, data []byte, perm os.FileMode) error {
	return fmt.Errorf("cannot write in file")
}

var user1 = user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
var user2 = user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
var user3 = user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
var user4 = user.User{FirstName: "Keith", LastName: "Hilpert", Email: "Keith.Hilpert@sammy.name", Address: "230 Kenna Harbor", Created: "March 11, 2019", Balance: "$8,419.47"}
var user5 = user.User{FirstName: "Brendan", LastName: "Mosciski", Email: "Keith.Hilpert@sammy.name", Address: "230 Kenna Harbor", Created: "March 11, 2019", Balance: "$8,419.47"}

func TestExportRecordsError(t *testing.T) {
	var w = &MockWriteErr{}
	collections := make(map[string][]user.User)
	collections["T"] = append(collections["T"], user1, user2, user3, user4, user5)
	err := Execute(collections, w)
	if err == nil {
		t.Errorf("Read dates error")
	}
}

func TestWriteRecordError(t *testing.T) {
	var w = &MockWriteErr{}
	var users []user.User
	var records []Records
	users = append(users, user1, user2)
	tempRecord := Records{Index: "T", Records: users, Total: len(users)}
	records = append(records, tempRecord)

	err := w.WriteRecord(records)
	if err == nil {
		t.Errorf("write record error")
	}
}

type MockWriteSuccess struct {
}

func (m MockWriteSuccess) WriteRecord(records []Records) error {
	return nil
}

func (m MockWriteSuccess) WriteFile(name string, data []byte, perm os.FileMode) error {
	return nil
}

func TestExportRecordsSuccess(t *testing.T) {
	var w = &MockWriteSuccess{}
	collections := make(map[string][]user.User)
	collections["T"] = append(collections["T"], user1, user2, user3)
	err := Execute(collections, w)
	if err != nil {
		t.Errorf("Read dates error")
	}
}

func TestWriteRecordSuccess(t *testing.T) {
	var w = &MockWriteSuccess{}
	var users []user.User
	var records []Records
	users = append(users, user1, user2)
	tempRecord := Records{Index: "T", Records: users, Total: len(users)}
	records = append(records, tempRecord)

	err := w.WriteRecord(records)
	if err != nil {
		t.Errorf("write record error")
	}
}
