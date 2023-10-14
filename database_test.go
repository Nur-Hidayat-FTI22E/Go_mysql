package gomysql

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test(t *testing.T) {

}

func TestOpenconnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:nra1223004@tcp(localhost:3306)/belajar_golang_database")

	if err != nil {
		panic(err)
	}
	defer db.Close()
}
