package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	log "github.com/inconshreveable/log15"
)

var (
	dirPath         = "/home/deepir/桌面/"
	ErrDirNotFound  = errors.New("dir ads not found")
	ErrFileNotFound = errors.New("file ads not found")
	ErrCreateDir    = errors.New("create dir err")
	ErrCreateFile   = errors.New("create file err")
	ErrReadFile     = errors.New("read file err")
	ErrWriteFile    = errors.New("write file err")
)

func main() {

	filebyte, err := ReadLocalFile(dirPath + "douyu.lst")
	if err != nil {
		log.Error(fmt.Sprintf("read file err=%s", err))
	}

	url_lines := byteToList(filebyte)

	var filename string
	var file *os.File
	for i := 0; i < len(url_lines); i++ {
		err := os.MkdirAll(dirPath+"url_file/"+strconv.Itoa(i/1000), 0755)
		if i%100 == 0 {
			filename = strconv.Itoa(int(i / 100))

			file, err = os.Create(dirPath + "url_file/" + strconv.Itoa(i/1000) + "/" + filename)
			if err != nil {
				log.Error(fmt.Sprintf("create file err%v", err))
			}
		}

		_, err = file.WriteString(url_lines[i] + "\n")
		if err != nil {
			log.Error(fmt.Sprintf("write file err%v", err))
		}

	}

}

func ReadLocalFile(filePath string) ([]byte, error) {

	isExist, _ := PathExists(filePath)

	if !isExist {
		log.Error(fmt.Sprintf(" file not Exists err"))
		return nil, ErrFileNotFound
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error(fmt.Sprintf("read file err%v", err))

		return nil, ErrFileNotFound
	}

	return data, nil
}

func byteToList(filebyte []byte) []string {

	urlStr := string(filebyte)
	url_lines := strings.Split(urlStr, "\n")

	return url_lines
}

func SaveFile(filename string, imageUrl []string) error {
	//	isExist, _ := PathExists(dirPath)

	//	if !isExist {
	//		err := os.MkdirAll(dirPath, 0755)
	//		if err != nil {
	//			log.Error(fmt.Sprintf("create dir err%v", err))
	//			return ErrCreateDir
	//		}
	//	}

	file, err := os.Create(dirPath + filename)
	if err != nil {
		log.Error(fmt.Sprintf("create file err%v", err))

		return ErrCreateFile
	}

	_, err = file.WriteString(toString(imageUrl))
	if err != nil {
		log.Error(fmt.Sprintf("write file err%v", err))
		os.Remove(dirPath + filename)
		return ErrWriteFile
	}

	return nil
}

func toString(s []string) string {

	var buffer bytes.Buffer
	for i := 0; i < len(s); i++ {

		if i == len(s)-1 {
			buffer.WriteString(s[i])
		} else {
			buffer.WriteString(s[i] + "\n")
		}

	}

	return buffer.String()
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
