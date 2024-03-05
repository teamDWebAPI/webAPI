package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetListAll struct {
	Message map[string]interface{} `json:"message"`
	Status  string                 `json:"status"`
}

type Dog struct {
	Breed    string      `json:"breed"`
	Subbreed interface{} `json:"sub_breed"`
}

type DogList struct {
	Message []Dog  `json:"message"`
	Status  string `json:"status"`
}

func (dogList *DogList) status() {
	if len(dogList.Message) != 0 {
		dogList.Status = "success"
	} else {
		dogList.Message = []Dog{}
		dogList.Status = "failed"
	}
}

// dog apiから呼び出したjsonを指定の構造体に格納する関数
func getResponseFromDogApi(st interface{}, endpoint string) {
	resp, err := http.Get(endpoint)
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

	if err := json.Unmarshal(body, &st); err != nil {
		fmt.Println(err)
		return
	}

}

var dogList DogList

// main()の前に呼ばれる特殊関数
func init() {
	var dogListAll GetListAll
	getResponseFromDogApi(&dogListAll, "https://dog.ceo/api/breeds/list/all")

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

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500") // localhost:5500のオリジンからのアクセスを許可（デモ用）
		json.NewEncoder(w).Encode(response)
	}
}

type getUrl struct {
	Message []string `json:"message"`
	Status  string   `json:"status"`
}

func getEndpoint(breedName string, subBreedName string, count string) string {
	endpoint := "https://dog.ceo/api/"

	if count == "" {
		count = "1"
	}

	if breedName != "" && subBreedName != "" {
		endpoint += "breed/" + breedName + "/" + subBreedName + "/images/random/" + count
	} else if breedName != "" && subBreedName == "" {
		endpoint += "breed/" + breedName + "/images/random/" + count
	} else if breedName == "" && subBreedName == "" {
		endpoint += "breeds/image/random/" + count
	}

	return endpoint
}

func getUrlHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	breedName := query.Get("breed")
	subBreedName := query.Get("sub-breed")
	count := query.Get("c")

	endpoint := getEndpoint(breedName, subBreedName, count)

	var urls getUrl

	getResponseFromDogApi(&urls, endpoint)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500") // localhost:5500のオリジンからのアクセスを許可（デモ用）
	json.NewEncoder(w).Encode(urls)
}

func main() {
	fmt.Println("Starting the server!")

	// ルートとハンドラ関数を定義
	http.HandleFunc("/api/list", getDogHandler)
	http.HandleFunc("/api/images", getUrlHandler)

	// 8000番ポートでサーバを開始
	http.ListenAndServe(":8000", nil)
}
