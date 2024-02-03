package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenShouldBeFine(t *testing.T) {
	city := "moscow"
	count := 2

	req := httptest.NewRequest("GET", fmt.Sprintf("/cafe?count=%d&city=%s", count, city), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	body := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, cafeList[city][:count], body)
}

func TestMainHandlerWhenCityIsUnsupported(t *testing.T) {
	city := "TheCityThatNeverExisted"
	count := 2

	req := httptest.NewRequest("GET", fmt.Sprintf("/cafe?count=%d&city=%s", count, city), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	body := responseRecorder.Body.String()
	assert.Equal(t, "wrong city value", body)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	city := "moscow"
	moreThanTotalCount := 5

	req := httptest.NewRequest("GET", fmt.Sprintf("/cafe?count=%d&city=%s", moreThanTotalCount, city), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	body := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, cafeList[city], body)
}
