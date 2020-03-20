package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("E:/Media"))))
	http.Handle("/media2/", http.StripPrefix("/media2/", http.FileServer(http.Dir("D:/Media"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/app", app)
	http.HandleFunc("/prime", prime)
	err:= http.ListenAndServe(":8000", nil)
	if err != nil {
		panic("Couldn't start server on port: 8000")
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	Respond(w, 200, "Please download the app to continue!  https://drive.google.com/open?id=1ohr1ce7CnQa_TmDCDqRkezSeKdfK40on", nil)
}

func app(w http.ResponseWriter, r *http.Request) {
	file, err:= os.Open("database.json")
	if err != nil {
		Respond(w, 104, "Unable to read database", nil)
		return
	}
	defer file.Close()
	data, err:= ioutil.ReadAll(file)
	if err != nil {
		Respond(w, 104, "Unable to read database", nil)
		return
	}

	type Database struct {
		Filename string
		Data_created string
		Date_modified string
		Version string
		Files []interface{}
	}

	db:= Database{}
	_= json.Unmarshal(data, &db)
	Respond(w, 200, "Fetch Successful", db)
}

func prime(w http.ResponseWriter, r *http.Request) {
	file, err:= os.Open("database-prime.json")
	if err != nil {
		Respond(w, 104, "Unable to read database", nil)
		return
	}
	defer file.Close()
	data, err:= ioutil.ReadAll(file)
	if err != nil {
		Respond(w, 104, "Unable to read database", nil)
		return
	}

	type Database struct {
		Filename string
		Data_created string
		Date_modified string
		Version string
		Files []interface{}
	}

	db:= Database{}
	_= json.Unmarshal(data, &db)
	Respond(w, 200, "Fetch Successful", db)
}

func Respond(w http.ResponseWriter, code int64, info string, data interface{}) {
	type Response struct {
		Code int64
		Info string
		Data interface{}
	}
	response:=&Response{Code:code, Info:info, Data:data}
	resp, err:= json.Marshal(response)
	if err != nil {
		log.Fatal("Error 102: Failed to generate response for target request")
	}
	_,_= w.Write(resp)
}