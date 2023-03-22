package xsdProcessor

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type SchemaFetcher interface {
	Fetch(path string) ([]byte, error)
}

type XMLFetcher struct {
}

func (f XMLFetcher) isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (f XMLFetcher) fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (f XMLFetcher) Fetch(location string) ([]byte, error) {
	if f.fileExists(location) {
		return FileSystemFetcher{}.Fetch(location)
	}

	if f.isUrl(location) {
		return HttpFetcher{}.Fetch(location)
	}

	return []byte{}, fmt.Errorf("can't fetch XML from: %s", location)
}

type FileSystemFetcher struct {
}

func (f FileSystemFetcher) Fetch(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

type HttpFetcher struct {
}

func (f HttpFetcher) Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return []byte{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
