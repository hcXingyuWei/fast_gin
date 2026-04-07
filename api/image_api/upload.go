package image_api

import (
	"fast_gin/global"
	"fast_gin/utils/find"
	"fast_gin/utils/md5"
	"fast_gin/utils/random"
	"fast_gin/utils/res"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

var WhiteList = []string{
	".jpg",
	".jpeg",
	".png",
	".webp",
}

func (ImageApi) UploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	//未上传
	if err != nil {
		res.FailWithMig("请选择文件", c)
		return
	}

	//文件过大
	if fileHeader.Size > global.Config.Upload.Size*1024*1024 {
		res.FailWithMig("上传文件过大", c)
		return
	}

	//后缀判断
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !find.InList(WhiteList, ext) {
		res.FailWithMig("上传文件类型非法", c)
		return
	}

	//处理文件名重复
	filePath := path.Join("uploads", global.Config.Upload.Dir, fileHeader.Filename)
	for {
		_, err1 := os.Stat(filePath)
		if os.IsNotExist(err1) {
			break
		}
		//文件存在 算上传文件图片和本身图片是否一样，如果一样，直接返回原来的地址
		uploadFile, _ := fileHeader.Open()
		oldFile, _ := os.Open(filePath)
		uploadFileHash := md5.MD5WithFile(uploadFile)
		oldFileHash := md5.MD5WithFile(oldFile)
		if uploadFileHash == oldFileHash {
			//上传的图片， 名称和内容一样
			res.Ok("/"+filePath, "上传成功", c)
			return
		}
		//上传的图片，名称一样，内容不一样
		fileNameNotExt := strings.TrimSuffix(fileHeader.Filename, ext)
		newFileName := fmt.Sprintf("%s_%s%s", fileNameNotExt, random.RandStr(3), ext)
		filePath = path.Join("uploads", global.Config.Upload.Dir, newFileName)
	}
	c.SaveUploadedFile(fileHeader, filePath)

	res.Ok("/"+filePath, "上传成功", c)
}
