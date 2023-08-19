package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Setup() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Failed to load .env")
	}

	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")
	if mongoConnectionString == "" {
		log.Fatalln("Database connection string not found.")
	}

	TESTDB := os.Getenv("DATABASE_TEST")
	EVENTSDB := os.Getenv("DATABASE_EVENTS")
	AUTHDB := os.Getenv("DATABASE_AUTH")
	PREFDB := os.Getenv("DATABASE_AUTH")

	if TESTDB == "" || EVENTSDB == "" || AUTHDB == "" {
		log.Fatalf("Critical DB not found. Found [%v %v %v]:\n", TESTDB , EVENTSDB, AUTHDB)
	}

	if PREFDB == "" {
		log.Println("User Preferences DB not found.")
	}

	sanityTestLogs := SanityTest(mongoConnectionString, TESTDB)
	log.Println(sanityTestLogs)
}
