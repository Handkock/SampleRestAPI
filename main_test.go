package main

import (
	"database/sql"
	"log"
	"testing"
	"time"
)

func prepareDb() *sql.DB {
	env := Environment{}
	env.loadEnvVars(true) // should be probably a best way
	database := InitDb(env)
	clearTable(*database)
	return database
}

func prepareTestCounter(database *sql.DB) *Counter {
	currentTime := time.Now()
	testCounter := &Counter{1, 0, currentTime.Format("2006-01-02"), currentTime.Format("2006-01-02")}
	testCounter.Create(*database)
	return testCounter
}

/*
	test db to be initialised
*/
func TestInitDb(t *testing.T) {
	database := prepareDb()
	var count int
	err := database.QueryRow("select count(*) from information_schema.tables where table_schema = 'counterdb_test' and table_name = 'counter';").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count != 1 {
		t.Errorf("Expected %v database to contain the %v table", "counterdb_test", "counter")
	}
}

/*
	test the model actions
*/
func TestCounter_Create(t *testing.T) {
	//database := prepareDb()
	database, _ := sql.Open("mysql", "root:root@tcp(db:3306)/counterdb_test")
	var count int
	database.QueryRow("SELECT COUNT(*) FROM counter").Scan(&count)
	if count != 0 {
		t.Errorf("Expected counter table to have no rows, %v rows found", count)
	}
	prepareTestCounter(database)
	database.QueryRow("SELECT COUNT(*) FROM counter;").Scan(&count)
	if count != 1 {
		t.Errorf("Expected counter table to have a single row, %v rows found", count)
	}
	clearTable(*database)
}

func TestCounter_Increment(t *testing.T) {
	database := prepareDb()
	testCounter := prepareTestCounter(database)

	var value int
	err := database.QueryRow("SELECT value FROM counter where id = ?;", testCounter.ID).Scan(&value)
	if err != nil {
		log.Fatal(err)
	}
	if value != 0 {
		t.Errorf("Expected counter value to be 0, %v given %T", value, value)
	}

	testCounter.Increment(*database)
	err2 := database.QueryRow("SELECT value FROM counter where id = ?;", testCounter.ID).Scan(&value)
	if err2 != nil {
		log.Fatal(err2)
	}
	if value != 1 {
		t.Errorf("Expected counter value table to be 1, %v given", value)
	}
	clearTable(*database)
}

func TestCounter_Decrement(t *testing.T) {
	database := prepareDb()
	currentTime := time.Now()
	testCounter := &Counter{1, 1, currentTime.Format("2006-01-02"), currentTime.Format("2006-01-02")}
	testCounter.Create(*database)

	var value int
	database.QueryRow("SELECT value FROM counter where id = ?;", testCounter.ID).Scan(&value)
	if value != 1 {
		t.Errorf("Expected counter value table to be 0, %v given", value)
	}

	testCounter.Decrement(*database)
	database.QueryRow("SELECT value FROM counter where id = ?;", testCounter.ID).Scan(&value)
	if value != 0 {
		t.Errorf("Expected counter value table to be 0, %v given", value)
	}

	// should not be decremented, as is is would be lower than 0 otherwise, which is invalid
	if testCounter.Decrement(*database) {
		t.Errorf("Expected the counter not to be descremented anymore")
	}
}

func TestCounter_GetCounter(t *testing.T) {
	database := prepareDb()
	testCounter := prepareTestCounter(database)

	testCounter.Increment(*database)
	testCounter.Increment(*database)
	testCounter.Increment(*database)

	dbValue := testCounter.GetCounter(*database)
	if 3 != dbValue {
		t.Errorf("Expected counter value to be 3, %v given", dbValue)
	}
	clearTable(*database)
}

func clearTable(database sql.DB) {
	database.Query("TRUNCATE counter")
}
