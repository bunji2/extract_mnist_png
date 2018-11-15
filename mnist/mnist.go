package mnist

import (
	"image"
	"image/color"
	"image/png"
	"os"

	gmnist "github.com/petar/GoMNIST"
)

// LoadMnist : MNIST データを読み出す関数
func LoadMnist(imageFile, labelFile string) (r []*image.Gray, l []gmnist.Label, err error) {
	var data *gmnist.Set
	data, err = gmnist.ReadSet(imageFile, labelFile)
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

// MakeMultiGray : cols x rows にイメージを配置した新しいイメージを作成する関数
func MakeMultiGray(cols, rows int, imgs []*image.Gray) (r *image.Gray) {
	if cols == 1 && rows == 1 {
		r = imgs[0]
		return
	}

	bounds := imgs[0].Bounds()
	wUnit := bounds.Size().X
	hUnit := bounds.Size().Y
	r = image.NewGray(image.Rect(0, 0, wUnit*cols, hUnit*rows))

	for n, img := range imgs {
		i := n % cols
		j := n / cols
		if i >= cols || j >= rows {
			break
		}
		bx := i * wUnit
		by := j * hUnit
		for x := 0; x < wUnit; x++ {
			for y := 0; y < hUnit; y++ {
				c := img.GrayAt(x, y)
				r.SetGray(bx+x, by+y, c)
			}
		}
	}
	return
}

// SavePng : イメージデータを PNG 形式でファイルに保存する関数
func SavePng(outFile string, img image.Image) (err error) {
	var f *os.File
	f, err = os.Create(outFile)
	if err != nil {
		return
	}
	defer f.Close()

	err = png.Encode(f, img)
	return
}
