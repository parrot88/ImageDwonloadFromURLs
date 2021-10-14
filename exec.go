//script as download images
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	pathImageUrlFile         = "imgUrl.dat"
	pathDownloadImagesFolder = "./DownloadImages/"
)

func main() {
	urls := getUrls()
	downloadAndSave(urls)
	fmt.Println("end")
}

func downloadAndSave(urls []string) {
	for i, v := range urls {
		v = strings.Replace(v, "\r", "", -1) //remove \r for windows
		urlArr := strings.Split(v, ".")      //get image file extension as .jpg .png
		path := pathDownloadImagesFolder + strconv.Itoa(i) + "." + urlArr[len(urlArr)-1]
		fmt.Println("url: " + v)
		fmt.Println("path: " + path)
		downloadFile(path, v)
	}
}

//get map of urls from dat file
func getUrls() []string {
	var strUrl string
	f, err := os.Open(pathImageUrlFile)
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		strUrl = string(buf[:n])
	}
	return strings.Split(strUrl, "\n")
}

//download file from url into filepath
func downloadFile(filepath string, url string) error {
	fmt.Println("download filepath: " + filepath)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}
