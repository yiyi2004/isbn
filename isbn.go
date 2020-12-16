package isbn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Book struct {
	CreateTime string   `json:"create_time"`
	Isbn       string   `json:"isbn"`
	Title      string   `json:"title"`
	BookInfo   BookInfo `json:"book_info"`
}

type BookInfo struct {
	Auther          string `json:"作者"`
	Translator      string `json:"译者"`
	PublishingHouse string `json:"出版社"`
	PublishDate     string `json:"出版年"`
	Pages           string `json:"页数"`
	Price           string `json:"定价"`
	BookBinding     string `json:"装帧"`
	Series          string `json:"丛书"`
	ISBN            string `json:"ISBN"`
}

// GetBookInfoByISBN -
func GetBookInfoByISBN(isbn string) (*Book, error) {
	client := &http.Client{}
	book := &Book{}

	url := "http://book.feelyou.top/isbn/" + isbn
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, book); err != nil {
		return nil, err
	}

	return book, nil
}
