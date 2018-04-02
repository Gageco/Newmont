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

for (i=0; i<21; i++) {
  jsonPost["data"] = i*i
  APIrequest()
  sleep(1000)


}

function APIrequest() {
  request({
    url: "http://localhost:8080/data",
    json: true,
    method: "POST",
    body: jsonPost,
  })
}
function sleep(milliseconds) {
  var start = new Date().getTime();
  for (var i = 0; i < 1e7; i++) {
    if ((new Date().getTime() - start) > milliseconds){
      break;
    }
  }
}
