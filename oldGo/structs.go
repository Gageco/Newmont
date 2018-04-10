package main

import (
  "time"
)

type Data struct {
  Time               time.Time     `json:"time"`
  Humidity           int           `json:"unit"`
  Temperature        int           `json:"data"`
}

type Datas []Data
