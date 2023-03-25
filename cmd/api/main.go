package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	YearsOld int `json:"yearsOld"`
}

type Response struct {
    StatusCode int `json:"statusCode"`
    Message string `json:"message"`
}

var newUsers []User

func main() {	
	muxtRouter := mux.NewRouter()

    muxtRouter.HandleFunc("/users", users).Methods("GET")
	muxtRouter.HandleFunc("/getNewUsers", getNewUsers).Methods("GET")
    muxtRouter.HandleFunc("/addUser", addUser).Methods("POST")
	muxtRouter.HandleFunc("/searchUserById", searchUserById).Methods("POST")

	fmt.Println("LIVE API LOCALHOST Server PORT 8080")

	log.Fatal(http.ListenAndServe(":8080", muxtRouter))
}

func getNewUsers (w http.ResponseWriter, r *http.Request) { 
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	if newUsers == nil {
		responseError := Response{StatusCode: 200, Message: "No new users"}
		json.NewEncoder(w).Encode(&responseError)
        return
    }

	json.NewEncoder(w).Encode(newUsers)
}

func getUsers () []User { 

	return []User {
		{
			Id: 1,
			Name:"John",
			YearsOld: 18,
		},
		{
			Id: 2,
			Name:"Thainan",
			YearsOld: 5,
		},
		{
			Id: 3,
			Name:"Luis",
			YearsOld: 11,
		},
		{
			Id: 4,
			Name:"Carlos",
			YearsOld: 13,
		},
		{
			Id: 5,
			Name:"Diego",
			YearsOld: 22,
		},
		{
			Id: 6,
			Name:"Manuel",
			YearsOld: 34,
		},
		{
			Id: 7,
			Name:"Carla",
			YearsOld: 31,
		},
	}
}

func users (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	json.NewEncoder(w).Encode(getUsers())
}

func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	defer r.Body.Close()

	var payloadUser User

	err := json.NewDecoder(r.Body).Decode(&payloadUser)
	if err != nil {
		responseError := Response{StatusCode: 406, Message: err.Error()}
		json.NewEncoder(w).Encode(&responseError)
        return
    }

   // decode `r.Body` `r *http.Request` to get data from request
   // decode the result to todo
   json.NewDecoder(r.Body).Decode(&payloadUser)

   // append new todo in todos
   newUsers = append(newUsers, payloadUser)

   // finally, encode again create the resutl with json format
   json.NewEncoder(w).Encode(newUsers)
}

func searchUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	defer r.Body.Close()

	var payloadUser User

	err := json.NewDecoder(r.Body).Decode(&payloadUser)
	if err != nil {
		responseError := Response{StatusCode: 406, Message: err.Error()}
		json.NewEncoder(w).Encode(&responseError)
        return
    }

   // decode `r.Body` `r *http.Request` to get data from request
   // decode the result to todo
   json.NewDecoder(r.Body).Decode(&payloadUser)
   // append new todo in todos
   //newUsers = append(newUsers, payloadUser)
   // finally, encode again create the resutl with json format
   //json.NewEncoder(w).Encode(newUsers)

    for _, user := range newUsers {
		if (user.Id == payloadUser.Id) {
			json.NewEncoder(w).Encode(payloadUser)
			return
		}
	}

	responseError := Response{StatusCode: 404, Message: "User not found"}

	json.NewEncoder(w).Encode(&responseError)
	return
}