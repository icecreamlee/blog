package controllers

import (
	"blog/internal/core"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func UEditor(c *gin.Context) {
	action := c.Query("action")
	switch action {
	case "config":
		c.Header("Content-type", "application/json; charset=utf-8")
		c.String(200, CONFIG_JSON)
		return
	case "uploadimage":
		c.Header("Content-type", "application/json; charset=utf-8")
		result := upload(core.RootPath+"web/statics/ueditor/upload/image/", c.Request)
		result["url"] = "/static" + strings.TrimPrefix(result["url"], core.RootPath+"web/statics")
		c.JSON(200, result)
		return
	}
}

func newFileName() string {
	var now = time.Now()
	var fileName = now.Format("200601021504")
	fileName = fileName + strconv.Itoa(now.Nanosecond())
	return fileName
}

func upload(pathString string, request *http.Request) map[string]string {
	var file, header, err = request.FormFile("upfile")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var filename = newFileName() + path.Ext(header.Filename)
	err = os.MkdirAll(pathString, 0775)
	if err != nil {
		panic(err)
	}
	targetFile, err := os.Create(path.Join(pathString, filename))
	if err != nil {
		panic(err)
	}
	defer targetFile.Close()
	io.Copy(targetFile, file)
	return map[string]string{
		"url":      pathString + filename,
		"title":    "",
		"original": header.Filename,
		"state":    "SUCCESS",
	}
}

const CONFIG_JSON = `{
  "imageActionName": "uploadimage",
  "imageUrlPrefix": "",
  "imageFieldName": "upfile",
  "imageMaxSize": 20480000,
  "imageAllowFiles": [
    ".png",
    ".jpg",
    ".jpeg",
    ".gif",
    ".bmp"
  ],
  "imageManagerActionName": "listimage",
  "imageManagerListSize": 20,
  "imageManagerUrlPrefix": "",
  "imageManagerInsertAlign": "none",
  "imageManagerAllowFiles": [
    ".png",
    ".jpg",
    ".jpeg",
    ".gif",
    ".bmp"
  ],
  "fileActionName": "uploadfile",
  "fileFieldName": "upfile",
  "fileUrlPrefix": "",
  "fileMaxSize": 51200000,
  "fileAllowFiles": [
    ".png",
    ".jpg",
    ".jpeg",
    ".gif",
    ".bmp",
    ".flv",
    ".swf",
    ".mkv",
    ".avi",
    ".rm",
    ".rmvb",
    ".mpeg",
    ".mpg",
    ".ogg",
    ".ogv",
    ".mov",
    ".wmv",
    ".mp4",
    ".webm",
    ".mp3",
    ".wav",
    ".mid",
    ".rar",
    ".zip",
    ".tar",
    ".gz",
    ".7z",
    ".bz2",
    ".cab",
    ".iso",
    ".doc",
    ".docx",
    ".xls",
    ".xlsx",
    ".ppt",
    ".pptx",
    ".pdf",
    ".txt",
    ".md",
    ".xml"
  ],
  "fileManagerActionName": "listfile",
  "fileManagerUrlPrefix": "",
  "fileManagerListSize": 20,
  "fileManagerAllowFiles": [
    ".png",
    ".jpg",
    ".jpeg",
    ".gif",
    ".bmp",
    ".flv",
    ".swf",
    ".mkv",
    ".avi",
    ".rm",
    ".rmvb",
    ".mpeg",
    ".mpg",
    ".ogg",
    ".ogv",
    ".mov",
    ".wmv",
    ".mp4",
    ".webm",
    ".mp3",
    ".wav",
    ".mid",
    ".rar",
    ".zip",
    ".tar",
    ".gz",
    ".7z",
    ".bz2",
    ".cab",
    ".iso",
    ".doc",
    ".docx",
    ".xls",
    ".xlsx",
    ".ppt",
    ".pptx",
    ".pdf",
    ".txt",
    ".md",
    ".xml"
  ]
}
`
