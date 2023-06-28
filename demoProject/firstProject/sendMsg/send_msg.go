package sendMsg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

const robotKey = "c4a79450-c10c-4549-907a-d746c220e286"

type FileContent struct {
	Msgtype string `json:"msgtype"`
	File    struct {
		MediaId string `json:"media_id"`
	} `json:"file"`
}

func sendFileMsg(media_id string) string {
	model := FileContent{
		Msgtype: "file",
		File: struct {
			MediaId string `json:"media_id"`
		}{
			MediaId: media_id,
		},
	}
	btyBody, error := json.Marshal(model)
	if error != nil {
		fmt.Println(error)
		return ""
	}

	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + robotKey

	clint := &http.Client{}

	reader := strings.NewReader(string(btyBody))
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := clint.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	bodyxx, err := io.ReadAll(resp.Body)
	fmt.Sprintf("result:%+v", bodyxx)

	str := string(bodyxx)

	fmt.Println("result:", str)

	return str
}

const sendKey = "c4a79450-c10c-4549-907a-d746c220e286"

func uploadFile(filePath string) string {
	//key := sendKey
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media?key=c4a79450-c10c-4549-907a-d746c220e286&type=file"
	oriPath := filePath
	dstPath := oriPath
	//原始路径和上传的文件路径是否相同
	var oriEquelDstPath = true
	if oriPath != dstPath {
		fileName := path.Base(oriPath)
		if len(fileName) > 0 && strings.Contains(oriPath, "/") {
			dstPath = fileName
			oriEquelDstPath = false
		}
		copyFile(dstPath, oriPath)
	}

	file, err := os.Open(dstPath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("upload", dstPath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	request.Header.Set("Content-Type", "multipart/form-data")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(responseBody))
	rModel := MediaModel{}
	json.Unmarshal(responseBody, &rModel)

	if !oriEquelDstPath {
		os.RemoveAll(dstPath)
	}

	return rModel.MediaId
}

type MediaModel struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

func SendFileMsg(filePath string) {
	mId := uploadFile(filePath)
	sendFileMsg(mId)
}

func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
