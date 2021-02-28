package upload

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/xiaohuazjg/blog_service/global"
	"github.com/xiaohuazjg/blog_service/pkg/util"
)

type FIleType int

const TypeImage FIleType = iota + 1

func GetFileName(name string) string {
	ext := GetFileExt(name)
	filename := strings.TrimSuffix(name, ext)
	filename = util.EncodeMD5(filename)
	return filename + ext
}

func GetFileExt(name string) string {
	return path.Ext(name)

}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath

}

func GetServerUrl() string {
	return global.AppSetting.UploadServerUrl
}

func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CheckContainExt(t FIleType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowext := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowext) == ext {
				return true
			}
		}
	}
	return false
}

func CheckMaxSize(t FIleType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

func CHeckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)

}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return err

}
