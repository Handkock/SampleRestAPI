package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	ApiVersion = "/api/"
	EntityName = "counter"
	ServerPort = "8000"
)

type App struct {
	Environment *Environment
	DB          *sql.DB
}

/*
	handles the request callbacks
*/
type CounterHandler struct {
	Counter *Counter
	DB      *sql.DB
}

func (cH *CounterHandler) increment(w http.ResponseWriter, r *http.Request) {
	cH.Counter.Increment(*cH.DB)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode("Counter succefully incremented")
}

func (cH *CounterHandler) decrement(w http.ResponseWriter, r *http.Request) {
	if !cH.Counter.Decrement(*cH.DB) {
		http.Error(w, "Value of the counter cannot be lower than 0", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode("Counter succefully decremented")
}

func (cH *CounterHandler) getCounter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, strconv.Itoa(cH.Counter.GetCounter(*cH.DB)))
}

func (app *App) RunServer() {
	// inititalize the counter
	currentTime := time.Now()
	counter := &Counter{1, 0, currentTime.Format("2006-01-02"), currentTime.Format("2006-01-02")}

	// create in in the db
	counter.Create(*app.DB)

	// register the handler
	handler := CounterHandler{counter, app.DB}

	r := mux.NewRouter()
	api := r.PathPrefix(ApiVersion + EntityName).Subrouter()

	// bind the routes to corresponding actions
	api.HandleFunc("/increment", handler.increment).Methods(http.MethodPut)
	api.HandleFunc("/decrement", handler.decrement).Methods(http.MethodPut)
	api.HandleFunc("", handler.getCounter).Methods(http.MethodGet)

	fmt.Printf("Server starting at port %v\n", ServerPort)
	log.Fatal(http.ListenAndServe(":"+ServerPort, r))
}
