package database

import(
	"database/sql"
    "fmt"
	"os"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
    DB *sql.DB
}

func NewDatabase() (Database, error) {

    USER := os.Getenv("DB_USERNAME")
    PASS := os.Getenv("DB_PASSWORD")
    HOST := os.Getenv("DB_HOSTNAME")
    DBNAME := os.Getenv("DB_NAME")

    if PASS == "" {
        DBPASS,_ := os.ReadFile("/run/secrets/DB_PASSWORD")
        PASS = strings.TrimSpace(string(DBPASS))
    }

    URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", USER, PASS, HOST, DBNAME)

        db, err := sql.Open("mysql", URL)

        if err != nil {
            return Database{}, err
        }

        fmt.Println("Database connection established.")

        //defer db.Close()

        return Database{
            DB: db,
        }, nil
}