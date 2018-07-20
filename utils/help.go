package utils

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
	//"swingbaby-go/config"
)

func MobiVerify(mobile string) bool {
	if m, _ := regexp.MatchString(`^(1[3|4|5|7|8][0-9]\d{4,8})$`, mobile); !m {
		return false
	}
	return true
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func Upload(filepath string, filename string, r *http.Request) (string, error) {
	if !checkFileIsExist(filepath) {
		err := os.Mkdir(filepath, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
			return "", err
		}
	}

	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile(filename)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer file.Close()

	imgname := GetRandomString(16)
	imgstr := fmt.Sprintf("%s.jpg", imgname)

	for checkFileIsExist(filepath + imgstr) {
		imgname = GetRandomString(16)
		imgstr = fmt.Sprintf("%s.jpg", imgname)
	}

	f, err := os.OpenFile(filepath+imgstr, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer f.Close()
	io.Copy(f, file)

	return strings.TrimLeft(filepath+imgstr, "."), nil
}
