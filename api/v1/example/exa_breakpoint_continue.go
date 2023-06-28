package example

import "github.com/gin-gonic/gin"

func (e *FileUploadAndDownloadApi) BreakpointContinue(c *gin.Context) {
	fileMd5 := c.Request.FormValue("fileMd5")
	fileName := c.Request.FormValue("fileName")
	chunkMd5 := c.Request.FormValue("chuankMd5")

}
