package handler

import (
	"SystemUtilManager/database_integration"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Agent struct {
	Id       int    `json:"id"`
	Hostname string `json:"hostname"`
}

type Data struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func GetResOfScan(w http.ResponseWriter, req *http.Request) {
	AgentID := req.Header.Get("AgentID")
	_, err := os.Stat(AgentID)
	if os.IsNotExist(err) {
		err := os.Mkdir(AgentID, 0777)
		if err != nil {
			fmt.Println("Error of creating dir", err)
		}
	}

	err = req.ParseMultipartForm(20)
	if err != nil {
		fmt.Println("Error of ParseMultipartForm:", err)
		http.Error(w, "Error of ParseMultipartForm", http.StatusBadRequest)
		return
	}

	file, _, err := req.FormFile("data")
	if err != nil {
		http.Error(w, "Error of FormFile", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	outFile, err := os.Create(AgentID + "/res_file")
	if err != nil {
		http.Error(w, "Error of CreateFile", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, "Error of CopyFile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File uploaded")
}

func Synchronization(w http.ResponseWriter, req *http.Request) {
	data, _ := io.ReadAll(req.Body)
	agent := Agent{}
	err := json.Unmarshal(data, &agent)
	if err != nil {
		return
	}
	db := database_integration.ConnectToDB()
	if database_integration.CheckAgents(agent.Id, db) {
		database_integration.UpdateTimeAgent(agent.Id, db)
		w.WriteHeader(http.StatusOK)
		return
	} else {
		http.Error(w, "Bad value of agent", http.StatusBadRequest)
		return
	}
}
