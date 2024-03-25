package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_handlerWrapper_IndexPage(t *testing.T) { // работает удивительно ведь не я это делала
	type want struct {
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
func TestRedirect(t *testing.T) { // работает на пол шишечки типо негативный
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
		// {
		// 	name: "positive test #1",
		// 	id:   "123456",
		// 	want: want{
		// 		code:        http.StatusTemporaryRedirect,
		// 		location:    "https://example.com",
		// 		contentType: "",
		// 	},
		// },
		{
			name: "negative test #1",
			id:   "invalid-id",
			want: want{
				code:        http.StatusBadRequest,
				location:    "",
				contentType: "",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Создаем тестовый запрос
			req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/"+test.id, bytes.NewReader([]byte("")))

			// Создаем тестовый обработчик
			handler := Init()

			// Выполняем запрос
			rr := httptest.NewRecorder()
			handler.Redirect(rr, req)

			// Проверяем код ответа
			assert.Equal(t, test.want.code, rr.Code)

			// Проверяем заголовок Location
			location := rr.Header().Get("Location")
			assert.Equal(t, test.want.location, location)

			// Проверяем тип контента
			assert.Equal(t, test.want.contentType, rr.Header().Get("Content-Type"))

			// Проверяем, что сгенерированная строка добавлена в хранилище
			if test.want.code == http.StatusTemporaryRedirect {
				require.NotNil(t, handler.storageInterface)
				value, err := handler.storageInterface.Get(test.id)
				require.NoError(t, err)
				assert.NotEmpty(t, value)
			}
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
