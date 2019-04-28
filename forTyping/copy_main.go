package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`
		<html>
			<head>
				<title>チャット</title>
			</head>
			<body>
				チャットしよう。
			</body>
		</html>
		`)); err!=nil{
			log.Fatal("w.Write:",err)
		}
	})

	//Webサーバーを開始します。
	if err:=http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatal("ListenAndServe", err)
	}
}