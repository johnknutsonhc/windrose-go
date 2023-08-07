package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

func windrose(w http.ResponseWriter, req *http.Request) {
	angle := req.URL.Query()["angle"][0]
	if angle == "" {
		angle = "0" // set a default
	}
	angleDeg, err := strconv.ParseFloat(angle, 64)
	if err != nil {
		angle = "0"
		err = nil
		// TODO: handle errors better
		// panic(err)
	}
	svgWindroseBuf := &bytes.Buffer{}
	err = GenWindrose(angleDeg, svgWindroseBuf)
	if err != nil {
		// TODO: what here?
		panic(err)
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	// w.Header().Set("Windrose-Angle-Deg", string(angleDeg))
	fmt.Fprintf(w, svgWindroseBuf.String())
}

func main() {
	http.HandleFunc("/windrose", windrose)

	http.ListenAndServe(":8090", nil)
}
