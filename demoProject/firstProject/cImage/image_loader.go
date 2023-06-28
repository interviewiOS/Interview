package cImage

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	imgext "github.com/shamsher31/goimgext"
	"image"
	_ "image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type ImageType int

const (
	ImageType_jpeg ImageType = iota
	ImageType_png
	ImageType_gif
	ImageType_unknown
)

// String type类型转换成字符串
func (f ImageType) String() [4]string {
	return [...]string{"ImageType_jpeg", "ImageType_png", "ImageType_gif", "ImageType_unknown"}
}

//func (i ImageType) String() string {
//	switch i {
//	case ImageType_jpeg:
//		return "jpg"
//	case ImageType_png:
//		return "png"
//	case ImageType_gif:
//		return "gif"
//	case ImageType_unknown:
//		return "unknown"
//	default:
//		return "unknown"
//	}
//}

type OneColourType int

const (
	OneColourType_true OneColourType = iota
	OneColourType_false
	OneColourType_unknown
)

// String type类型转换成字符串
//func (f OneColourType) String() [3]string {
//	return [...]string{"纯色", "非纯色", "未知x"}
//}

func (o OneColourType) String() string {
	switch o {
	case OneColourType_true:
		return "纯色"
	case OneColourType_false:
		return "非纯色"
	case OneColourType_unknown:
		return "未知"
	default:
		return "未知x"
	}
}

const cachPath = "./imageFile/cache/"
const csvPath = "./imageFile/csv/"
const csvTotalCacheFileName = "oneColour.csv"
const csvOnceCacheFileName = "oneColourList.csv"

// GetImageTYPE 返回图片类型image/png imageFile/gif ...
func GetImageTYPE(p string) (ImageType, error) {
	file, err := os.Open(p)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	buff := make([]byte, 512)

	_, err = file.Read(buff)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	return getImageTYPEByByte(buff)

}

func LoadIamgeTest() {

	//modelMap := readCacheCsv()
	//if len(modelMap) > 0 {
	//	imageOneColourCache = modelMap
	//}

	list := []string{
		//"https://puep.qpic.cn/qqgameedu/0/096910ec78d802c8a6b0681dcc5ca4fe_0/0",
		//"https://puep.qpic.cn/qqgameedu/0/fd4ab7736559f096cd5f76f27d74a07e_0/0",
		//"https://puep.qpic.cn/qqgameedu/0/20762fe5397b78f6490474fcff7e3413_0/0",
		"https://photo.16pic.com/00/13/89/16pic_1389821_b.jpg",
		"https://www.pngfind.com/pngs/m/470-4703547_icon-png-download-imageFile-transparent-png.png",
		"https://www.pngfind.com/pngs/m/470-4703547_icon-png-download-imageFile-transparent-png.png",
		"https://hbimg.huaban.com/32f065b3afb3fb36b75a5cbc90051b1050e1e6b6e199-Ml6q9F_fw320",
	}
	var oneCount, noCount, unCount int
	var repeatMap = make(map[string]int)
	for _, imgUrl := range list {
		sum, _ := repeatMap[imgUrl]
		repeatMap[imgUrl] = sum + 1
		oneResult, _ := AssertImageWithUrl(imgUrl)
		switch oneResult {
		case OneColourType_true:
			oneCount++
		case OneColourType_false:
			noCount++
		case OneColourType_unknown:
			unCount++
		}
	}
	fmt.Printf("总检查次数：%d\n", count)
	fmt.Printf("图片总数：%d,单色图片数：%d,非单色图片数：%d,未知：%d\n", len(list), oneCount, noCount, unCount)
	fmt.Printf("总url：%+v\n", repeatMap)
	writeRepeatMapToCsv(repeatMap)
	count = 0
	writeAllCacheToCsv()

}

func writeRepeatMapToCsv(repeatMap map[string]int) {
	if len(repeatMap) > 0 {
		var dataList [][]string
		for key, value := range repeatMap {
			dataList = append(dataList, []string{
				key,
				strconv.Itoa(value),
			})
		}
		titleList := []string{
			"url",
			"出现次数",
		}

		timestamp := fmt.Sprintf("%d", time.Now().Unix())
		csvPath := csvPath + "repeat" + timestamp + ".csv"
		WriteCSVByDList(titleList, dataList, csvPath)
	}
}

