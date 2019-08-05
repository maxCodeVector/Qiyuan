package hander

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"qiyuan/service"
)

func HandleUpload(c *gin.Context) {

	orderID := c.Param("order_id")
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get file err : %s", err.Error()))
		return

	}
	//获取文件名
	filename := header.Filename
	filePath := service.UploadFile(orderID, file, filename)
	//以json格式返回文件存放路径
	c.JSON(http.StatusOK, gin.H{"filepath": filePath})

}


func HandleDownload(c *gin.Context) {
	fileName := c.Query("fileUrl")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("../"+fileName)
}


func HandleExport(c *gin.Context){
	outFile := "../file/order.xlsx"
	service.ExportData(outFile)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", outFile))
	//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(outFile)
}