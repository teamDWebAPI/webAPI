package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetListAll struct {
	Message map[string]interface{} `json:"message"`
}

type Dog struct {
	Breed    string      `json:"breed"`
	Subbreed interface{} `json:"sub_breed"`
}

type DogList struct {
	Message []Dog  `json:"message"`
	Status  string `json:"status"`
}

func (dogList DogList) status() {
	fmt.Println(dogList.Message)
	if len(dogList.Message) != 0 {
		dogList.Status = "success"
	} else {
		dogList.Status = "failed"
	}
}

var dogList DogList

func init() {
	var dogListAll GetListAll

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

	if err := json.Unmarshal(body, &dogListAll); err != nil {
		fmt.Println(err)
		return
	}

	for key, value := range dogListAll.Message {
		var dog Dog

		dog.Breed = key
		dog.Subbreed = value

		dogList.Message = append(dogList.Message, dog)
	}
	dogList.status()
}

func getDogHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if len(query) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500") // localhost:5500のオリジンからのアクセスを許可（デモ用）
		json.NewEncoder(w).Encode(dogList)
	} else {
		keyword := query.Get("keyword")
		var response DogList

		for _, dog := range dogList.Message {
			if len(dog.Breed) >= len(keyword) && dog.Breed[:len(keyword)] == keyword {
				response.Message = append(response.Message, dog)
			}
		}
		response.status()
		fmt.Println(len(response.Message))
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500") // localhost:5500のオリジンからのアクセスを許可（デモ用）
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	fmt.Println("Starting the server!")

	// ルートとハンドラ関数を定義
	http.HandleFunc("/api/list", getDogHandler)

	// 8000番ポートでサーバを開始
	http.ListenAndServe(":8000", nil)
}
