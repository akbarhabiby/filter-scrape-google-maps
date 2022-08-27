package helpers

import (
	"encoding/json"
	"fmt"
	"os"
)

func ImportJSON(fileName string, result interface{}) {
	fmt.Printf("Importing %s ...\n", fileName)
	timeLog := Timelog(fmt.Sprintf("%s %s", "Import", fileName))
	bt, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bt, &result)
	if err != nil {
		panic(err)
	}
	timeLog()
}

func ExportJSON(fileName string, jsonData interface{}) {
	fmt.Printf("Exporting %s ...\n", fileName)
	timeLog := Timelog(fmt.Sprintf("%s %s", "Export", fileName))
	bt, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fileName, bt, 0644)
	if err != nil {
		panic(err)
	}
	timeLog()
}
