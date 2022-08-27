package helpers

import (
	"encoding/json"
	"os"
)

func ImportJSON(fileName string, result interface{}) {
	bt, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bt, &result)
	if err != nil {
		panic(err)
	}
}

func ExportJSON(fileName string, jsonData interface{}) {
	bt, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fileName, bt, 0644)
	if err != nil {
		panic(err)
	}
}
