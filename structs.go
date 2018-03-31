package main

type Data struct {
  Time        string        `json:"time"`
  Id          int           `json:"id"`
  Data        string        `json:"data"`
}

type Datas []Data
