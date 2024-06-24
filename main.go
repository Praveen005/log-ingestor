//In main.go we will basically have routes
//the functions that will help us to work with these routes will be in controllers
//our models will be very simple, it will just have struct of how the user model will look like
//main.go is the entry point to our project
package main

//there should be space after import
import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"

	"log-ingestor/controllers"
)

func main(){
	r := httprouter.New()
	//we need a new session
	uc := controllers.NewLogDataController(getClient())
	// r.GET("/LogData/:id", uc.GetLogData)
	r.POST("/LogData", uc.CreateLogData )
	// r.DELETE("/LogData/:id", uc.DeleteLogData)
	http.ListenAndServe("localhost:9000", r)
}

//it will return a mongoDb session
func getClient() *mongo.Client {
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// clientOptions := options.Client().ApplyURI("mongodb://writeUser:secret@localhost:27017")
	clientOptions := options.Client().ApplyURI("mongodb://writeUser:secret@mongo1:27017,mongo2:27017,mongo3:27017/?replicaSet=myReplicaSet")
	
	clientOptions.SetRetryWrites(true) // Enable retryable writes
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	// Ping the server to ensure that the client can successfully connect
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}