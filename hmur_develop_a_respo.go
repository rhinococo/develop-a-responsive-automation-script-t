package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// AutomationScript represents a single automation script
type AutomationScript struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// AutomationScripts is a collection of AutomationScript
type AutomationScripts []AutomationScript

var automationScripts = AutomationScripts{}

func init() {
	automationScripts = AutomationScripts{
		{ID: "1", Name: "Script 1", Status: "Running", StartTime: time.Now().Format(time.RFC3339), EndTime: ""},
		{ID: "2", Name: "Script 2", Status: "Failed", StartTime: time.Now().Add(-10 * time.Minute).Format(time.RFC3339), EndTime: time.Now().Format(time.RFC3339)},
		{ID: "3", Name: "Script 3", Status: "Pending", StartTime: "", EndTime: ""},
	}
}

func getScripts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(automationScripts)
}

func getScript(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range automationScripts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&AutomationScript{})
}

func createScript(w http.ResponseWriter, r *http.Request) {
	var newScript AutomationScript
	_ = json.NewDecoder(r.Body).Decode(&newScript)
	automationScripts = append(automationScripts, newScript)
	json.NewEncoder(w).Encode(newScript)
}

func updateScript(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range automationScripts {
		if item.ID == params["id"] {
			_ = json.NewDecoder(r.Body).Decode(&item)
			automationScripts[index] = item
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&AutomationScript{})
}

func deleteScript(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range automationScripts {
		if item.ID == params["id"] {
			automationScripts = append(automationScripts[:index], automationScripts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(&AutomationScripts{})
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/scripts", getScripts).Methods("GET")
	router.HandleFunc("/scripts/{id}", getScript).Methods("GET")
	router.HandleFunc("/scripts", createScript).Methods("POST")
	router.HandleFunc("/scripts/{id}", updateScript).Methods("PUT")
	router.HandleFunc("/scripts/{id}", deleteScript).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}