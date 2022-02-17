package export

import (
	"awesomeProject2/user"
	"encoding/json"
	"fmt"
	"os"
)

type Records struct {
	Index   string      `json:"index"`
	Records []user.User `json:"records"`
	Total   int         `json:"total_records"`
}

// IExport interface
type IWrite interface {
	WriteRecord(records []Records) error
}

// Write structure
type Write struct {
}

func Export(collections map[string][]user.User, write IWrite) error {
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

func (w *Write) WriteRecord(records []Records) error {
	for _, record := range records {
		file, err := json.Marshal(record)
		if err != nil {
			return fmt.Errorf("marshal error: %v", err)
		}
		name := record.Index + ".json"
		err = os.WriteFile(name, file, 0644)
		if err != nil {
			return fmt.Errorf("cannot write the file: %v", err)
		}
	}
	return nil
}
