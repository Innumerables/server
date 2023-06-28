package example

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type FileUploadAndDownloadRouter struct{}

func (e *FileUploadAndDownloadRouter) InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	fileUploadAndDownloadRouter := Router.Group("fileUploadAndloadRouter")
	exaFileUploadAndDownloadApi := v1.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi
	{
		fileUploadAndDownloadRouter.POST("upload", exaFileUploadAndDownloadApi.Upload)
		fileUploadAndDownloadRouter.POST("getFileList", exaFileUploadAndDownloadApi.GetFileList)
		fileUploadAndDownloadRouter.POST("deleteFile", exaFileUploadAndDownloadApi.DeleteFile)
		fileUploadAndDownloadRouter.POST("editFileName", exaFileUploadAndDownloadApi.EditFileName)
		fileUploadAndDownloadRouter.POST("breakpointContinue", exaFileUploadAndDownloadApi.BreakpointContinue)
		fileUploadAndDownloadRouter.GET("findFile", exaFileUploadAndDownloadApi.FindFile)
		fileUploadAndDownloadRouter.POST("breakpointContinueFinish", exaFileUploadAndDownloadApi.BreakpointContinueFinish)
		fileUploadAndDownloadRouter.POST("removeChunk", exaFileUploadAndDownloadApi.RemoveChunk)
	}
}
