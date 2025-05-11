package _case

import (
	"fmt"
	"log"
	"net/http"
)

func HttpSimpleCase() {

	http.HandleFunc("/go", myHandle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

// 处理函数
func myHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr)
	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)
	fmt.Println(r.Header)
	fmt.Println(r.Body)

	//	这个常用
	fmt.Println(r.URL.Query().Get("username"))
	w.Write([]byte("hello world"))
}
