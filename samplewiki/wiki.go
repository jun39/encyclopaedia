package main

import (
	"fmt"
	"io/ioutil"
)

// 各ページのデータを保存するページ型を定義
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	// 一括して書き込むioutil.WireFileメソッド
	// 最後の引数はファイルパーミッション
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// ページのタイトルを読み込む関数
// 最後は返り値の型が*oageとerrorの２つなのでかっこで囲んでいる
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"

	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
		// go言語では探したものがなかった場合はnul値だけでなく、なかったファイル名（err）を返り値で返す
	}
	return &Page{Title: title, Body: body}, nil
}
func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
