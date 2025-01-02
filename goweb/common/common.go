package common

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)



func IsContain(items interface{}, item interface{}) bool {
	switch items.(type) {
	case []int:
		intArr := items.([]int)
		for _, value := range intArr {
			if value == item.(int) {
				return true
			}
		}
	case []string:
		strArr := items.([]string)
		for _, value := range strArr {
			if value == item.(string) {
				return true
			}
		}
	default:
		return false
	}
	return false
}

// remove removes a string from a slice of strings
func RemoveStr(slice []string, s string) []string {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// 判断传入的文件夹路径是否存在
func IsDirExist(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return s.IsDir()
}

// 将传入的图片base64字符串转换为图片保存到指定路径、
func SaveBase64Img(base64Str string, path string) (string, error) {
	b, err := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, base64Str)
	if err != nil {
		return "", err
	}
	if !b {
		return "", fmt.Errorf("invalid base64 image string")
	}

	re, err := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
	if err != nil {
		return "", err
	}
	allData := re.FindAllSubmatch([]byte(base64Str), 2)
	if len(allData) == 0 || len(allData[0]) < 2 {
		return "", fmt.Errorf("invalid base64 image string")
	}
	fileType := string(allData[0][1]) //png, jpeg 后缀获取

	base64Str2 := re.ReplaceAllString(base64Str, "")

	date := time.Now().Format("2006-01-02")
	if ok := IsDirExist(path + "/" + date); !ok {
		err := os.Mkdir(path+"/"+date, 0666)
		if err != nil {
			return "", err
		}
	}

	curFileStr := strconv.FormatInt(time.Now().UnixNano(), 10)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(99999)

	var file string = path + "/" + date + "/" + curFileStr + strconv.Itoa(n) + "." + fileType
	byte, err := base64.StdEncoding.DecodeString(base64Str2)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(file, byte, 0666)
	if err != nil {
		return "", err
	}
	return file, nil
}
