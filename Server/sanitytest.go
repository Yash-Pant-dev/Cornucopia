package main

import (
	"log"
	"strings"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
)

func SanityTest(dbConnectionString, TESTDB string) string {
	var sanityTestLogs strings.Builder
	sanityTestLogs.WriteString("Sanity Check started.")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbConnectionString))
	if err != nil {
		log.Fatalln("Failed to connect to Mongo Server.")
	}
	sanityTestLogs.WriteString("Client successfully connected.\n")

	defer func() {
		if client.Disconnect(context.TODO()); err != nil {
			sanityTestLogs.WriteString("Disconnect failed.")
			log.Fatalln("Failed to disconnect.")
		}
		sanityTestLogs.WriteString("Disconnected successfully.\n")
		sanityTestLogs.WriteString("Sanity Check ended.\n")
	}()

	// Testing in the Sanity Check Database.
	testDB_Events := client.Database(TESTDB).Collection("Events")
	testDB_Auth := client.Database(TESTDB).Collection("Auth")
	
	newEvent := Event{ID: "1", Name: "FirstEvent", Desc: "SanityEvent", RDate: "17/11/2001/17:30"}
	newUser := Auth{Username: "Yash", Pass: "StrongPass", Salt: "SomeRandomString"}
	_, errEvent := testDB_Events.InsertOne(context.TODO(), newEvent)
	_, errAuth := testDB_Auth.InsertOne(context.TODO(), newUser)
	

	if errEvent != nil && errAuth != nil {
		sanityTestLogs.WriteString("Error in inserting.\n")
		// sanityTestLogs.WriteString("EventCollection: %v \n",errEvent.Error())
		// sanityTestLogs.WriteString("AuthCollection : %v \n",errAuth.Error())
		log.Fatalf("Error in sanity insertion. %v %v\n", errEvent, errAuth)
	}

	return sanityTestLogs.String()
}
