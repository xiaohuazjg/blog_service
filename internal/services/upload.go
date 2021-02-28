package services

import (
	"errors"
	"mime/multipart"
	"os"

	"github.com/xiaohuazjg/blog_service/global"
	"github.com/xiaohuazjg/blog_service/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType,
	file multipart.File,
	fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not support")

	}

	if upload.CheckMaxSize(fileType, fileName) {
		return nil, errors.New("exceeded maixmum file limit")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}
	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileName, dst); err != nil {
		return nil, err
	}
	accessurl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{
		Name:      fileName,
		AccessUrl: accessurl,
	}, nil
}
