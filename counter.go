package main

import (
	"database/sql"
	"fmt"
	"log"
)

const TableName = "counter"

/*
	counter model
*/
type Counter struct {
	ID      int
	value   int
	created string
	updated string
}

/*
	writes the counter information into db
*/
func (p *Counter) Create(db sql.DB) {
	db.Query(fmt.Sprintf("INSERT INTO %v () VALUES (?,?,?,?)", TableName), p.ID, p.value, p.created, p.updated)

}

/*
	increments und updates the db record
*/
func (p *Counter) Increment(db sql.DB) {
	newValue := p.GetCounter(db) + 1
	_, err := db.Query(fmt.Sprintf("UPDATE %v SET value = ? WHERE id = ?", TableName), newValue, p.ID)
	if err != nil {
		log.Fatal(err)
	}
}

/*
	decrements und updates the db record
*/
func (p *Counter) Decrement(db sql.DB) bool {
	newValue := p.GetCounter(db) - 1
	if newValue < 0 {
		return false
	}
	_, err := db.Query(fmt.Sprintf("UPDATE %v SET value = ? WHERE id = ?", TableName), newValue, p.ID)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

/*
	gets the current value from the db
*/
func (p *Counter) GetCounter(db sql.DB) int {
	var value int
	err := db.QueryRow(fmt.Sprintf("SELECT value FROM %v WHERE id = ?", TableName), p.ID).Scan(&value)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
