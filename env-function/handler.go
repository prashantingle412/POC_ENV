package function

import (
	// function "POCjoin/join-function"
	"fmt"
	"log"
	"net/http"
	"os"
	// test2 "github.com/prashantingle412/test2/Model"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	// simple.Hello()
	// test2.User{}
	// function.User{}
	log.Println("successfully accessed")
	//accessed envirnment variables from saprate db files
	host := os.Getenv("host")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")
	log.Println("database credentials from external yml file ::", host, user, password, dbname)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("envirnment variables printed on  %s", "done")))
}
