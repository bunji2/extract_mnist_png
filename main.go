package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	mnist "github.com/petar/GoMNIST"
)

const (
	imageFile = "dataset/train-images-idx3-ubyte.gz"
	labelFile = "dataset/train-labels-idx1-ubyte.gz"
)

func main() {
	os.Exit(run())
}

func run() int {

	imgs, labels, err := loadMnist(imageFile, labelFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	for i, img := range imgs {
		label := labels[i]
		savePng(fmt.Sprintf("x%06d_%d.png", i, label), img)
	}
	return 0
}

func loadMnist(imageFile, labelFile string) (r []*image.Gray, l []mnist.Label, err error) {
	var data *mnist.Set
	data, err = mnist.ReadSet(imageFile, labelFile)
	if err != nil {
		return
	}

	l = data.Labels

	numberImages := len(data.Images)
	row := data.NRow
	col := data.NCol

	rect := image.Rect(0, 0, row-1, col-1)
	r = make([]*image.Gray, numberImages)
	for i := 0; i < numberImages; i++ {
		r[i] = image.NewGray(rect)
		for x := 0; x < col; x++ {
			for y := 0; y < row; y++ {
				c := color.Gray{Y: data.Images[i][x+y*(col)]}
				r[i].SetGray(x, y, c)
			}
		}
	}
	return

}

func savePng(outFile string, img image.Image) (err error) {
	var f *os.File
	f, err = os.Create(outFile)
	if err != nil {
		return
	}
	defer f.Close()

	err = png.Encode(f, img)
	return
}
