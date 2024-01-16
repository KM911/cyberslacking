package util

import "os"

func IsFile(file_ string) bool {
	fileInfo, err := os.Stat(file_)
	if err != nil {
		panic(err)
	}
	return !fileInfo.IsDir()
}

func IsExist(file_ string) bool {
	_, err := os.Stat(file_)
	return err == nil
}

func IsDir(file_ string) bool {
	fileInfo, err := os.Stat(file_)
	if err != nil {
		panic(err)
	}
	return fileInfo.IsDir()
}

func ReadDir(_src string) ([]string, []string) {
	files, err := os.ReadDir(_src)
	if err != nil {
		panic(err)
	}
	var fileNames []string
	var dirNames []string
	for _, file := range files {
		if file.IsDir() {
			dirNames = append(dirNames, file.Name())
		} else {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames, dirNames
}
