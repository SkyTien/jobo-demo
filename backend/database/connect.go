package database

import (
	config "Goal-Back-End/conf"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var SqlDB *sql.DB

func Init() {
	DBconfig := config.GetConfig().Database
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DBconfig.Host,
		DBconfig.Port,
		DBconfig.User,
		DBconfig.Password,
		DBconfig.DBname)
	var err error
	SqlDB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer SqlDB.Close()

	err = SqlDB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
