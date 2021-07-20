package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"
)

// Coordinates represents coordinates on a cartesian plane
type Coordinates struct {
	X, Y int
}

// Arrow is an arrow
type Arrow struct {
	Start, Finish    Coordinates
	StartX, StartY   int
	FinishX, FinishY int
	Width            float32
}

// Canvas represent SVG canvas
type Canvas struct {
	Width, Height int
}

func main() {
	svgBaseTmplData, err := ioutil.ReadFile("windrose_base.svg.tmpl")
	if err != nil {
		panic(err)
	}
	svgArrowTmplData, err := ioutil.ReadFile("windrose_arrow.svg.tmpl")
	if err != nil {
		panic(err)
	}
	svgArrowTmpl, err := template.New("arrow").Parse(string(svgArrowTmplData))
	if err != nil {
		panic(err)
	}
	svgBaseTmpl, err := template.New("arrow").Parse(string(svgBaseTmplData))
	if err != nil {
		panic(err)
	}
	/*
		double angle_rad = (angle - 90) * (M_PI / 180.0);
		fprintf(stderr, "angle supplied: %f\n", angle);

		struct Coordinates c_start;
		c_start.x = (CANVAS_SIZE / 2);
		c_start.y = (CANVAS_SIZE / 2);
		struct Coordinates c_finish;
		c_finish.x = c_start.x + (cos(angle_rad) * ARROW_LENGTH);
		c_finish.y = c_start.y + (sin(angle_rad) * ARROW_LENGTH);
		struct Canvas canvas;
		canvas.width = CANVAS_SIZE;
		canvas.height = CANVAS_SIZE;
	*/

	arrow := Arrow{}
	arrow.StartX = 306
	arrow.StartY = 300
	arrow.FinishX = 415
	arrow.FinishY = 23
	arrow.Width = 2.5

	svgArrowBuf := &bytes.Buffer{}
	err = svgArrowTmpl.Execute(svgArrowBuf, arrow)
	if err != nil {
		panic(err)
	}

	err = svgBaseTmpl.Execute(os.Stdout, svgArrowBuf.String())
	if err != nil {
		panic(err)
	}
}
