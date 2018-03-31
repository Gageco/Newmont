package main

import (
  "time"
)

type Data struct {
  Time        time.Time     `json:"time"`
  Unit        int           `json:"unit"`
  Data        int           `json:"data"`
  Id          int           `json:"id"`
}

type Datas []Data
