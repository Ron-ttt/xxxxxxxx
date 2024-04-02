package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_handlerWrapper_IndexPage(t *testing.T) { // работает удивительно ведь не я это делала
	type want struct { // я не ебу что для негативного надо работает и заебись
		code int
		//response    string
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test #1",
			want: want{
				code: 201,
				//response:    "https://example.com",
				contentType: "text/plain",
			},
		},
		{
			name: "negative test #1",
			want: want{
				code: 400,
				//response:    "https://example.com",
				contentType: "text/plain",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hw := Init()
			reqBody := strings.NewReader("https://example.com")
			request := httptest.NewRequest(http.MethodPost, "/", reqBody)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			hw.IndexPage(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.NotEmpty(t, string(resBody))
			assert.Equal(t, test.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func Test_handlerWrapper_Redirect(t *testing.T) {
	type want struct {
		code        int
		location    string
		contentType string
	}

	tests := []struct {
		name string
		id   string
		want want
	}{
		{
			name: "positive test #1",
			id:   "123456",
			want: want{
				code:        http.StatusTemporaryRedirect,
				location:    "http://love_nika",
				contentType: "",
			},
		},
		{
			name: "negative test #1",
			id:   "invalid",
			want: want{
				code:        http.StatusBadRequest,
				location:    "",
				contentType: "",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Создаем тестовый обработчик
			handler := MInit()

			// Создаем тестовый запрос

			r := mux.NewRouter()
			r.HandleFunc("/{id}", handler.Redirect)
			// Выполняем запрос
			w2 := httptest.NewRecorder()
			r.ServeHTTP(w2, httptest.NewRequest(http.MethodGet, handler.baseURL+test.id, nil))

			// Проверяем код ответа
			assert.Equal(t, test.want.code, w2.Result().StatusCode)

			// Проверяем заголовок Location
			//location := rr.Header().Get("Location")
			//assert.Equal(t, test.want.location, location)

			// Проверяем тип контента
			//assert.Equal(t, test.want.contentType, rr.Header().Get("Content-Type"))

			// Проверяем, что сгенерированная строка добавлена в хранилище

		})
	}
}

// func Test_handlerWrapper_Redirect(t *testing.T) {
// 	type want struct {
// 		code int
// 		//response string
// 	}
// 	tests := []struct {
// 		name string
// 		want want
// 	}{
// 		{
// 			name: "positive test #1",
// 			want: want{
// 				code: 307,
// 			},
// 		},
// 		{
// 			name: "negative test #1",
// 			want: want{
// 				code: 400,
// 			},
// 		},
// 	}
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			// здесь будет запрос и проверка ответа
// 			hw := Init()
// 			request := httptest.NewRequest(http.MethodGet, "/status", nil)
// 			// создаём новый Recorder
// 			w := httptest.NewRecorder()
// 			hw.Redirect(w, request)

// 			res := w.Result()
// 			// проверяем код ответа
// 			assert.Equal(t, test.want.code, res.StatusCode)
// 			// получаем и проверяем тело запроса
// 			//defer res.Body.Close()
// 			//resBody, err := io.ReadAll(res.Body)

// 			//require.NoError(t, err)
// 			//assert.Equal(t, test.want.response, string(resBody))

// 		})
// 	}
// }
