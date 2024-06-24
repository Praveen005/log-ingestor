package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log-ingestor/models"
	"net/http" 

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
)

//we had to define this struct because we want to create GetLogData() CreateLogData()  DeleteLogData() functions
// GetLogData() CreateLogData()  DeleteLogData() are not simple functions, but struct methods
type LogDataController struct{
	session *mongo.Client
}

//remeber session 's' is taken as parameter here fron getSession()
//it returns a pointer to LogDataController
func NewLogDataController(s *mongo.Client) *LogDataController{
	return &LogDataController{s}
}



//Again a struct method
//Response is what we will be returning from this function
//ResponseWriter helps in sending the response to the LogData about the staus of request
func (uc LogDataController) CreateLogData(w http.ResponseWriter, r *http.Request, _ httprouter.Params){ // you don't need params here, since it's a POST request

	u := models.LogData{}

	//decoding the json value that you'd be getting
	json.NewDecoder(r.Body).Decode(&u)

	

	//inserted
	uc.session.Database("logIngestorDB").Collection("logs").InsertOne(context.TODO(), u)

	//now that data has been inserted, you want to tell the client that, hey, it has been inserted
	uj, err := json.Marshal(u)

	//if you get error
	if err != nil{
		fmt.Println(err)
	}
	//if all went well, you create header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)

}