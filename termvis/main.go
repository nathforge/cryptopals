package main

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/golang/freetype/truetype"
	flags "github.com/jessevdk/go-flags"

	"github.com/nathforge/cryptopals/termvis/ansievent"
	"github.com/nathforge/cryptopals/termvis/dispdrawer"
	"github.com/nathforge/cryptopals/termvis/display"
)

func main() {
	log.SetFlags(0)

	var opts struct {
		OutputPath     string  `long:"output-path" default:"."`
		OutputBasename string  `long:"output-basename" default:"frame"`
		FontFilename   string  `long:"font-filename" required:"true"`
		TermWidth      int     `long:"term-width" default:"80"`
		TermHeight     int     `long:"term-height" default:"24"`
		FontSize       float64 `long:"font-size" default:"14"`
	}
	args, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}
	if len(args) > 0 {
		log.Fatal(errors.New("Unexpected arguments"))
	}

	disp := display.New(opts.TermWidth, opts.TermHeight)

	fontFile, err := os.Open(opts.FontFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer fontFile.Close()
	fontBytes, err := ioutil.ReadAll(fontFile)
	if err != nil {
		log.Fatal(err)
	}
	font, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}
	regularFace := truetype.NewFace(font, &truetype.Options{Size: opts.FontSize})
	//boldFace := regularFace

	dispDrawer := dispdrawer.New(disp, regularFace, nil)
	dispDrawer.CellPaddingLeft = 0
	dispDrawer.CellPaddingTop = 1
	dispDrawer.CellPaddingRight = 0
	dispDrawer.CellPaddingBottom = 1

	frame := 0
	saveImage := func() {
		filename := path.Join(opts.OutputPath, fmt.Sprintf("%s%04d.png", opts.OutputBasename, frame))
		log.Printf("Writing %s\n", filename)

		img := image.NewRGBA(dispDrawer.ImageSize())
		dispDrawer.Draw(img)

		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := png.Encode(f, img); err != nil {
			log.Fatal(err)
		}
		frame++
	}
	disp.BeforeAnsiEvent = func(event interface{}) {
		if v, ok := event.(ansievent.ED); ok {
			if v.Param == 2 {
				saveImage()
			}
		}
	}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	err = disp.WriteBytes(bytes)
	if err != nil {
		log.Fatal(err)
	}

	saveImage()
}
