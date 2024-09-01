package main

import (
	"SystemUtilManager/router"
	"fmt"
	"net/http"
)

func main() {

	err := http.ListenAndServe("localhost:8090", router.Router)
	if err != nil {
		fmt.Println("Error of listen and serve")
	}
}
