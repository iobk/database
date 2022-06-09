package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn, err := sql.Open("mysql", "guest:guest@tcp(ketisedu4.synology.me:3306)/testdb")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	for i := 0; i < 10; i++ {
		_, err := conn.Exec("INSERT INTO user1(userName,phoneNumber,workerName,dateTime) VALUE ('js kim',?,'david',now())", 100+i)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("/ ", i)
	}

}
