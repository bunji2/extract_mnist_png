package main

// MNIST データを PNG 形式でファイルに保存する
// Usage: extract_mnist_png [ cols rows ]
// Example:
//   extract_mnist_png       ---->  40x20 にまとめて保存
//   extract_mnist_png 80 80 ---->  80x80 にまとめて保存
//   extract_mnist_png 1 1   ---->  個別に保存

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bunji2/extract_mnist_png/mnist"
)

const (
	imageFile = "dataset/train-images-idx3-ubyte.gz"
	labelFile = "dataset/train-labels-idx1-ubyte.gz"
	usage     = "Usage: %s [cols rows]\n"
)

func main() {
	os.Exit(run())
}

func run() int {

	imgs, _, err := mnist.LoadMnist(imageFile, labelFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	cols := 40
	rows := 20

	if len(os.Args) >= 3 {
		cols, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprintf(os.Stderr, usage, os.Args[0])
			return 2
		}
		rows, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprintf(os.Stderr, usage, os.Args[0])
			return 3
		}
		if cols < 1 || rows < 1 {
			fmt.Fprintln(os.Stderr, "cols and rows should be positive number")
			fmt.Fprintf(os.Stderr, usage, os.Args[0])
			return 4
		}
	}
	for i := 0; i < len(imgs); i += (rows * cols) {
		till := i + rows*cols
		if till >= len(imgs) {
			till = len(imgs)
		}
		mImgs := mnist.MakeMultiGray(cols, rows, imgs[i:till])
		mnist.SavePng(fmt.Sprintf("x%d.png", i), mImgs)
	}
	return 0
}
