package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	err := tarFolder("/Users/herrdu/tmp/15", "/Users/herrdu/tmp")
	// err := unTarFile("/Users/herrdu/tmp/15.tar", "/Users/herrdu/download")
	if err != nil {
		log.Println(err)
	}
}

func tarFolder(src, dest string) error {
	fileName := filepath.Base(src)
	target := filepath.Join(dest, fmt.Sprintf("%s.tar", fileName))

	tarFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	tw := tar.NewWriter(tarFile)
	defer tw.Close()

	fileInfo, err := os.Stat(src)
	if err != nil {
		return nil
	}

	var baseDir string
	if fileInfo.IsDir() {
		baseDir = filepath.Base(src)
	}

	return filepath.Walk(src, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := tar.FileInfoHeader(fileInfo, fileInfo.Name())
		if err != nil {
			return err
		}
		header.Name = filepath.Join(baseDir, strings.TrimPrefix(filePath, src))

		fmt.Println("baseDir", header.Name)

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(tw, file)
			return err
		} else {
			fmt.Println(filePath)
		}
		return nil
	})
}

func unTarFile(src, dest string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()
	tarReader := tar.NewReader(file)

	for {
		hdr, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(dest, hdr.Name)
		fileInfo := hdr.FileInfo()

		if fileInfo.IsDir() {
			if err = os.MkdirAll(path, fileInfo.Mode()); err != nil {
				return err
			}
			continue
		}
		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fileInfo.Mode())
		if err != nil {
			return err
		}
		if _, err := io.Copy(file, tarReader); err != nil {
			log.Println(err)
		}
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
