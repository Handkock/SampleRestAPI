package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

type Environment struct {
	dbHost, dbPort, dbName, dbUserName, dbPassword string
}

func (env *Environment) loadEnvVars(test bool) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading environment variables")
	}

	// TODO handle through .env.test
	env.dbHost = os.Getenv("DB_HOST")
	env.dbPort = os.Getenv("DB_PORT")
	env.dbName = os.Getenv("DB_NAME")
	if test {
		env.dbName = os.Getenv("DB_NAME_TEST")
	}
	env.dbUserName = os.Getenv("DB_USERNAME")
	env.dbPassword = os.Getenv("DB_PASSWORD")
}

/*
	initialise the database connection from the env variables
*/
func InitDb(env Environment) *sql.DB {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
			env.dbUserName,
			env.dbPassword,
			env.dbHost,
			env.dbPort,
			env.dbName,
		),
	)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	return db
}

func cleanUp(db *sql.DB) {
	// close the db connection
	db.Close()
}

func main() {
	env := Environment{}
	env.loadEnvVars(false)
	database := InitDb(env)

	// clean termination
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Println("test", sig)
			cleanUp(database)
			os.Exit(1)
		}
	}()

	app := App{&env, database}
	app.RunServer()
}
