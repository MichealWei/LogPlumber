package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func createFolderAndFile(baseFolder string, subFolder string, fileExt string) (*os.File, error) {
	err := os.MkdirAll(baseFolder, os.ModePerm)
	if err != nil {
		return nil, err
	}
	FileName := currentDir + subFolder + strconv.FormatInt(time.Now().Unix(), 10) + fileExt

	File, err := os.OpenFile(FileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}

	return File, nil
}

func getDirectorySize(directory string) (int, error) {
	totalSize := 0
	err := filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				totalSize = +int(info.Size())

			}
			return nil
		})
	return totalSize, err
}

func createFolder(folderName string) error {
	var err error
	if IsFolderExist(folderName) {
		return nil
	}
	err = os.Mkdir(folderName, 0777)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Chmod(folderName, 0777)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func IsFolderExist(path string) bool {

	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}
