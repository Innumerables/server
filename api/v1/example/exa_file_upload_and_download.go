package example

import (
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	"server/model/example"
	exampleRes "server/model/example/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileUploadAndDownloadApi struct {
}

// 上传文件
func (fileUp *FileUploadAndDownloadApi) Upload(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接受文件失败", zap.Error(err))
		response.FailWithMessage("接受文件失败", c)
		return
	}
	file, err = fileUploadAndDownloadService.UploadFile(header, noSave)
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(exampleRes.ExaFileResponse{File: file}, "上传成功", c)
}

// 查询文件名,和加载存在的文件
func (fileUp *FileUploadAndDownloadApi) GetFileList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := fileUploadAndDownloadService.GetFileList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// 删除文件
func (fileUp *FileUploadAndDownloadApi) DeleteFile(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = fileUploadAndDownloadService.DeleteFile(file)
	if err != nil {
		global.GVA_LOG.Error("删除失败", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (fileUp *FileUploadAndDownloadApi) EditFileName(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = fileUploadAndDownloadService.EditFileName(file)
	if err != nil {
		global.GVA_LOG.Error("编辑失败", zap.Error(err))
		response.FailWithMessage("编辑失败", c)
		return
	}
	response.OkWithMessage("编辑成功", c)
}
