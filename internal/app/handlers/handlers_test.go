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

// !func Test_handlerWrapper_IndexPage(t *testing.T) {
// 	type args struct {
// 		res http.ResponseWriter
// 		req *http.Request
// 	}
// 	tests := []struct {
// 		name string
// 		hw   handlerWrapper
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			//tt.hw.IndexPage(tt.args.res, tt.args.req)
// 		})
// 	}
// }

func Test_handlerWrapper_IndexPage(t *testing.T) {
	type want struct {
		code int
		//response    string
		contentType string
	}
	tests := []struct {
		name    string
		reqBody string
		want    want
	}{
		{
			name:    "positive test #1",
			reqBody: "http://site1.com",
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
			reqreqBody := strings.NewReader(string(test.reqBody))
			request := httptest.NewRequest(http.MethodPost, "/", reqreqBody)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			hw.IndexPage(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resreqBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			require.NotNil(t, resreqBody)
			assert.Equal(t, test.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}

// !func Test_handlerWrapper_Redirect(t *testing.T) {
// 	type want struct {
// 		code     int
// 		response string
// 	}
// 	tests := []struct {
// 		name string
// 		want want
// 	}{
// 		{
// 			name: "positive test #1",
// 			want: want{
// 				code:     307,
// 				response: "jhvmj",
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
// 			defer res.reqBody.Close()
// 			resreqBody, err := io.ReadAll(res.reqBody)

// 			require.NoError(t, err)
// 			assert.Equal(t, test.want.response, string(resreqBody))

// 		})
// 	}
// }

// !func TestIndexPage(t *testing.T) {
// 	hw := Init()

// 	reqreqBody := []byte("https://example.com")
// 	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(reqreqBody))
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	w := httptest.NewRecorder()

// 	hw.IndexPage(w, req)

// 	resp := w.Result()
// 	reqBody := make([]byte, resp.ContentLength)
// 	resp.reqBody.Read(reqBody)

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// 	assert.Equal(t, "text/plain", resp.Header.Get("content-type"))
// 	assert.Contains(t, string(reqBody), "localhost")
// 	assert.Equal(t, 31, len(string(reqBody)))

// 	value, _ := hw.storageInterface.Get("rez1")
// 	assert.Equal(t, string(reqreqBody), value)
// }

// !func TestIndexPage(t *testing.T) {
// 	// Создаем тестовый запрос
// 	req := httptest.NewRequest(http.MethodPost, localhost, bytes.NewReader([]byte("https://example.com")))

// 	// Создаем тестовый обработчик
// 	handler := Init()

// 	// Выполняем запрос
// 	rr := httptest.NewRecorder()
// 	handler.IndexPage(rr, req)

// 	// Проверяем код ответа
// 	assert.Equal(t, http.StatusCreated, rr.Code)

// 	// Проверяем длину сгенерированной строки
// 	reqBody := rr.reqBody.String()
// 	assert.Len(t, reqBody, 6)

// 	// Проверяем, что сгенерированная строка начинается с localhost
// 	assert.True(t, strings.HasPrefix(reqBody, localhost))

// 	// Проверяем, что сгенерированная строка добавлена в хранилище
// 	require.NotNil(t, handler.storageInterface)
// 	value, err := handler.storageInterface.Get(reqBody[len(localhost):])
// 	require.NoError(t, err)
// 	assert.Equal(t, "https://example.com", value)
// }

func TestIndexPage(t *testing.T) {
	// Создаем тестовый запрос
	req := httptest.NewRequest(http.MethodPost, localhost, bytes.NewReader([]byte("https://example.com")))

	// Создаем тестовый обработчик
	handler := Init()

	// Выполняем запрос
	rr := httptest.NewRecorder()
	handler.IndexPage(rr, req)

	// Проверяем код ответа
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Проверяем длину сгенерированной строки
	reqBody := rr.Body.String()
	assert.Len(t, reqBody[len(localhost):], 6)

	// Проверяем, что сгенерированная строка начинается с localhost
	assert.True(t, strings.HasPrefix(reqBody, localhost))

	// Проверяем, что сгенерированная строка добавлена в хранилище
	require.NotNil(t, handler.storageInterface)
	value, err := handler.storageInterface.Get(reqBody[len(localhost):])
	require.NoError(t, err)
	assert.NotEmpty(t, value)
}

// !func TestIndexPage(t *testing.T) {
// 	type want struct {
// 		code        int
// 		response    string
// 		contentType string
// 	}

// 	tests := []struct {
// 		name string
// 		want want
// 	}{
// 		{
// 			name: "positive test #1",
// 			want: want{
// 				code:        http.StatusCreated,
// 				response:    localhost + "123456",
// 				contentType: "text/plain",
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			// Создаем тестовый запрос
// 			req := httptest.NewRequest(http.MethodPost, localhost, bytes.NewReader([]byte("https://example.com")))

// 			// Создаем тестовый обработчик
// 			handler := Init()

// 			// Выполняем запрос
// 			rr := httptest.NewRecorder()
// 			handler.IndexPage(rr, req)

// 			// Проверяем код ответа
// 			assert.Equal(t, test.want.code, rr.Code)

// 			// Проверяем длину сгенерированной строки
// 			reqBody := rr.reqBody.String()
// 			assert.Len(t, reqBody[len(localhost):], 6)

// 			// Проверяем тип контента
// 			assert.Equal(t, test.want.contentType, rr.Header().Get("Content-Type"))

// 			// Проверяем, что сгенерированная строка добавлена в хранилище
// 			require.NotNil(t, handler.storageInterface)
// 			value, err := handler.storageInterface.Get(reqBody[len(localhost):])
// 			require.NoError(t, err)
// 			assert.NotEmpty(t, value)
// 		})
// 	}
// }

// !func TestRedirect(t *testing.T) {
// 	type want struct {
// 		code        int
// 		location    string
// 		contentType string
// 	}

// 	tests := []struct {
// 		name string
// 		id   string
// 		want want
// 	}{
// 		{
// 			name: "positive test #1",
// 			id:   "123456",
// 			want: want{
// 				code:        http.StatusTemporaryRedirect,
// 				location:    "https://example.com",
// 				contentType: "",
// 			},
// 		},
// 		{
// 			name: "negative test #1",
// 			id:   "invalid-id",
// 			want: want{
// 				code:        http.StatusBadRequest,
// 				location:    "",
// 				contentType: "",
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			// Создаем тестовый запрос
// 			req := httptest.NewRequest(http.MethodGet, localhost+test.id, bytes.NewReader([]byte("")))

// 			// Создаем тестовый обработчик
// 			handler := Init()

// 			// Выполняем запрос
// 			rr := httptest.NewRecorder()
// 			handler.Redirect(rr, req)

// 			// Проверяем код ответа
// 			assert.Equal(t, test.want.code, rr.Code)

// 			// Проверяем заголовок Location
// 			location := rr.Header().Get("Location")
// 			assert.Equal(t, test.want.location, location)

// 			// Проверяем тип контента
// 			assert.Equal(t, test.want.contentType, rr.Header().Get("Content-Type"))

// 			// Проверяем, что сгенерированная строка добавлена в хранилище
// 			if test.want.code == http.StatusTemporaryRedirect {
// 				require.NotNil(t, handler.storageInterface)
// 				value, err := handler.storageInterface.Get(test.id)
// 				require.NoError(t, err)
// 				assert.NotEmpty(t, value)
// 			}
// 		})
// 	}
// }

// !func TestRedirect(t *testing.T) {
// 	type want struct {
// 		code        int
// 		location    string
// 		contentType string
// 	}

// 	tests := []struct {
// 		name string
// 		id   string
// 		want want
// 	}{
// 		{
// 			name: "positive test #1",
// 			id:   "123456",
// 			want: want{
// 				code:        http.StatusTemporaryRedirect,
// 				location:    "https://example.com",
// 				contentType: "",
// 			},
// 		},
// 		{
// 			name: "negative test #1",
// 			id:   "invalid-id",
// 			want: want{
// 				code:        http.StatusBadRequest,
// 				location:    "",
// 				contentType: "",
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			// Создаем тестовый запрос
// 			req := httptest.NewRequest(http.MethodGet, localhost+test.id, bytes.NewReader([]byte("")))

// 			// Создаем тестовый обработчик
// 			handler := Init()

// 			// Добавляем сгенерированную строку в хранилище
// 			if test.want.code == http.StatusTemporaryRedirect {
// 				require.NotNil(t, handler.storageInterface)
// 				handler.storageInterface.Add(test.id, test.want.location)
// 			}

// 			// Выполняем запрос
// 			rr := httptest.NewRecorder()
// 			handler.Redirect(rr, req)

// 			// Проверяем код ответа
// 			assert.Equal(t, test.want.code, rr.Code)

// 			// Проверяем заголовок Location
// 			location := rr.Header().Get("Location")
// 			assert.Equal(t, test.want.location, location)

// 			// Проверяем тип контента
// 			assert.Equal(t, test.want.contentType, rr.Header().Get("Content-Type"))
// 		})
// 	}
// }
