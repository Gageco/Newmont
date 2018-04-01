var request = require("request")
var fs = require('fs')

//POST Requesst
var jsonPost = {"unit": 11, "data": 20}


request({
  url: "http://localhost:8080/data",
  json: true,
  method: "POST",
  body: jsonPost,
})

for (i=0; i<16; i++) {
  jsonPost["data"] = i*i
  request({
    url: "http://localhost:8080/data",
    json: true,
    method: "POST",
    body: jsonPost,
  })
}
