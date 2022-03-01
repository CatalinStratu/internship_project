package export

import (
	"encoding/json"
	"fmt"
	"os"

	"service/user"
)

//Records record structure
type Records struct {
	Index   string      `json:"index"`
	Records []user.User `json:"records"`
	Total   int         `json:"total_records"`
}

// Write interface
type Write interface {
	WriteRecord(records []Records) error
	WriteFile(name string, data []byte, perm os.FileMode) error
}

// WriteFile structure
type WriteFile struct{}

//WriteFile writes data to the named file, creating it if necessary.
func (w WriteFile) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}

//Execute collections in files, using records slice
func Execute(collections map[string][]user.User, write Write) error {
	var records []Records
	for i, j := range collections {
		tempRecord := Records{Index: i, Records: j, Total: len(j)}
		records = append(records, tempRecord)
	}
	err := write.WriteRecord(records)
	if err != nil {
		return fmt.Errorf("write record error: %v", err)
	}
	return nil
}

// WriteRecord Write record in file
func (w *WriteFile) WriteRecord(records []Records) error {
	for _, record := range records {
		file, err := encoding(record)
		if err != nil {
			return fmt.Errorf("marshal error: %w", err)
		}
		name := record.Index + ".json"
		err = w.WriteFile(name, file, 0644)
		if err != nil {
			return fmt.Errorf("cannot write the file: %w", err)
		}
	}
	return nil
}

func encoding(record Records) ([]byte, error) {
	file, err := json.Marshal(record)
	if err != nil {
		return nil, fmt.Errorf("marshal error: %w", err)
	}
	return file, nil
}
