package bigxxby_test

import (
	bigxxby "bigxxby/funcs"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGetRequestsOnly(t *testing.T) {
	// Создаем новый запрос GET
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создаем ResponseRecorder (реализует http.ResponseWriter) для записи ответа
	rr := httptest.NewRecorder()

	// Запускаем обработчик HTTP (handler) с помощью ServeHTTP и записываем ответ
	handler := http.HandlerFunc(bigxxby.MainHandler) // Замените YourHandlerFunction на ваш обработчик
	handler.ServeHTTP(rr, req)

	// Проверяем статус код
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Проверяем, что запрос был GET
	if req.Method != "GET" {
		t.Errorf("Expected GET request, got %s", req.Method)
	}
}

func Teest(t *testing.T) {
	t.Error("FAIL")
}

func TestGetRequestsOnlyArtists(t *testing.T) {
	for i := 1; i <= 52; i++ {
		log.Println("check")
		str := strconv.Itoa(i)
		// Создаем новый запрос GET
		req, err := http.NewRequest("GET", "/artsits/2"+str, nil)
		if err != nil {
			t.Fatal(err)
		}

		// Создаем ResponseRecorder (реализует http.ResponseWriter) для записи ответа
		rr := httptest.NewRecorder()

		// Запускаем обработчик HTTP (handler) с помощью ServeHTTP и записываем ответ
		handler := http.HandlerFunc(bigxxby.ArtistIdHandler) // Замените YourHandlerFunction на ваш обработчик
		handler.ServeHTTP(rr, req)

		// Проверяем статус код
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Проверяем, что запрос был GET
		if req.Method != "GET" {
			t.Errorf("Expected GET request, got %s", req.Method)
		}
	}
}
