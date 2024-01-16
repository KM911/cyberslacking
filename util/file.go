package util

import "os"

func ReadAll(_src string) string {
	_data, err := os.ReadFile(_src)
	if err != nil {
		panic(err)
	}
	return string(_data)
}

func CreatFile(_src string, _data string) {
	err := os.WriteFile(_src, []byte(_data), 0666)
	if err != nil {
		panic(err)
	}
}

func DeleteFile(_src string) {
	err := os.Remove(_src)
	if err != nil {
		panic(err)
	}
}

func CopyFile(_src string, dst string) {
	_data := ReadAll(_src)
	CreatFile(dst, _data)
}