// getImageTYPEByByte 传入[]byte类型数据 返回图片类型
func getImageTYPEByByte(bty []byte) (ImageType, error) {
	length := len(bty)
	var buff = make([]byte, length)
	copy(buff, bty)
	filetype := http.DetectContentType(buff)
	ext := imgext.Get()
	for i := 0; i < len(ext); i++ {
		if strings.Contains(ext[i], filetype[6:len(filetype)]) {
			fmt.Println("图片类型", filetype)
			return swichImageType(filetype), nil
		}
	}
	fmt.Println("图片类型xx", filetype)
	return ImageType_unknown, errors.New("Invalid imageFile type")

}

func getImageTypeByImageUrl(imageUrl string) (ImageType, error) {
	url := imageUrl
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ImageType_unknown, err
	}
	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	imageType := swichImageType(contentType)
	return imageType, nil
}

func swichImageType(strType string) ImageType {
	if strings.Contains(strType, "jpg") || strings.Contains(strType, "jpeg") {
		return ImageType_jpeg
	} else if strings.Contains(strType, "gif") {
		return ImageType_gif
	} else if strings.Contains(strType, "png") {
		return ImageType_png
	} else {
		return ImageType_unknown
	}
}

var imageOneColourCache = make(map[string]ImageUrlModel)
var count = 0

// AssertImageWithUrl 判断url的类型
func AssertImageWithUrl(imgUrl string) (OneColourType, ImageType) {
	if len(imgUrl) <= 0 {
		fmt.Println("imgurl为空")
		return OneColourType_unknown, ImageType_unknown
	}

	result, succ := imageOneColourCache[imgUrl]
	cacheResult := result.ColorType
	if succ {
		if cacheResult != OneColourType_unknown {
			return cacheResult, ImageType_unknown
		}
	}
	count++
	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!", err)
		return OneColourType_unknown, ImageType_unknown
	}
	defer res.Body.Close()

	//reader := bufio.NewReader(res.Body)
	//bytess, _ := io.ReadAll(reader)
	//imageType, _ := getImageTYPEByByte(bytess)
	//contentType := res.Header.Get("Content-Type")
	//imageType := swichImageType(contentType)
	//success, errorss := oneColourImage(res.Body.(io.Reader), imageType)
	//if errorss != nil {
	//	return OneColourType_unknown
	//}

	//获得get请求响应的reader对象
	reader := bufio.NewReader(res.Body)
	//获得图片格式 将图片保存在本地
	fileName := path.Base(imgUrl)

	imgPath := cachPath + fileName
	error1 := writeBodyToLocFile(reader, imgPath)
	if error1 != nil {
		fmt.Println("保存文件失败", error1)
		return OneColourType_unknown, ImageType_unknown
	}
	imageType, errType := GetImageTYPE(imgPath)
	if errType != nil {
		return OneColourType_unknown, ImageType_unknown
	}
	//判断图片是否是纯色
	success, err2 := OneColourImageByPathAndType(imgPath, imageType)
	if err2 != nil {
		fmt.Println("图片判断失败", err2)
		return OneColourType_unknown, imageType
	}

	if success {
		imageOneColourCache[imgUrl] = ImageUrlModel{
			Url:       imgUrl,
			ImgType:   imageType,
			ColorType: OneColourType_true,
		}
		return OneColourType_true, imageType
	} else {
		imageOneColourCache[imgUrl] = ImageUrlModel{
			Url:       imgUrl,
			ImgType:   imageType,
			ColorType: OneColourType_false,
		}
		return OneColourType_false, imageType
	}
	return OneColourType_unknown, imageType
}

func writeBodyToLocFile(reader io.Reader, path string) error {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("A error occurred!", err)
		return err
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
	return nil
}

// oneColourImage 判断图片是否是纯色图片
func oneColourImage(reader io.Reader, imageType ImageType) (bool, error) {
	var img image.Image
	var err error
	//img, _, err = image.Decode(reader)
	if imageType == ImageType_jpeg {
		img, err = jpeg.Decode(reader)
	} else if imageType == ImageType_png {
		img, err = png.Decode(reader)
	} else if imageType == ImageType_gif {
		img, err = gif.Decode(reader)
	}
	if err != nil {
		fmt.Println("图片解码失败", err)
		return false, err
	}

	// 获取第一个像素点的颜色
	color1 := img.At(0, 0)

	// 遍历所有像素点，判断是否为纯色图片
	isPureColor := true
	var count = 0
	dx := img.Bounds().Dx()
	dy := img.Bounds().Dy()
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			count += 1
			if !ColorEqual(color1, img.At(x, y)) {
				isPureColor = false
				break
			}
		}
		if !isPureColor {
			break
		}
	}
	fmt.Println("统计次数", count)

	// 输出结果
	if isPureColor {
		fmt.Println("图片是纯色图片")
	} else {
		fmt.Println("图片不是纯色图片")
	}
	return isPureColor, nil
}

