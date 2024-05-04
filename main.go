package main

import (
 	"fmt"
 	"log"
 	"main.go/logicFibonachi"
 	"net/http"
 	"os"
 	"strconv"
)

func render(w http.ResponseWriter, r *http.Request, filePath string) {
 	file, err := os.ReadFile(filePath)
 	if err != nil {
  		fmt.Println("Что-то не так:(")
  		return
 	}
 	fmt.Fprintf(w, "%s", file)
}

func serveHandler(w http.ResponseWriter, r *http.Request) {
 	if r.Method == "GET" {
 		render(w, r, "./html/inputFibonachi.html")
 	} else if r.Method == "POST" {
  		formHandler(w, r)
 	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
 	fs := logic.FibonacciService{}
 	render(w, r, "./html/resultFibonachi.html")
 	number, err := strconv.Atoi(r.FormValue("numberValue"))

 	if err != nil || number < 0 {
  		fmt.Fprintf(w, "<h2>Введённое число некорректно!!!</h2>")
 		} else {

  	if fs.IsFibonacci(number) {
   		prev, next := fs.GetAdjacentFibonacci(number)
   		fmt.Fprintf(w, "<div>Предыдущее число: %d</div>", prev)
   		fmt.Fprintf(w, "<div>Следующее число: %d</div>", next)
  		} else {
   		closest := fs.GetNearestFibonacci(number)
   		fmt.Fprintf(w, "<div>В ряду Фибоначи такого числа нету, но ближайшее к нему число: %d</div>", closest)
  		}
 	}

 	fmt.Fprintf(w, "<a href=\"/\">Вернуться</a>")
}

func main() {
 	http.HandleFunc("/", serveHandler)

 	log.Fatal(http.ListenAndServe(":8080", nil))
}