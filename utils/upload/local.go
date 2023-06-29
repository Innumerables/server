package upload

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"server/global"
	"server/utils"

	"go.uber.org/zap"
)

type Local struct {
}

// 上传文件到本地，通过拷贝赋值方式
func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	//读取文件后缀
	ext := path.Ext(file.Filename)
	//读取文件名并进行加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	//拼接新的文件名
	filename := name + "_" + time.Now().Format("20060606150405") + ext
	//尝试创建路径
	mkdirErr := os.Mkdir(global.GVA_CONFIG.Local.StorePath, os.ModePerm)
	if mkdirErr != nil {
		global.GVA_LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	//拼接路径和文件名
	p := global.GVA_CONFIG.Local.StorePath + "/" + filename
	filepath := global.GVA_CONFIG.Local.Path + "/" + filename

	f, openError := file.Open()
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close()

	out, createErr := os.Create(p)
	if createErr != nil {
		global.GVA_LOG.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))

		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.GVA_LOG.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return filepath, filename, nil
}

func (*Local) DeleteFile(key string) error {
	p := global.GVA_CONFIG.Local.StorePath + "/" + key
	if strings.Contains(p, global.GVA_CONFIG.Local.StorePath) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败+err:" + err.Error())
		}
	}
	return nil
}