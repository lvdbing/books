package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path"

	"github.com/lvdbing/books/blog-service/global"
	"github.com/lvdbing/books/blog-service/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	filename := upload.GetFilename(fileHeader.Filename)
	uploadPath := upload.GetSavePath()
	dst := path.Join(uploadPath, filename)
	if !upload.CheckExt(fileType, filename) {
		return nil, errors.New("file suffix is not supported.")
	}
	if upload.CheckSavePath(uploadPath) {
		err := upload.CreateSavePath(uploadPath, os.ModePerm)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to create save directory: %v", err))
		}
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}
	if upload.CheckPermission(uploadPath) {
		return nil, errors.New("insufficient file permissions.")
	}
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	accessUrl := path.Join(global.AppSetting.UploadUrl, filename)
	return &FileInfo{Name: filename, AccessUrl: accessUrl}, nil
}
