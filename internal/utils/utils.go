package utils

import "os"

func GetFile(file string) (*os.File, error) {
	path, _ := os.Getwd()
	return os.Open(path + "/assets/" + file)
}
