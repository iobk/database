package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn, err := sql.Open("mysql", "guest:guest@tcp(ketisedu4.synology.me:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM user1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var userName, phoneNumber, worker string
	var userID uint
	var dTime string
	for rows.Next() {
		err := rows.Scan(&userID, &userName, &phoneNumber, &worker, &dTime)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(userID, userName, phoneNumber, worker, dTime)
	}

}
