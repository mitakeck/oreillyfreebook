package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Downloader : downloader
type Downloader struct {
}

var (
	baseURL = "http://www.oreilly.com/%s/free/"

	// const で宣言出来ないんご
	categorys = [...]string{"data", "business", "design", "iot", "programming", "security", "web-platform", "webops"}

	// worker size : 並列ダウンロード数
	worker = 4
	// delay time : サーバに負荷を与えないように一定の時間待つ
	waitSec = 5
)

func (d *Downloader) createURI(category string) ([]string, error) {
	if category == "all" {
		us := make([]string, len(categorys))
		for i, cat := range categorys {
			us[i] = fmt.Sprintf(baseURL, cat)
		}
		return us, nil
	}
	for _, cat := range categorys {
		if cat == category {
			return []string{fmt.Sprintf(baseURL, category)}, nil
		}
	}
	return nil, fmt.Errorf("Invalid category : %s", category)
}

func (d *Downloader) download(wg *sync.WaitGroup, q chan string, directory string) {
	defer wg.Done()
	client := &http.Client{}

	for {
		url, ok := <-q
		if !ok {
			return
		}

		fileName, _ := pop(strings.Split(url, "/"))
		fileName = filepath.Join(directory, fileName)
		fmt.Printf("Download : %s\n", fileName)
		output, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Error while creating %s : %v\n", fileName, err)
			continue
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("Unable to create request (%s) : %v\n", url, err)
			continue
		}

		res, resErr := client.Do(req)
		if resErr != nil {
			fmt.Printf("Error while downloading %s : %v\n", url, err)
			continue
		}
		if res.StatusCode == 404 {
			fmt.Printf("Error while downloading %s : HTTP Status 404 \n", url)
			RmErr := os.Remove(fileName)
			if RmErr != nil {
				fmt.Printf("Error while remove %s : %v", fileName, err)
			}
			continue
		}

		_, err = io.Copy(output, res.Body)
		if err != nil {
			fmt.Printf("Error while downloading %s : %v\n", url, err)
			continue
		}

		// clean up
		output.Close()
		res.Body.Close()
		// delay
		time.Sleep(time.Duration(waitSec) * time.Second)
	}
}

// Download : download file
func (d *Downloader) Download(category string, format string, directory string, searchWord string) error {
	list, err := d.getFileList(category, format, searchWord)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	q := make(chan string, len(list))

	// make worker
	for i := 0; i < worker; i++ {
		wg.Add(1)
		go d.download(&wg, q, directory)
	}

	for _, url := range list {
		q <- url
	}
	close(q)

	wg.Wait()
	return nil
}

func isContain(s string, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func (d *Downloader) getFileList(category string, format string, searchWord string) ([]string, error) {
	us, err := d.createURI(category)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	for _, url := range us {
		doc, err := goquery.NewDocument(url)
		if err != nil {
			return nil, err
		}

		doc.Find(".cover-showcase a").Each(func(i int, s *goquery.Selection) {
			u, _ := s.Attr("href")
			u = strings.Replace(u, "free/", "free/files/", 1)
			uSplit := strings.Split(u, ".csp")
			if len(uSplit) != 2 {
				return
			}
			u = uSplit[0] + "." + format
			if isContain(u, searchWord) {
				result = append(result, u)
			}
		})
	}
	return result, nil
}
