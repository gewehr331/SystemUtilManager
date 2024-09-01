package database_integration

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func ConnectToDB() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=postgres host=localhost dbname=systemUtilDB sslmode=disable")
	if err != nil {
		fmt.Println("Error of opening database", err)
		return nil
	}
	return db
}

func CreateTable(db *sql.DB) {
	_, err := db.Query("CREATE TABLE agents (id integer, hostname varchar(128), sync_date varchar(4096)); ")
	if err != nil {
		fmt.Println("Error of creating table", err)
	}

}
func GetIDByHostname(db *sql.DB, hostname string) int {
	res, err := db.Query("SELECT id FROM agents WHERE hostname=$1;", hostname)
	if err != nil {
		fmt.Println("Error of creating query to get ID", err)
	}
	var id int
	res.Next()
	err = res.Scan(&id)
	if err != nil {
		fmt.Println("Error of scanning id results", err)
		return 0
	}
	return id
}
func GetHostnameByID(db *sql.DB, id int) string {
	res, err := db.Query("SELECT hostname FROM agents WHERE id=$1;", id)
	if err != nil {
		fmt.Println("Error of creating query to get hostname", err)
	}
	var hostname string
	res.Next()
	err = res.Scan(&hostname)
	if err != nil {
		fmt.Println("Error of scanning query to get id")
		return ""
	}
	return hostname

}
func GetAllAgents(db *sql.DB) *sql.Rows {
	res, err := db.Query("SELECT * FROM agents;")
	if err != nil {
		fmt.Println("Error of select all", err)
	}
	return res

}

func AddAgentToDB(hostname string, db *sql.DB) bool {
	rows, err := db.Query("SELECT COUNT(*);")
	if err != nil {
		fmt.Println("Error of select count all", err)
		return false
	}
	var numOfAgents int32
	RegTime := time.Now()
	rows.Next()
	err = rows.Scan(&numOfAgents)
	if err != nil {
		fmt.Println("Error of rows scan", err)
		return false
	}

	rows, err = db.Query("INSERT INTO agents VALUES ($1, $2, $3);", numOfAgents, hostname, RegTime)
	if err != nil {
		fmt.Println("Error of adding agent to table", err)
	}
	return true

}

func UpdateTimeAgent(id int, db *sql.DB) {
	NewTime := time.Now()
	_, err := db.Query("UPDATE agents SET sync_date = $1 WHERE id = $2", NewTime, id)
	if err != nil {
		fmt.Println("Error of updating time for agent", err)
	}

}

func CheckAgents(id int, db *sql.DB) bool { // Добавить обработку, если обнаружено несколько агентов в БД
	rows, err := db.Query("select count(*) from agents where id=$1", id)
	if err != nil {
		fmt.Println("Error of checking agent in db", err)
	}

	var numOfAgents int32
	rows.Next()
	err = rows.Scan(&numOfAgents)
	if err != nil {
		fmt.Println("Error of rows scan", err)
	}
	if numOfAgents == 1 {
		return true
	} else {
		return false
	}
}

func CheckHostnameAgents(hostname string, db *sql.DB) bool { // Добавить обработку, если обнаружено несколько агентов в БД
	rows, err := db.Query("select count(*) from agents where hostname=$1", hostname)
	if err != nil {
		fmt.Println("Error of checking agent in db", err)
	}

	var numOfAgents int32
	rows.Next()
	err = rows.Scan(&numOfAgents)
	if err != nil {
		fmt.Println("Error of rows scan", err)
	}
	if numOfAgents == 1 {
		return true
	} else {
		return false
	}
}
