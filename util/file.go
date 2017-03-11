package util

import (
	"io"
	"mime/multipart"
	"os"
)

func SaveFile(f *multipart.FileHeader, path string) error {
	fh, err := f.Open()
	defer fh.Close()
	if err != nil {
		return err
	}

	// 打开保存文件句柄
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer fp.Close()
	if err != nil {
		return err
	}
	if _, err = io.Copy(fp, fh); err != nil {
		return err
	}
	return nil
}
