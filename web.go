package main

import (
    "context"
    "fmt"
    "net/http"
    "gopkg.in/mgo.v2"
    "os"
    "strings"
)
var ctx = context.Background()
func getallenv(w http.ResponseWriter, r *http.Request){
	for _, element:= range os.Environ(){
		variabel := strings.Split(element,"=")
		fmt.Fprintln(w,variabel[0],"=>",variabel[1])
	}
}

func test(w http.ResponseWriter, r *http.Request){
    sess, err := mgo.Dial(os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT"))
    if err != nil {
        fmt.Fprintln(w,err)
        fmt.Println(err)
    }
    defer sess.Close()
    err = sess.Ping()
    if err != nil {
        fmt.Fprintln(w,err)
        fmt.Println(err)
    }
        fmt.Fprintln(w,"MongoDB server is healthy.")
}
func main() {
  http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w,"Belajar Integrasi Network Docker")
	})

	http.HandleFunc("/test",test)
    http.HandleFunc("/env",getallenv)
	http.ListenAndServe(":8080",nil)
}
