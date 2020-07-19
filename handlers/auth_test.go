package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleAuth(t *testing.T) {
	mux := http.NewServeMux()            // テストを実行するマルチプレクサを生成
	mux.HandleFunc("/user/", HandleAuth) // テスト対象のハンドラを付加

	jsonStr := strings.NewReader(`{"email":"test@test.com", "password":"test"}`)

	writer := httptest.NewRecorder()                         // 返されたhttp レスポンスを取得
	request, _ := http.NewRequest("POST", "/user/", jsonStr) // テストしたいハンドラ宛のリクエストを作成
	mux.ServeHTTP(writer, request)                           // テスト対象のハンドラにリクエストを送信

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	var user User
	json.Unmarshal(writer.Body.Bytes(), &user)
	if user.Email != "test@test.com" {
		t.Errorf("Email is %v", user.Email)
	}
	if user.Password != "test" {
		t.Errorf("Password is %v", user.Password)
	}
}