package database_integration

import (
	"database/sql"
	"fmt"
)

func CreateTableAgentsTasks(db *sql.DB) {
	_, err := db.Query("CREATE TABLE agents_tasks (id integer, task varchar(128)); ")
	if err != nil {
		fmt.Println("Error of creating table agents_tasks", err)
	}
}

func AddTask(db *sql.DB, id int, task string) {
	_, err := db.Query("INSERT INTO agents_tasks VALUES ($1, $2);", id, task)
	if err != nil {
		fmt.Println("Error of adding task", err)
	}
}
func GetNumOfTasksDB(db *sql.DB, id int) int {
	rows, err := db.Query("SELECT COUNT(task) FROM agents_tasks WHERE id=$1", id)
	if err != nil {
		fmt.Println("Error of getting tasks", err)
	}
	var count int
	rows.Next()
	err = rows.Scan(&count)
	if err != nil {
		fmt.Println("Error of scan in GetNumOfTasksDB", err)
	}
	return count
}
func GetTasksFromDB(db *sql.DB, id int) *sql.Rows {
	rows, err := db.Query("SELECT task FROM agents_tasks WHERE id=$1", id)
	if err != nil {
		fmt.Println("Error of getting tasks", err)
	}
	_, err = db.Query("DELETE FROM agents_tasks WHERE id=$1", id)
	if err != nil {
		fmt.Println("Error of deleting tasks", err)
	}
	return rows
}
