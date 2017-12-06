package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"strconv"
)

type LongUrl struct {
	Url  string `json:"url"`
}

type ShortUrl struct {
	Key  string `json:"key"`
}

var(
	keyId = 0
	URLStore = make(map[int]string)
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", AddURL).Methods("POST")
	r.HandleFunc("/{key}", GetURL)/*.Methods("GET")*/ //�������� ������ �������� ����������
	http.ListenAndServe(":8082", r)
}

func AddURL(w http.ResponseWriter, r *http.Request) {
	//r - ���������� ������
	//w - �����
	var MyLongUrl LongUrl
	json.NewDecoder(r.Body).Decode(&MyLongUrl) // �������� � itemAdd ���������� �� json-������

	var MyShortUrl ShortUrl
	MyShortUrl.Key = strconv.Itoa(keyId)
	URLStore[keyId] = MyLongUrl.Url //�������� URL � ����
	keyId += 1

	json.NewEncoder(w).Encode(&MyShortUrl)
}

func GetURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	InputString := vars["key"]

	Input, err := strconv.Atoi(InputString)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Location", URLStore[Input])
	w.WriteHeader(http.StatusMovedPermanently)
	var MyLongUrl LongUrl
	MyLongUrl.Url = URLStore[Input]
	/*j, err := json.Marshal(MyLongUrl)
	if err != nil {
		panic(err)
	}
	w.Write(j)*/
}