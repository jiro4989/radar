package main

import (
	"io"
	"os"

	"github.com/ajstarks/svgo"
)

// Radar はレーダーチャート1つ分のデータ
type Radar struct {
	Title  string `json:"title"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Max    int    `json:"max"`
	Values []struct {
		Text  string `json:"text"`
		Value int    `json:"value"`
	} `json:"values"`
}

func main() {
	rad := Radar{
		Title:  "title",
		Width:  500,
		Height: 500,
	}
	WriteSVG(rad, os.Stdout)
}

func WriteSVG(rad Radar, r io.Writer) {
	w, h, title := rad.Width, rad.Height, rad.Title
	canvas := svg.New(r)
	canvas.Start(w, h)
	canvas.Circle(w/2, h/2, 100)
	canvas.Text(w/2, h/2, title, "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()
}
