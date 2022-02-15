package export

import (
	api "awesomeProject2/service"
	"encoding/json"
	"io/ioutil"
)

type Records struct {
	Index   string     `json:"index"`
	Records []api.User `json:"records"`
	Total   int        `json:"total_records"`
}

func Export(collections map[string][]api.User) {
	var records []Records
	for i, j := range collections {
		tempRecord := Records{Index: i, Records: j, Total: len(j)}
		records = append(records, tempRecord)
		file, _ := json.MarshalIndent(tempRecord, "", " ")
		var name string
		name = i + ".json"
		_ = ioutil.WriteFile(name, file, 0644)
	}
}