// OneColourImageByPath 传入路径判断图片是否是纯色
func OneColourImageByPath(path string) (bool, error) {
	imagePath := path
	if len(imagePath) <= 0 {
		return false, errors.New("传入的path为空")
	}
	imageType, error := GetImageTYPE(imagePath)
	if error != nil {
		fmt.Println("获得imgtype失败", error)
		return false, error
	}
	return OneColourImageByPathAndType(imagePath, imageType)
}

// OneColourImageByPathAndType 传入路径判断图片是否是纯色
func OneColourImageByPathAndType(iPath string, imageType ImageType) (bool, error) {
	imagePath := iPath
	if len(imagePath) <= 0 {
		return false, errors.New("传入的path为空")
	}
	byt, err := os.ReadFile(imagePath)
	if err != nil {
		fmt.Sprintln("errror:", err)
		return false, err
	}
	file := bytes.NewReader(byt)
	//如果能判断出图片类型在后边加后缀方便人工校验
	if imageType != ImageType_unknown && !strings.Contains(path.Base(imagePath), ".") {
		rpath := fmt.Sprintf("%s.%+v", imagePath, imageType)
		error := os.Rename(imagePath, rpath)
		print(error)
	}
	rbyt := file.Len()
	print(rbyt)
	return oneColourImage(file, imageType)
}

// 判断两个颜色是否相等
func ColorEqual(color color.Color, eColor color.Color) bool {
	r, g, b, a := color.RGBA()
	er, eg, eb, ea := eColor.RGBA()
	return r == er &&
		g == eg &&
		b == eb &&
		a == ea
}

func writeAllCacheToCsv() {
	filePath := csvPath + csvTotalCacheFileName
	writeCacheToCsv(filePath, imageOneColourCache)
}

func clearCache() {
	imageOneColourCache = make(map[string]ImageUrlModel)
	os.RemoveAll(cachPath)
}

func writeCacheToCsv(filePath string, cacheMap map[string]ImageUrlModel) {
	var dataList [][]string
	for _, value := range cacheMap {
		itemList := []string{
			value.Url,
			fmt.Sprintf("%+v", value.ImgType),
			fmt.Sprintf("%+v", value.ColorType),
		}
		dataList = append(dataList, itemList)
	}
	if len(dataList) > 0 {
		titleList := []string{
			"域名",
			"图片类型",
			"是否是单色图片",
		}
		WriteCSVByDList(titleList, dataList, filePath)
	}
}

func readCacheCsv() map[string]ImageUrlModel {
	os.Mkdir(cachPath, os.ModePerm)
	os.Mkdir(csvPath, os.ModePerm)
	filePath := csvPath + csvTotalCacheFileName
	file, error := os.Open(filePath)
	if error != nil {
		fmt.Println("Error while opening file:", error)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	recorder, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error while reading CSV data:", err)
		return nil
	}
	modelDic := make(map[string]ImageUrlModel)
	for _, row := range recorder {
		imgType, error1 := strconv.Atoi(row[1])
		colour, error2 := strconv.Atoi(row[2])
		if error1 != nil || error2 != nil {
			continue
		}
		modelDic[row[0]] = ImageUrlModel{
			Url:       row[0],
			ImgType:   ImageType(imgType),
			ColorType: OneColourType(colour),
		}
	}
	return modelDic
}

// WriteCSVByDList 最简单的双重输出
func WriteCSVByDList(title []string, writeList [][]string, path string) {
	if len(writeList) <= 0 {
		return
	}
	//OpenFile读取文件，不存在时则创建，使用覆盖模式
	csvFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("error, file open failed, path = %+v\n!", path)
		return
	}

	// 写入UTF-8 BOM，防止中文乱码
	csvFile.WriteString("\xEF\xBB\xBF")

	//创建写入接口
	writerCsv := csv.NewWriter(csvFile)
	if len(title) > 0 {
		writerCsv.Write(title)
	}

	writerCsv.WriteAll(writeList)
	writerCsv.Flush()
}

type ImageUrlModel struct {
	Url       string
	ImgType   ImageType
	ColorType OneColourType
}
