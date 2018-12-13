package main

import (
	"archive/zip"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

}

func zipFolder(src, dest string) error {
	fileName := filepath.Base(src)
	targetPath := filepath.Join(src, fmt.Sprintf("%s.zip", fileName))
	zipFile, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zw := zip.NewWriter(zipFile)
	defer zw.Close()

	srcFileInfo, err := os.Stat(src)

	if err != nil {
		return err
	}
	var baseDir string
	if srcFileInfo.IsDir() {
		baseDir = fileName
	}

	filepath.Walk(src, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			return err
		}
		header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, src))
		if fileInfo.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := zw.CreateHeader(header)

		return nil
	})

	return nil
}
