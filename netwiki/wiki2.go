package main

import (
	"fmt"
	"log"
	"net/http"
)

// // ハンドラはプログラム中で関数やサブルーチンなどの形で実装され、メモリ上に展開されるが、
// 通常のプログラムの流れには組み込まれず、普段は待機している。そのハンドラが対応すべき処理要求が発生すると
// プログラムの流れを中断してハンドラが呼び出され、要求された処理を実行する。対応付けられた事象の種類により
// // 「イベントハンドラ」「割り込みハンドラ」などの種類がある。

// 待機して、リクエストを待っている
// func handler(w http.ResponseWriter, r *http.Request) {
// 	// http.ResponseWriter値は、HTTPサーバの応答を組み立て、それに書き込むことにより、HTTPクライアントにデータを送信します。
// 	//
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// 	// リクエストのurlのパスから最初（インデックス番号０）からスライスしたものを返す
// }

func viewHandler（w http.ResponseWriter、r * http.Request）{ 
    title：= r.URL.Path [len（ "/ view /"）：] 
    p、_：= loadPage（title）
    fmt.Fprintf（w、 " <h1>％s </ h1> <div>％s </ div> "、p.Title、p.Body）
}

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// 	// fatalでリクエストがあったら値を渡してプロセスが死ぬ
// }

func main（）{ 
    http.HandleFunc（ "/ view /"、viewHandler）
    log.Fatal（http.ListenAndServe（ "：8080"、nil））
}