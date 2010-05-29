package graph

import (
	svglib "svg"
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
)

var (
	width  int
	height int
	svg    *svglib.SVG
	colors = [3]string{"stroke:black;", "stroke:blue;", "stroke:red;"}
)

func background(v int) { svg.Rect(0, 0, width, height, svg.RGB(v, v, v)) }

func drawGraph(x, y []int, style ...string) {
	if len(x) != len(y) || len(x) == 0 || len(y) == 0 {
		return
	}
	x1, y1 := x[0], y[0]
	for i := 1; i < len(x); i++ {
		svg.Line(x1*100+100, height-100-y1/100, x[i]*100+100, height-100-y[i]/100, style)
		x1, y1 = x[i], y[i]
	}
}

func drawAxis(xrate, yrate int) {
	style := "stroke:black;"
	xlen := (width - xrate - 100) / xrate
	ylen := (height - yrate - 100) / yrate
	svg.Line(100, height-100, xlen*xrate+100, height-100, style)
	svg.Line(100, height-100, 100, height-100-(ylen*yrate), style)
	for i := 1; i <= xlen; i++ {
		svg.Line(i*xrate+100, height-100, i*xrate+100, height-90, style)
		svg.Text(i*xrate+90, height-80, fmt.Sprint(1000*i), style)
	}
	for i := 1; i <= ylen; i++ {
		svg.Line(100, height-100-(i*yrate), 90, height-100-(i*yrate), style)
		svg.Text(60, height-95-(i*yrate), fmt.Sprint(i * 1000), style)
	}
}

type Values struct {
	X, Y []int
}

type Graph struct {
	values     []Values
	xlen, ylen int
}

func NewGraph(xlen, ylen int, values []Values) *Graph {
	return &Graph{values, xlen, ylen}
}

func (this *Graph) Draw() {
	drawAxis(this.xlen, this.ylen)
	for i, v := range this.values {
		drawGraph(v.X, v.Y, colors[i])
	}
}

func DrawToFile(file string, graph *Graph) {
	width, height = 800, 800
	buffer := bytes.NewBuffer(nil)//create an empty bytes buffer
	svg = svglib.New(buffer)//initialize svg with buffer
	svg.Start(width, height)
	background(255)

	svg.Title("CSC 341 Assignment")

	svg.Text(120, 30, "Comparison of Selection Sort, Tree Sort, Quick Sort", colors[0]+"font-family:Calibri; font-size:30px;")
	svg.Text(10, 70, "Time (microseconds)", colors[0]+"font-family:Calibri")
	svg.Text(width-300, height-60, "Number of sorted elements", colors[0]+"font-family:Calibri")

	svg.Grid(0, 0, width, height, 10, "stroke:black;opacity:0.15")

	graph.Draw()
	
	DrawTable(graph.values)

	svg.End()
	
	ioutil.WriteFile(file, buffer.Bytes(), 0666)//write the buffer to file
}

func DrawTable(values []Values){
	//svg.Text(500,200, "Insertion : Blue, Quick : Red, Tree : Black", colors[0]) tree, insertion, quick
	svg.Text(500,200, "Tree Sort", colors[0]+"font-family:Calibri")
	svg.Line(580, 195, 640, 195, colors[0])
	svg.Text(500,220, "Insertion Sort", colors[0]+"font-family:Calibri")
	svg.Line(580, 215, 640, 215, colors[1])
	svg.Text(500,240, "Quick Sort", colors[0]+"font-family:Calibri")
	svg.Line(580, 235, 640, 235, colors[2])
	
	names := []string{"Tree Sort", "Insertion Sort", "Quick Sort"}
	startWidth, startHeight, w, h := 500, 300, 80, 20
	svg.Text(startWidth, startHeight, "No of Data", colors[0]+"font-family:Calibri")
	for i, n := range names {
		svg.Text(startWidth+w*(i+1), startHeight, n, colors[0]+"font-family:Calibri")
	}
	for i:=1; i < 7; i++ {
		svg.Text(startWidth, startHeight+h*(i), fmt.Sprint(1000 * i), colors[0]+"font-family:Calibri;font-size:14px;")
	}
	for i, val := range values {
		for j, y := range val.Y[1:]{
			svg.Text(startWidth+w*(i+1), startHeight+h*(j+1), fmt.Sprint(y), colors[0]+"font-family:Calibri;font-size:14px;")
		}
	}
}

func Pow(x, y int) int{
	return int(math.Pow(float64(x), float64(y)))
}
