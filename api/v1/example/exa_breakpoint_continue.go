package example

import (
	"fmt"
	"io"
	"mime/multipart"
	"server/global"
	"server/model/common/response"
	"server/model/example"
	exampleRes "server/model/example/response"
	"strconv"

	"server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (e *FileUploadAndDownloadApi) FindFile(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	chuankTotal, _ := strconv.Atoi(c.Query("chunkTotal"))
	file, err := fileUploadAndDownloadService.FindOrCreateFile(fileMd5, fileName, chuankTotal)
	if err != nil {
		global.GVA_LOG.Error("查找失败", zap.Error(err))
		response.FailWithMessage("查找失败", c)
	}
	response.OkWithDetailed(exampleRes.FileResponse{File: file}, "查找成功", c)

}

func (e *FileUploadAndDownloadApi) BreakpointContinue(c *gin.Context) {
	fileMd5 := c.Request.FormValue("fileMd5")
	fileName := c.Request.FormValue("fileName")
	chunkMd5 := c.Request.FormValue("chunkMd5")
	chunkNumber, _ := strconv.Atoi(c.Request.FormValue("chunkNumber"))
	chunkTotal, _ := strconv.Atoi(c.Request.FormValue("chunkTotal"))
	_, FileHeader, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	f, err := FileHeader.Open()
	if err != nil {
		global.GVA_LOG.Error("文件读取失败!", zap.Error(err))
		response.FailWithMessage("文件读取失败", c)
		return
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	cen, _ := io.ReadAll(f)
	if !utils.CheckMd5(cen, chunkMd5) {
		global.GVA_LOG.Error("检查md5失败!", zap.Error(err))
		response.FailWithMessage("检查md5失败", c)
		return
	}
	file, err := fileUploadAndDownloadService.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		global.GVA_LOG.Error("查找或创建记录失败!", zap.Error(err))
		response.FailWithMessage("查找或创建记录失败", c)
		return
	}
	pathC, err := utils.BreakPointContinue(cen, fileName, chunkNumber, chunkTotal, fileMd5)
	if err != nil {
		global.GVA_LOG.Error("断点续传失败!", zap.Error(err))
		response.FailWithMessage("断点续传失败", c)
		return
	}

	if err = fileUploadAndDownloadService.CreateFileChunk(file.ID, pathC, chunkNumber); err != nil {
		global.GVA_LOG.Error("创建文件记录失败!", zap.Error(err))
		response.FailWithMessage("创建文件记录失败", c)
		return
	}
	response.OkWithMessage("切片创建成功", c)
}

func (e *FileUploadAndDownloadApi) BreakpointContinueFinish(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	filePath, err := utils.MakeFile(fileName, fileMd5)
	if err != nil {
		global.GVA_LOG.Error("文件创建失败", zap.Error(err))
		response.FailWithDetailed(exampleRes.FilePathResponse{FilePath: filePath}, "文件创建失败", c)
	}
	response.OkWithDetailed(exampleRes.FilePathResponse{FilePath: filePath}, "文件创建成功", c)
}

func (e *FileUploadAndDownloadApi) RemoveChunk(c *gin.Context) {
	var file example.ExaFile
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.RemoveChunk(file.FileMd5)
	if err != nil {
		global.GVA_LOG.Error("缓存切片删除失败", zap.Error(err))
		return
	}
	err = fileUploadAndDownloadService.RemoveChunk(file.FileMd5, file.FilePath)
	if err != nil {
		global.GVA_LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("缓存切片删除成功", c)
}
