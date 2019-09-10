package main
import (
        "fmt"
        "net/http"
        "strconv"
        "sync"
)

var counter int
var mutex =  &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "test")
}
func incrementCounter(w http.ResponseWriter, r *http.Request) {
        mutex.Lock()
        counter++
        fmt.Fprintf(w, strconv.Itoa(counter))
        mutex.Unlock()
}


func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
                fmt.Fprintf(w,"This is a website server by a Golang HTTP Server.")})

        http.HandleFunc("/increment", incrementCounter)

                fs := http.FileServer(http.Dir("static/"))
                http.Handle("/static/", http.StripPrefix("/static/", fs))
                http.ListenAndServe(":3001", nil)

}
