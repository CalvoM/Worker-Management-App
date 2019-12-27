package main
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)
//db credentials
const(
	host = <HOST>
	port = 5432
	user = <DB user>
	password = <DB password>
	dbname = <DB name>
)
var workers []*Worker
func main(){
	defer keepCalm()
	//Have connection to DB first.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)
	db,err := sql.Open("postgres",psqlInfo)
	if err !=nil{
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil{
		panic(err)
	}
	fmt.Println("Successfully DB Connected")
	//then make server
	r := mux.NewRouter()
	r.HandleFunc("/registerworker/",func(w http.ResponseWriter, rw *http.Request){
		RegisterWorkerHandler(w,rw,db)
	})
	r.HandleFunc("/workers/",func(w http.ResponseWriter, rw *http.Request){
		GetWorkersHandler(w,rw,db)
	})
	log.Println("Server started")
	http.ListenAndServe(":8082",r)
}

/*
@brief Handle all panics called in the the program
*/
func keepCalm(){
	if r:=recover(); r!=nil{
		log.Println("Recovering :",r)
	}
}

/*
@brief Do insertion to gotest table
*/
func insertToDb(worker Worker,db *sql.DB){
	_,err := db.Exec(`insert into gotest (name,age,salary,email) values($1,$2,$3,$4)`,worker.Name,worker.Age,worker.Salary,worker.Email)
	if err != nil{
		panic(err)
	}
	log.Println("Insertion to DB successful!")
}

/*
@brief receive the worker credentials and add to DB
*/
func RegisterWorkerHandler(_ http.ResponseWriter, rw *http.Request,db *sql.DB){
	var worker Worker
	err :=json.NewDecoder(rw.Body).Decode(&worker)
	if err !=nil{
		panic(err)
	}
	insertToDb(worker,db)
}

/*
@brief get all the workers in the DB
*/
func GetWorkersHandler(w http.ResponseWriter, rw *http.Request,db *sql.DB){
	rows,err :=db.Query(`select * from gotest`)
	if err !=nil{
		panic(err)
	}
	defer rows.Close()
	// Loop through the rows returned
	for rows.Next(){
		var(
			id int
			name string
			age int8
			salary float64
			email string
		)
		err := rows.Scan(&id,&name,&age,&salary,&email)
		if err !=nil {
			panic(err)
		}
		worker := Worker{name,age,salary,email}
		workers =append(workers,&worker)
	}
	err = rows.Err()
	if err !=nil{
		panic(err)
	}
	resp := Response{workers}
	log.Println(resp.Data)
	json.NewEncoder(w).Encode(resp)
	//Empty the workers slice or it accrue for each call to server, thus replication of workers
	workers = workers[:0]
}