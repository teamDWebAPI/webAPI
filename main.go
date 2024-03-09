package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type GetListAll struct {
	Message map[string][]string `json:"message"`
	Status  string              `json:"status"`
}

type Dog struct {
	ID       int      `json:"id"`
	Breed    string   `json:"breed"`
	Subbreed []string `json:"sub_breed"`
}

type ResponseDetail struct {
	Message Dog    `json:"message"`
	Status  string `json:"status"`
}

func (detail *ResponseDetail) status() {
	if detail.Message.ID != 0 {
		detail.Status = "success"
	} else {
		detail.Message.Subbreed = []string{}
		detail.Status = "failed"
	}
}

type BreedNameList struct {
	Message []string `json:"message"`
	Status  string   `json:"status"`
}

func (dogList *BreedNameList) status() {
	if len(dogList.Message) != 0 {
		dogList.Status = "success"
	} else {
		dogList.Message = []string{}
		dogList.Status = "failed"
	}
}

type getUrl struct {
	Message []string `json:"message"`
	Status  string   `json:"status"`
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

var breedNameList BreedNameList
var breedDetailList []Dog

// main()の前に呼ばれる特殊関数
func init() {
	var dogListAll GetListAll
	getResponseFromDogApi(&dogListAll, "https://dog.ceo/api/breeds/list/all")

	var id int = 1
	for key, value := range dogListAll.Message {
		var dog Dog

		dog.ID = id
		dog.Breed = key
		dog.Subbreed = value

		breedNameList.Message = append(breedNameList.Message, dog.Breed)
		breedDetailList = append(breedDetailList, dog)
		id++
	}
	breedNameList.status()
}

func showListHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// クエリが指定されていないときlistALLを返す
	if len(query) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500") // localhost:5500のオリジンからのアクセスを許可（デモ用）
		json.NewEncoder(w).Encode(breedNameList)
	} else {
		// filterで指定された文字から始まるbreedのlistを返す
		filter := query.Get("filter")
		// クエリsort=ascend（昇順） or descend（降順）で返すリストのソートを行う
		sor := query.Get("sort")

		var response BreedNameList
		if len(filter) > 0 {
			for _, dog := range breedNameList.Message {
				if len(dog) >= len(filter) && dog[:len(filter)] == filter {
					response.Message = append(response.Message, dog)
				}
			}
		} else {
			response = breedNameList
		}

		if sor == "ascend" {
			sort.Strings(response.Message)
		} else if sor == "descend" {
			sort.Slice(response.Message, func(i, j int) bool {
				return response.Message[i] > response.Message[j]
			})
		}

		response.status()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500") // localhost:5500のオリジンからのアクセスを許可（デモ用）
		json.NewEncoder(w).Encode(response)
	}
}

// itemの詳細を返すハンドラー
func detailHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])

	var responseDetail ResponseDetail
	if err != nil {
		// idではなくbreedで詳細を取得
		for _, doc := range breedDetailList {
			if doc.Breed == parts[len(parts)-1] {
				responseDetail.Message = doc
			}
		}
	} else {
		// idで詳細を取得
		for _, doc := range breedDetailList {
			if doc.ID == id {
				responseDetail.Message = doc
			}
		}
	}
	responseDetail.status()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500") // localhost:5500のオリジンからのアクセスを許可（デモ用）
	json.NewEncoder(w).Encode(responseDetail)
}

// クエリの有無によってをDogapiのエンドポイントを指定し返す関数
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
	fmt.Println(endpoint)
	return endpoint
}

// dogapiから写真のURLを返すハンドラー
func getUrlHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	breedName := query.Get("breed")
	subBreedName := query.Get("sub-breed")
	count := query.Get("count")

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
	http.HandleFunc("/api/list", showListHandler) // eg. /api/list?filter=a&sort=ascend
	http.HandleFunc("/api/item/", detailHandler)  // eg. /api/item/{id} or /api/item/{breed}
	http.HandleFunc("/api/images", getUrlHandler) // eg. /api/images?breed=hound&sub-breed=afghan?count=3

	// 8000番ポートでサーバを開始
	http.ListenAndServe(":8000", nil)
}
