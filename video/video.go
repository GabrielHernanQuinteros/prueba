package video

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Hola() {
	fmt.Println("culo")
}

func ConectarDB(parConnectionString string) (*sql.DB, error) {
	return sql.Open("mysql", parConnectionString)
}
