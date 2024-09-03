package core

import (
	"SystemUtilManager/database_integration"
	"fmt"
)

func GetTask(id int) []string {

	db := database_integration.ConnectToDB()
	defer db.Close()
	count := database_integration.GetNumOfTasksDB(db, id)
	var results []string
	var buffer string
	sqlRows := database_integration.GetTasksFromDB(db, id)
	sqlRows.Next()
	for i := 0; i < count; i++ {
		err := sqlRows.Scan(&buffer)
		if err != nil {
			fmt.Println("Error of scan in GetTask", err)
		}
		results = append(results, buffer)
		sqlRows.Next()
	}
	return results
}
func AddTask(id int, task string) {
	db := database_integration.ConnectToDB()
	database_integration.AddTask(db, id, task)
	db.Close()
}
