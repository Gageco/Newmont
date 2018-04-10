package main

import(
  "net/http"
  "github.com/gorilla/mux"
)

type Route struct {                   //a struct for the Routes that we can use
  Name          string
  Method        string
  Pattern       string
  HandlerFunc   http.HandlerFunc
}

type Routes []Route                   //a struct that is a collection of Route

func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(false)
  for _, route := range routes {
    router.Methods(route.Method).
          Path(route.Pattern).
          Name(route.Name).
          Handler(route.HandlerFunc)
  }

  return router
}

var routes = Routes{
  Route{
    "INDEX",
    "GET",
    "/",
    Index,
  },
  Route{
    "dataPOST",
    "POST",
    "/data",
    dataPOST,
  },
  Route{
    "dataGET",
    "GET",
    "/data",
    dataGET,
  },
}
