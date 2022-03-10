package main

import (
	"flag"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/mishazawa/printbot/parser"
	"github.com/mishazawa/printbot/printer"
)

var (
	filePath string
	x, y, z  int
	remove   bool
	host     string
	port     string
	password string
)

func init() {
	flag.StringVar(&filePath, "file", "", "File path.")
	flag.IntVar(&x, "x", 0, "Horizontal axis.")
	flag.IntVar(&y, "y", 0, "Vertical axis.")
	flag.IntVar(&z, "z", 0, "Depth axis.")
	flag.BoolVar(&remove, "remove", false, "Remove picture from desired coords.")
	flag.StringVar(&host, "host", os.Getenv("RCON_HOST"), "RCON server ip.")
	flag.StringVar(&port, "port", os.Getenv("RCON_PORT"), "RCON server port.")
	flag.StringVar(&password, "password", os.Getenv("RCON_PASSWORD"), "RCON server password.")
	flag.Parse()
}

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	ext := filepath.Ext(filePath)

	if ext == ".jpg" || ext == ".jpeg" {
		if err := processImage(file); err != nil {
			panic(err.Error())
		}
	}

	if ext == ".obj" {
		if err := processModel(file); err != nil {
			panic(err.Error())
		}
	}

}

func processImage(file *os.File) error {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	builder := printer.NewPrinter(host+":"+port, password)
	defer builder.Close()

	if remove {
		if err := builder.Remove(width, height, x, y, z); err != nil {
			return err
		}
	} else {
		if err := builder.Draw(img, x, y, z); err != nil {
			return err
		}
	}

	return nil
}

func processModel(file *os.File) error {
	voxels, err := parser.Parse(file)

	if err != nil {
		return err
	}

	builder := printer.NewPrinter(host+":"+port, password)
	defer builder.Close()

	if err := builder.Print(voxels, x, y, z, remove); err != nil {
		return err
	}

	return nil
}
