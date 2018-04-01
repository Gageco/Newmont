package main

import (
	"log"
	"net/http"
  "fmt"
	"time"
	"github.com/wcharczuk/go-chart"
  // "math"
)

var repeat int

func getXVal() []time.Time {
  ret := []time.Time{}
  for i:=0; i<len(datas); i++ {
    ret = append(ret, datas[i].Time)
  }

  return ret
}

func getYVal() []float64 {
  ret := []float64{}
  for i:=0; i<len(datas); i++ {
    ret = append(ret, float64(datas[i].Data))
  }
  return ret
}

func drawChart(res http.ResponseWriter, req *http.Request) {
  xVal := getXVal()
  yVal := getYVal()
  repeat+=1
  graph := chart.Chart{
  		XAxis: chart.XAxis{
  			Style: chart.Style{
  				Show: true, //enables / displays the x-axis
  			},
				ValueFormatter: chart.TimeHourValueFormatter,
  		},
  		YAxis: chart.YAxis{
  			Style: chart.Style{
  				Show: true, //enables / displays the y-axis
  			},
  		},
  		Series: []chart.Series{
  			chart.TimeSeries{
  				Style: chart.Style{
  					Show:        true,
  					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
  					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
  				},
  				XValues: xVal,
  				YValues: yVal,
  			},
  		},
  	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

// func drawChartWide(res http.ResponseWriter, req *http.Request) {
// 	graph := chart.Chart{
// 		Width: 1920, //this overrides the default.
// 		Series: []chart.Series{
// 			chart.ContinuousSeries{
// 				XValues: []float64{1.0, 2.0, 3.0, 4.0},
// 				YValues: []float64{1.0, 4.0, 9.0, 16.0},
// 			},
// 		},
// 	}
//
// 	res.Header().Set("Content-Type", "image/png")
// 	graph.Render(chart.PNG, res)
// }

func deployGraph() {
	fmt.Println("Graph Deployed at localhost:8081")
	http.HandleFunc("/", drawChart)
	// http.HandleFunc("/wide", drawChartWide)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
