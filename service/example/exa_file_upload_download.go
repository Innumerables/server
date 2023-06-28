package example

import (
	"errors"
	"mime/multipart"
	"server/global"
	"server/model/common/request"
	"server/model/example"
	"strings"

	"server/utils/upload"
)

type FileUploadAndDownloadService struct{}

func (e *FileUploadAndDownloadService) UploadFile(header *multipart.FileHeader, noSave string) (file example.ExaFileUploadAndDownload, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := example.ExaFileUploadAndDownload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return f, e.Upload(f) //将存放位置保存到数据库中
	}
	return
}

func (e *FileUploadAndDownloadService) Upload(file example.ExaFileUploadAndDownload) error {
	return global.GVA_DB.Create(&file).Error
}

func (e *FileUploadAndDownloadService) GetFileList(pageInfo request.PageInfo) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	keyword := pageInfo.Keyword
	db := global.GVA_DB.Model(&example.ExaFileUploadAndDownload{})
	var fileList []example.ExaFileUploadAndDownload
	if len(keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+keyword+"%") //用于输入查询条件，条件为文件名
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileList).Error
	return
}

func (e *FileUploadAndDownloadService) FindFile(id uint) (example.ExaFileUploadAndDownload, error) {
	var file example.ExaFileUploadAndDownload
	err := global.GVA_DB.Where("id = ?", id).First(&file).Error
	return file, err
}

// 删除上传文件，同时在本地的位置也删，以及数据库存放的记录
func (e *FileUploadAndDownloadService) DeleteFile(file example.ExaFileUploadAndDownload) (err error) {
	var fileFromDb example.ExaFileUploadAndDownload
	fileFromDb, err = e.FindFile(file.ID) //使用了MD5加密，需要得到加密后端关键值，因为在存放文件名称时使用了加密
	if err != nil {
		return
	}
	oss := upload.NewOss()
	err = oss.DeleteFile(fileFromDb.Key) //根据关键值删除指定的文件
	if err != nil {
		return errors.New("文件删除失败")
	}
	err = global.GVA_DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return
}

func (e *FileUploadAndDownloadService) EditFileName(file example.ExaFileUploadAndDownload) (err error) {
	var fileFromDb example.ExaFileUploadAndDownload
	err = global.GVA_DB.Where("id = ?", file.ID).First(&fileFromDb).Update("name", file.Name).Error
	return
}
