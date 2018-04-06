package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "io"
  "io/ioutil"
  "time"
  // "math/rand"
  //"github.com/gorilla/mux"
)

var dataID int
var datas Datas

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Welcome Newmont Challenge Contestants!")
  fmt.Fprintln(w, "Graph    -    http://api.io4.in:8080/graph")
  fmt.Fprintln(w, "Data     -    http://api.io4.in:8080/data")

}
func dataGET(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  // w.WriteHeader(422)
  err := json.NewEncoder(w).Encode(datas)
  checkErr(err)
}

func dataPOST(w http.ResponseWriter, r *http.Request) {
  var data Data

  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))                  //limit the requests that can come in
  checkErr(err)

  err = r.Body.Close()
  checkErr(err)

  if err := json.Unmarshal(body, &data); err != nil {                          // <- this is a pretty neat layout, its a line that dones something then checks for error, good to remember for future projects
    panic(err)
  }

  dataID +=1
  data.Id = dataID
  data.Time = time.Now()

  datas = append(datas, data)

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  err = json.NewEncoder(w).Encode(data)
  checkErr(err)
}
