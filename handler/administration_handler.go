package handler

import (
	"SystemUtilManager/database_integration"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
)

type TableRows struct {
	Rows template.HTML
}

type AgentRows struct {
	AgentApps  template.HTML
	AgentPorts template.HTML
}

func AdminPanel(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		tmplt, err := template.ParseFiles("ui/html/agents.html")
		if err != nil {
			fmt.Println("Error of parsing template", err)
		}
		db := database_integration.ConnectToDB()
		defer db.Close()
		agents := database_integration.GetAllAgents(db)
		var tableRows string
		for agents.Next() {
			var AgentId string
			var Hostname string
			var syncDate string

			err := agents.Scan(&AgentId, &Hostname, &syncDate)
			if err != nil {
				fmt.Println("error of scan", err)
				return
			}
			tableRows += "<tr><td><a href=\"/agent/" + AgentId + "\">" + AgentId + "</a></td><td><a href=\"/agent/" + AgentId + "\">" + Hostname + "</a></td><td><a href=\"/agent/" + AgentId + "\">" + syncDate + "</a></td></tr>"

		}

		data := TableRows{Rows: template.HTML(tableRows)}

		err = tmplt.Execute(w, data)
		if err != nil {
			fmt.Println("Error of executing template", err)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func Agents(w http.ResponseWriter, req *http.Request) {
	agentInfo := mux.Vars(req)

	file, err := os.Open(string(agentInfo["id"]) + "/res_file")
	if err != nil {
		fmt.Println("Error of reading agent file", err)
	}

	decoder := json.NewDecoder(file)
	var dataObjects []Data
	for {
		var object Data
		err := decoder.Decode(&object)
		if err != nil {
			break
		} else {
			dataObjects = append(dataObjects, object)
		}
	}
	var Apps string
	var Ports string

	for _, dataObject := range dataObjects {
		if dataObject.Type == "App" {
			Apps += "<tr><td>" + dataObject.Value + "</td></tr>"
		} else if dataObject.Type == "Port" {
			Ports += "<tr><td>" + dataObject.Value + "</td></tr>"
		}

	}
	tmplt, err := template.ParseFiles("ui/html/agent_information.html")
	if err != nil {
		fmt.Println("Error of parsing template agent_information.html", err)
	}

	data := AgentRows{AgentApps: template.HTML(Apps), AgentPorts: template.HTML(Ports)}

	err = tmplt.Execute(w, data)
	if err != nil {
		fmt.Println("Error of executing template agent_information", err)
	}
}
