package main

import (
  "time"
)

type Data struct {
  Time               time.Time        `json:"time"`
  Humidity           int        `json:"humidity"`
  Temperature        int        `json:"temperature"`
}

type Datas []Data
