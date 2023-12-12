package migrationutils

import (
	"fmt"
	"os"
	"time"
)

// Make this a seperate function. And it has parameter "name". you are gonna use this name parameter while your new migration
func Init() {

	timestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("%s.go", timestamp)

	//Make all text seperate to seperate .txt file. And use approach like this instead of sprintf https://www.digitalocean.com/community/tutorials/how-to-use-templates-in-go
	content := fmt.Sprintf(`package migration

	import (
		"go-backend/db"
		"go-backend/logger"
	)
	
	//What will happen if i dont want to create migration which dont be PERSON :D. Person should be a variable. And we could define that while we creating migration
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
