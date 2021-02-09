package main

import (
	"encoding/json"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
)

func main() {
	err := convertToCode("src/341.png")
	if err != nil {
		panic(err)
	}
}

func convertToCode(fname string) (err error) {
	// rows := 5
	// cols := 19
	f, err := os.Open(fname)
	if err != nil {
		return
	}
	img, err := png.Decode(f)
	if err != nil {
		return
	}

	chars := ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` + "`" + `abcdefghijklmnopqrstuvwxyz{|}~`
	chari := -1

	type Position struct {
		X int `json:"x"`
		Y int `json:"y"`
	}
	type Glyph struct {
		Glyph     string     `json:"glyph"`
		Positions []Position `json:"positions"`
	}
	var glyphs = []Glyph{}

	for row := 0; row < 5; row++ {
		for col := 0; col < 19; col++ {
			chari++

			fmt.Printf("%s\n", string(chars[chari]))
			positions := []Position{}
			for y := 0; y < 4; y++ {
				for x := 0; x < 3; x++ {
					y0 := (row*5 + 1) + y
					x0 := (col*4 + 1) + x
					r, _, _, a := img.At(x0, y0).RGBA()
					val := 0
					if r == 0 && a == 65535 {
						val = 1
					}
					fmt.Println(x, y, val)
					if val == 1 {
						positions = append(positions, Position{x, y})
					}
				}
			}
			glyphs = append(glyphs, Glyph{
				string(chars[chari]),
				positions,
			})

		}
	}

	b, err := json.MarshalIndent(glyphs, "", " ")
	if err != nil {
		return
	}

	err = ioutil.WriteFile("glyphs.json", b, 0644)

	return
}
