package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 모든 API 테스트는 서버가 정상적으로 실행 중인 상태에서 실행되어야 합니다.

var serverURL = "http://localhost:8080" // 기본값, .env 값으로 덮어쓰기

func TestMain(m *testing.M) {
	// .env에서 포트 값을 불러옴
	if port := os.Getenv("WIZCRAFT_PORT"); port != "" {
		serverURL = "http://localhost:" + port
	}

	// 서버 시작 대기 (5초 대기)
	time.Sleep(5 * time.Second)

	// 모든 테스트 실행
	exitVal := m.Run()

	// 서버 종료 로직 (필요 시 추가)
	os.Exit(exitVal)
}

func TestCreateUser(t *testing.T) {
	mockUser := map[string]string{
		"email":    "testuser@example.com",
		"password": "securepassword",
	}

	requestBody, _ := json.Marshal(mockUser)
	resp, err := http.Post(serverURL+"/user", "application/json", bytes.NewBuffer(requestBody))

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseData map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&responseData)
	assert.Equal(t, mockUser["email"], responseData["email"])
}

func TestCreateUser_InvalidData(t *testing.T) {
	invalidData := `{ "email": "", "password": "" }`

	resp, err := http.Post(serverURL+"/user", "application/json", bytes.NewBufferString(invalidData))

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetUser(t *testing.T) {
	resp, err := http.Get(serverURL + "/user/123")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdateUser(t *testing.T) {
	updateData := `{ "email": "updateduser@example.com", "password": "newpassword" }`

	client := &http.Client{}
	req, _ := http.NewRequest("PUT", serverURL+"/user/123", bytes.NewBufferString(updateData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDeleteUser(t *testing.T) {
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", serverURL+"/user/123", nil)

	resp, err := client.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
