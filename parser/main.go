package parser

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

const SCALE = 25.0

type Voxel struct {
	X, Y, Z int
}

func Parse(f *os.File) ([]Voxel, error) {
	scanner := bufio.NewScanner(f)
	voxels := make([]Voxel, 0)

	for scanner.Scan() {
		txt := strings.Split(scanner.Text(), " ")

		if err := scanner.Err(); err != nil {
			return nil, err
		}

		v := txt[0]

		if v == "v" {
			x, _ := strconv.ParseFloat(txt[1], 32)
			y, _ := strconv.ParseFloat(txt[2], 32)
			z, _ := strconv.ParseFloat(txt[3], 32)
			voxels = append(voxels, Voxel{scaleVal(x), scaleVal(y), scaleVal(z)})
		}

	}
	return voxels, nil
}

func scaleVal(v float64) int {
	return int(math.Round(v * SCALE))
}
