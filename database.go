package main

import (
	"database/sql"
	"fmt"
	"log"
)

type testDB struct {
	Id 		int
	Name 	sql.NullString
}


func Insertion(db *sql.DB, name string, major string, schoolId int) {
	insert, err := db.Query("INSERT INTO student VALUES (Null, ?, ?, ?)",name,major,schoolId)
	if err!=nil {log.Fatal("An error when inserting to mysql database: ",err)}
	defer insert.Close()
	fmt.Printf("Successfully inserted %s,to a MySQL Database!\n", name)
}

func Selection(db *sql.DB) {
	results, err := db.Query("SELECT student_id, student_name FROM student")
	if err!=nil {log.Fatal("An error when reading from mysql database: ",err)}
	for results.Next() {
		var user testDB

		// Scan have to pass a number of arguments = number of column return
		err := results.Scan(&user.Id, &user.Name)
		if err!=nil {log.Print("fatal: ",err)}
		fmt.Printf("ID: %d, name: %s\n",user.Id,user.Name.String)
	}
}
