package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arr_user []Users
	var response Response

	db, e := MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Success")

	rows, err := db.Query("Select id,first_name,last_name from person")

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.FirstName, &users.LastName); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_user = append(arr_user, users)
		}

	}

	response.Status = 1
	response.Message = "Success ok"
	response.Data = arr_user
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", returnAllUsers).Methods("GET")
	http.Handle("/", router)
	fmt.Println("starting web server at http://localhost:8000/")
	http.ListenAndServe(":8000", router)
}
