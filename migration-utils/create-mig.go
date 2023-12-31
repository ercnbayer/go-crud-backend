package migrationutils

import (
	"fmt"
	"go-backend/logger"
	"os"
	"text/template"
	"time"
)

type Migration struct {
	Name           string
	Timestamp      string
	TableName      string
	SuccessInitLog string
}

func Init(structName string) {

	timestamp := time.Now().Format("20060102150405")
	templateFile := "migration-utils/migration.tmpl"
	fileName := fmt.Sprintf("migration/%s.go", timestamp)

	migrationFile := []Migration{{Name: "person", Timestamp: timestamp, TableName: "persons", SuccessInitLog: "Table Init"}}

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		//panic(err)
		logger.Error(err.Error())
	}

	File, err := os.Create(fileName)

	if err != nil {
		logger.Error(err)
	}

	err = tmpl.Execute(File, migrationFile)

	if err != nil {

		File.Close()
		logger.Error()

	}

	File.Close()
	// end main

}
