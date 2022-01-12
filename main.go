package main

import (
	AsciiArt "AsciiArt/asciiart"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

var mainTemplate, errorTemplate, styleTemplate1 *template.Template

func main() {
	var err error
	mainTemplate, err = template.ParseFiles("PageTemplate/index.html")
	CheckError(err)
	errorTemplate, err = template.ParseFiles("PageTemplate/error.html")
	CheckError(err)
	styleTemplate1, err = template.ParseFiles("PageTemplate/style.css")
	CheckError(err)
	//We register our handlers on server routes using the http.HandleFunc convenience function.
	//It sets up the default router in the net/http package and takes a function as an argument.
	fs := http.FileServer(http.Dir("PageTemplate"))
	http.Handle("/PageTemplate/", http.StripPrefix("/PageTemplate/", fs))
	http.HandleFunc("/", PathPage)
	http.HandleFunc("/ascii-art", StatusPage)
	fmt.Println("Server listening on port 8000...")
	//Finally, we call the ListenAndServe with the port and a handler.
	//nil tells it to use the default router we’ve just set up.
	log.Fatal(http.ListenAndServe(":8000", nil))
}

//CheckError is looking over errors
func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

//StatusPage is executing request(respond to request)
func StatusPage(w http.ResponseWriter, r *http.Request) {
	//checking method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		errorTemplate.Execute(w, "405: Status method not allowed")
		return
	}

	//Getting FontStyle
	font := r.FormValue("banner")
	if font != "standard" && font != "shadow" && font != "thinkertoy" {
		w.WriteHeader(http.StatusBadRequest)
		errorTemplate.Execute(w, "400: Bad Request")
		return
	}

	//Getting Text
	text := r.FormValue("data")
	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
		errorTemplate.Execute(w, "400: Bad Request")
		return
	}

	//removing carriage ret and checking symbols
	var textUpdated string
	for _, symbol := range text {
		if symbol == 13 {
			continue
		}
		if !(symbol >= 32 && symbol <= 126 || symbol == 10) {
			w.WriteHeader(http.StatusBadRequest)
			errorTemplate.Execute(w, "400: Bad Request")
			return
		}
		textUpdated += string(symbol)
	}

	//getting ascii art
	Art, err := AsciiArt.GetASCII(textUpdated, font)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorTemplate.Execute(w, "500: Internal Server Error")
		return
	}
	ButtonClicked := r.FormValue("buttons")
	if ButtonClicked == "Download" {
		lenS := strconv.Itoa(len(Art))
		w.Header().Set("Content-Disposition", "attachment; filename=asciiart.txt")
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", lenS)
		w.Write([]byte(Art))
	}else if ButtonClicked=="Submit"{
		mainTemplate.Execute(w, Art)
	}else {
		w.WriteHeader(http.StatusBadRequest)
		errorTemplate.Execute(w, "400: Bad Request")
	}
}

//PathPage is response to any request excluding ascii art
func PathPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		errorTemplate.Execute(w, "404: Not found")
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		errorTemplate.Execute(w, "405: Status method not allowed")
		return
	}
	mainTemplate.Execute(w, nil)
}
