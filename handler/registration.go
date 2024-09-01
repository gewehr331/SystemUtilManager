package handler

import (
	"SystemUtilManager/database_integration"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type NotRegisteredAgent struct {
	Hostname string `json:"hostname"`
}

func registration(w http.ResponseWriter, req *http.Request) {
	data, _ := io.ReadAll(req.Body)
	agent := NotRegisteredAgent{}
	err := json.Unmarshal(data, &agent)
	var res int
	if err != nil {
		fmt.Println("Error of unmarshalling unregistered agent", err)
	}
	db := database_integration.ConnectToDB()
	if database_integration.CheckHostnameAgents(agent.Hostname, db) {
		database_integration.AddAgentToDB(agent.Hostname, db)
		res = database_integration.GetIDByHostname(db, agent.Hostname)

	} else {
		res = 0
	}
	fmt.Fprintf(w, strconv.Itoa(res))
}

func main() {

}
