package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"nativepub.net/idrequest/cron"
	"nativepub.net/idrequest/csv"
)

type Environment struct {
	synchronizedMap *csv.SynchronizedMap
}

func main() {
	synchronizedMap := csv.NewSynchronizedMap()
	env := &Environment{synchronizedMap: synchronizedMap}
	cron.ReloadMapEvery30Minutes(synchronizedMap)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/promotions/{key}", env.Promotions)
	log.Fatal(http.ListenAndServe(":1321", router))
}

func (env *Environment) Promotions(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	routeVariables := mux.Vars(request)

	record, ok := env.synchronizedMap.Load(routeVariables["key"])
	if ok != true {
		http.Error(writer, "Key not found", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": routeVariables["key"], "price": record[0], "expiration_date": record[1]}
	responseJson, err := json.Marshal(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseJson)
}
