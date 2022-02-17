package export

import (
	"awesomeProject2/user"
	"fmt"
	"testing"
)

type MockIWriteErr interface {
	WriteRecord(records []Records) error
}

type MockWriteErr struct {
}

func (m MockWriteErr) WriteRecord(records []Records) error {
	return fmt.Errorf("cannot create record")
}

func TestExportRecordsError(t *testing.T) {
	var w = &MockWriteErr{}
	user1 := user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	user2 := user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	user3 := user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	collections := make(map[string][]user.User)
	collections["T"] = append(collections["T"], user1, user2, user3)
	err := Export(collections, w)
	if err == nil {
		t.Errorf("Read dates error")
	}
}

type MockIWriteSuccess interface {
	WriteRecord(records []Records) error
}

type MockWriteSuccess struct {
}

func (m MockWriteSuccess) WriteRecord(records []Records) error {
	return nil
}

func TestExportRecordsSuccess(t *testing.T) {
	var w = &MockWriteSuccess{}
	user1 := user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	user2 := user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	user3 := user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	collections := make(map[string][]user.User)
	collections["T"] = append(collections["T"], user1, user2, user3)
	err := Export(collections, w)
	if err != nil {
		t.Errorf("Read dates error")
	}
}

// FileWriter implements is an abstraction of ioutil.WriterFile
type FileWriter struct {
}

type FileMode struct {
}

// WriteFile implements the Writer interface that's been created so that ioutil.WriteFile can be mocked
func (w FileWriter) WriteFile(name string, data []byte, perm FileMode) error {
	return fmt.Errorf("cannot write in file")
}
func TestWriteRecordErr(t *testing.T) {
	var w = &Write{}
	user1 := user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	user2 := user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	var users []user.User
	var records []Records
	users = append(users, user1, user2)
	tempRecord := Records{Index: "T", Records: users, Total: len(users)}
	records = append(records, tempRecord)

	err := w.WriteRecord(records)
	if err == nil {
		t.Errorf("invalide URL")
	}
}
