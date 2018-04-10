package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
  "errors"
	"net/http"
  "path/filepath"
  "strings"
	"time"
	"strconv"
  // "bufio"
  "os"
)

func main() {
	// /var/www/htmls/data.csv
  GRAPH_FILE_LOCATION := "/Users/gagecoprivnicar/Documents/Github/Newmont/sia/site/graph/data.csv"

  startup(GRAPH_FILE_LOCATION)

  for i:=0; i< 10; i+=0 {
  for i:=0; i< 10; i+=0 {

    // path, err := checkFiles()
    // if err != nil {
    //   fmt.Println(err)
    //   break
    // }
		var siaData Data
    siaData, err := downloadFile()
    if err != nil {
      fmt.Println(err)
      break
    }

    err = sendToGraph(GRAPH_FILE_LOCATION, siaData)
    if err != nil {
      fmt.Println(err)
      break
    }

    // err = deleteFile(path)
    // if err != nil {
    //   fmt.Println(err)
    //   break
    // }
    time.Sleep(30 * time.Second)
  }
  time.Sleep(5 * time.Second)
  }


}

func checkFiles() (string, error) {
  var files siaFiles
  fmt.Println("Checking For Files on Sia")

  url := "http://localhost:9980/renter/files"
  request, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return "", err
  }
  request.Header.Set("User-Agent","Sia-Agent")

  client := &http.Client{}
  response, err := client.Do(request)
  if err != nil {
    return "", err
  }
  defer response.Body.Close()

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return "", err
  }
  data := bytes.TrimSpace(body)

  data = bytes.TrimPrefix(data, []byte("// "))

  err = json.Unmarshal(data, &files)
  if err != nil {
    return "", err
  }

  // fmt.Println(files.Files == nil)
  if files.Files != nil {
    if files.Files[0].SiaPath == "data.txt" {
      if files.Files[0].Progress == 100 {
        return files.Files[0].SiaPath, nil
      } else {
        newErr := errors.New("File Not Fully Uploaded Yet")
        return "", newErr
      }
    }
    newErr := errors.New("No data.txt found")
    return "", newErr
  }
  newErr := errors.New("No Files Found")
  return "", newErr
}

func downloadFile() (Data, error) {
	var siaData Datas
	var fakeData Data
  fmt.Println("Downloading Files From Sia")

  dir, err := filepath.Abs("./")
  if err != nil {
    return fakeData, err
  }
  dir = strings.Replace(dir, "/", "%2F", -1)

	// url := "http://newmont.io4.in:8080/data"
	url := "http://localhost:8080/data"

  request, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return fakeData, err
  }
  request.Header.Set("User-Agent","Sia-Agent")

  client := &http.Client{}
  response, err := client.Do(request)
  if err != nil {
    return fakeData, err
  }
  defer response.Body.Close()

  body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fakeData, err
	}
	data := bytes.TrimSpace(body)

	data = bytes.TrimPrefix(data, []byte("// "))

	err = json.Unmarshal(data, &siaData)
	if err != nil {
    if err.Error() != "unexpected end of JSON input" {
      return fakeData, err
    }
  }

	fmt.Println(siaData[len(siaData)-1])

	if len(siaData) != 0 {
		return siaData[len(siaData)-1], nil
	}
	newErr := errors.New("No Data on Server Yet")
	return fakeData, newErr

}

func sendToGraph(csvLocation string, siaData Data) (error) {
  // var pulledData fileData

  // inFile, err := os.Open("./data.txt")
  // if err != nil {
  //   return err
  // }
	//
  // scanner := bufio.NewScanner(inFile)
	// scanner.Split(bufio.ScanLines)
	//
  // scanner.Scan()                  //line 1
  // text := scanner.Text()
  // inFile.Close()
  // // fmt.Println(text)
  // json.Unmarshal([]byte(text), &pulledData)
  // if err != nil {
  //   return err
  // }

  f, err := os.OpenFile(csvLocation, os.O_APPEND|os.O_WRONLY, 0644)
  csvString := siaData.Time + "," + strconv.Itoa(siaData.Temperature) + "," + strconv.Itoa(siaData.Humidity) + "\n"
  _, err = f.WriteString(csvString)
  // fmt.Println(n)
  if err != nil {
    fmt.Println(err)
  }

  f.Close()

  return nil
}

func deleteFile(siaPath string) (error) {
  fmt.Println("Deleting File")
  var siaResp generalResp

  url := "http://localhost:9980/renter/delete/" + siaPath

  request, err := http.NewRequest("POST", url, nil)
  if err != nil {
    return err
  }
  request.Header.Set("User-Agent","Sia-Agent")

  client := &http.Client{}
  response, err := client.Do(request)
  if err != nil {
    return err
  }
  defer response.Body.Close()

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return err
  }
  data := bytes.TrimSpace(body)

  data = bytes.TrimPrefix(data, []byte("// "))

  err = json.Unmarshal(data, &siaResp)
  if err != nil {
    if err.Error() != "unexpected end of JSON input" {
      return err
    }
  }

  fmt.Println(siaResp)

  return nil
}

func startup(csvLocation string) {
  csvHeaders := []byte("Date,Temperature,Humidity\n")
  err := ioutil.WriteFile(csvLocation, csvHeaders, 0644)
  if err != nil {
    fmt.Println(err)
  }

}
