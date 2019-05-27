package api

import (
	"bytes"
	"log"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
	"errors"
)

// HTTP POST
// inParam postのパラメータ
// outParam response json を構造体に格納するための初期値データ
func Post(url string, inParam interface{}, outParam interface{}) (interface{}, error) {

	// リクエストパラメータをjsonに変換する
	in, err := json.Marshal(inParam)
	// POSTリクエスト
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(in))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// status code が200以外の場合、現状エラー扱いとする
	if resp.StatusCode != http.StatusOK {
		fmt.Println("status code: ", resp.StatusCode)
		return nil, errors.New("status not ok")
	}
	// response body のjsonデータを取得する
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// response body のjsonをデータ型に変換する
	json.Unmarshal(body, &outParam)
	return outParam, nil
}

// func Get(url string, outParam interface{}) (interface{}, error) {
// 	// リクエストパラメータをjsonに変換する
// 	in, err := json.Marshal(inParam)
// 	// POSTリクエスト
// 	resp, err := http.Get(url, "application/json", bytes.NewBuffer(in))
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	// status code が200以外の場合、現状エラー扱いとする
// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Println("status code: ", resp.StatusCode)
// 		return nil, errors.New("status not ok")
// 	}
// 	// response body のjsonデータを取得する
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}

// 	// response body のjsonをデータ型に変換する
// 	json.Unmarshal(body, &outParam)
// 	return outParam, nil
// }