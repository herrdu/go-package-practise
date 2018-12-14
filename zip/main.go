package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// zipFolder("/Users/herrdu/tmp/15", "/Users/herrdu/tmp")
	unZip("/Users/herrdu/tmp/15.zip", "/Users/herrdu/download/15zip")
}

func zipFolder(src, dest string) error {
	fileName := filepath.Base(src)
	targetPath := filepath.Join(dest, fmt.Sprintf("%s.zip", fileName))
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

		if fileInfo.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)

		return nil
	})

	return nil
}

func unZip(src, dest string) error {
	zr, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer zr.Close()

	if err := os.MkdirAll(dest, 0755); err != nil {
		return err
	}

	for _, file := range zr.File {
		path := filepath.Join(dest, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}
		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}
	return nil
}
