package migrationutils

import (
	"fmt"
	"os"
	"time"
)

func Init() {

	timestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("%s.go", timestamp)
	content := fmt.Sprintf(`package migration

	import (
		"go-backend/db"
		"go-backend/logger"
	)
	
	type  Person%s struct {

	}
	
	
	func (table Person%s) TableName() string {
		return "persons"
	}
	func PersonUp%s() error {

		return nil
	}
	func PersonDown%s() error {
		
		return nil
	}
	
	func init() {
	
		Migrations_Arr = append(Migrations_Arr, Migration{
			Name:   "Person%s",
			UpFn:   PersonUp%s,
			DownFn: PersonDown%s,
		})
		logger.Info("TABLE INIT")
	
	}
	`, timestamp, timestamp, timestamp, timestamp, timestamp, timestamp, timestamp)
	err := os.WriteFile("./migration/"+fileName, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
}
