package main

type siaFiles struct {
  Message    string       `json:"message"`
  Files      []siaFile    `json:"files"`
}

type siaFile struct {
  SiaPath    string    `json:"siapath"`
  Path       string    `json:"localpath"`
  Size       int       `json:"filesize"`
  Available  bool      `json:"available"`
  Renewing   bool      `json:"renewing"`
  Redundancy float64   `json:"redundancy"`
  Progress   float64   `json:"uploadprogress"`
}

type generalResp struct {
  Message    string    `json:"message"`
}

type fileData struct {
  Temperature    string    `json:"temperature"`
  Humidity       string    `json:"humidity"`
  Date           string    `json:"time"`
}

type Data struct {
  Time               string        `json:"time"`
  Humidity           int        `json:"humidity"`
  Temperature        int        `json:"temperature"`
}

type Datas []Data

type Message struct {
  Message    string    `json:"message"`
}
