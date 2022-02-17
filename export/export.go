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

func Export(collections map[string][]user.User) error {
	var records []Records
	for i, j := range collections {
		tempRecord := Records{Index: i, Records: j, Total: len(j)}
		records = append(records, tempRecord)
	}
	err := WriteRecord(records)
	if err != nil {
		return fmt.Errorf("write record error: %v", err)
	}
	return nil
}

func WriteRecord(records []Records) error {
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
