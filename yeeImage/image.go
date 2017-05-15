/**
 * Created by angelina on 2017/5/15.
 */

package yeeImage

import (
	"image"
	"strings"
	"os"
	"path"
	"image/png"
	"image/jpeg"
	"image/gif"
	"github.com/nfnt/resize"
)

// LoadImage
// 从文件中解码出image
func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

// SaveImage
// 将image保存到文件中，根据后缀判断保存形式，如果后缀不匹配，则png
func SaveImage(p string, img image.Image) (err error) {
	imgFile, err := os.Create(p)
	if err != nil {
		return err
	}
	defer imgFile.Close()
	switch strings.ToLower(path.Ext(p)) {
	case ".png":
		err = png.Encode(imgFile, img)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(imgFile, img, nil)
	case ".gif":
		err = gif.Encode(imgFile, img, nil)
	default:
		err = png.Encode(imgFile, img)
	}
	return
}

// ResizeImage
// 图片缩放
func ResizeImage(inputImagePath string, width, height uint, outputFileName string) error {
	inputImage, err := LoadImage(inputImagePath)
	if err != nil {
		return err
	}
	outputImage := resize.Resize(width, height, inputImage, resize.Lanczos3)
	err = SaveImage(outputFileName, outputImage)
	if err != nil {
		return err
	}
	return nil
}
