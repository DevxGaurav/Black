package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("../../Media"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/app", app)
	err:= http.ListenAndServe(":8000", nil)
	if err != nil {
		panic("Couldn't start server on port: 8000")
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	Respond(w, 200, "Please download the app to continue!", "")
}

func app(w http.ResponseWriter, r *http.Request) {
	file, err:= os.Open("database.json")
	if err != nil {
		Respond(w, 104, "Unable to read database", "")
		return
	}
	defer file.Close()
	data, err:= ioutil.ReadAll(file)
	if err != nil {
		Respond(w, 104, "Unable to read database", "")
		return
	}
	resp:= string(data)
	resp=strings.ReplaceAll(resp, "\r\n", "")
	Respond(w, 200, "Fetch Successful", resp)
}

func Respond(w http.ResponseWriter, code int64, info string, data string) {
	type Response struct {
		Code int64
		Info string
		Data string
	}
	response:=&Response{Code:code, Info:info, Data:data}
	resp, err:= json.Marshal(response)
	if err != nil {
		log.Fatal("Error 102: Failed to generate response for target request")
	}
	_,_= w.Write(resp)
}