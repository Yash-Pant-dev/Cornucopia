package main

// API Server for GO.
import (
// "encoding/json"

// "net/http"

// "github.com/gin-gonic/gin"

// "context"
// "log"
// // "os"
// "strings"
// // "github.com/joho/godotenv"
// // "go.mongodb.org/mongo-driver/bson"
// "go.mongodb.org/mongo-driver/mongo"
// "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	Setup()

	// router := gin.Default()

	// router.Run("127.0.0.1:8080")

}

// func setup() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatalln("Failed to load .env")
// 	}

// 	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")
// 	if mongoConnectionString == "" {
// 		log.Fatalln("Database connection string not found.")
// 	}

// 	TESTDB := os.Getenv("DATABASE_TEST")
// 	EVENTSDB := os.Getenv("DATABASE_EVENTS")
// 	AUTHDB := os.Getenv("DATABASE_AUTH")
// 	PREFDB := os.Getenv("DATABASE_AUTH")

// 	if TESTDB == "" || EVENTSDB == "" || AUTHDB == "" {
// 		log.Fatalln("Critical DB not found. [Test %v / Events %v/ Auth %v]", TESTDB, EVENTSDB, AUTHDB)
// 	}

// 	if PREFDB == "" {
// 		log.Println("User Preferences DB not found.")
// 	}

// 	sanityTestLogs := sanityTest(mongoConnectionString, TESTDB)
// 	log.Println(sanityTestLogs)
// }

// func sanityTest(dbConnectionString, TESTDB string) string {
// 	var sanityTestLogs strings.Builder
// 	sanityTestLogs.WriteString("Sanity Check started.")

// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbConnectionString))
// 	if err != nil {
// 		log.Fatalln("Failed to connect to Mongo Server.")
// 	}
// 	sanityTestLogs.WriteString("Client successfully connected.\n")

// 	defer func() {
// 		if client.Disconnect(context.TODO()); err != nil {
// 			sanityTestLogs.WriteString("Disconnect failed.")
// 			log.Fatalln("Failed to disconnect.")
// 		}
// 		sanityTestLogs.WriteString("Disconnected successfully.\n")
// 		sanityTestLogs.WriteString("Sanity Check ended.\n")
// 	}()

// 	// Testing in the Sanity Check Database.
// 	testDB_Events := client.Database(TESTDB).Collection("Events")
// 	testDB_Auth := client.Database(TESTDB).Collection("Auth")

// 	newEvent := Event{ID: "1", Name: "FirstEvent", Desc: "SanityEvent", RDate: "17/11/2001/17:30"}
// 	newUser := Auth{Username: "Yash", Pass: "StrongPass", Salt: "SomeRandomString"}
// 	_, errEvent := testDB_Events.InsertOne(context.TODO(), newEvent)
// 	_, errAuth := testDB_Auth.InsertOne(context.TODO(), newUser)

// 	if errEvent != nil && errAuth != nil {
// 		sanityTestLogs.WriteString("Error in inserting.\n")
// 		// sanityTestLogs.WriteString("EventCollection: %v \n",errEvent.Error())
// 		// sanityTestLogs.WriteString("AuthCollection : %v \n",errAuth.Error())
// 		log.Fatalln("Error in sanity insertion. %v %v", errEvent, errAuth)
// 	}

// 	return sanityTestLogs.String()
// }

// func getHomepage(c *gin.Context) {
// 	print("Got Request")
// 	print(json.Marshal(firstPage))
// 	// jdata, _ := json.Marshal(firstPage)
// 	c.JSON(http.StatusOK, firstPage)
// }
