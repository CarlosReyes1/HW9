package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func upload(res http.ResponseWriter, req *http.Request) {

	tpl, err := template.ParseFiles("./temp.html")
	if err != nil {
		log.Fatal("error:", err)
	}

	tpl.Execute(res, nil)

	if req.Method == "Post" {
		_, src, err := req.FormFile("Name")
		if err != nil {
			fmt.Println(err)
		}
		dest, err := src.Open()
		if err != nil {
			fmt.Println(err)
		}

		io.Copy(res, dest)
	}

}

func main() {
	http.HandleFunc("/", upload)
	http.ListenAndServe(":8080", nil)
}
