package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type Post struct {
    Message map[string]interface{} `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    var post Post

    resp, err := http.Get("https://dog.ceo/api/breeds/list/all")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        fmt.Println("Error: status code", resp.StatusCode)
        return
    }

    body, _ := io.ReadAll(resp.Body)

    if err := json.Unmarshal(body, &post); err != nil {
        fmt.Println(err)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    // 構造体をJSONにエンコードしてレスポンスとして送信
    json.NewEncoder(w).Encode(post)
}

func main() {
    fmt.Println("Starting the server!")

    // ルートとハンドラ関数を定義
    http.HandleFunc("/api/hello", helloHandler)

    // 8000番ポートでサーバを開始
    http.ListenAndServe(":8000", nil)
}
