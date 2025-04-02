package tests

import (
	"api/controllers"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetInitial(t *testing.T) {
	router := gin.Default()
	router.GET("/", controllers.GetInitial)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Welcome to the API Managerment Albums", response["message"])
}

func TestPostAlbums(t *testing.T) {
	router := gin.Default()
	router.POST("/albums", controllers.PostAlbums)

	newAlbum := map[string]any{"title": "Um álbum qualquer", "artist": "Fulano de Tal", "price": 2.99}
	jsonValue, _ := json.Marshal(newAlbum)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	_, err := io.ReadAll(w.Body)
	assert.Nil(t, err)
}

func TestGetAlbums(t *testing.T) {
	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response []map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	delete(response[0], "_id")
	updateResponse, err := json.Marshal(response[0])
	assert.Nil(t, err)
	expectedResponse := `{"title": "Um álbum qualquer", "artist": "Fulano de Tal", "price": 2.99}`
	assert.JSONEq(t, expectedResponse, string(updateResponse))
}
