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

func TestGetAlbumsEmpty(t *testing.T) {
	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	var response map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "no albums found", response["error"])
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

func TestGetAlbumByID(t *testing.T) {
	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)
	var albums []map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &albums)
	assert.Nil(t, err)
	id := albums[0]["_id"].(string)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/albums/"+id, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var album map[string]any
	err = json.Unmarshal(w.Body.Bytes(), &album)
	assert.Nil(t, err)
	expectedResponse := `{"_id": "` + id + `", "title": "Um álbum qualquer", "artist": "Fulano de Tal", "price": 2.99}`
	assert.JSONEq(t, expectedResponse, string(w.Body.String()))
}

func TestUpdateAlbum(t *testing.T) {
	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.PUT("/albums/:id", controllers.UpdateAlbum)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)
	var albums []map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &albums)
	assert.Nil(t, err)
	id := albums[0]["_id"].(string)

	updatedAlbum := map[string]any{"title": "Um álbum atualizado", "artist": "Ciclano de Tal", "price": 3.99}
	jsonValue, _ := json.Marshal(updatedAlbum)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/albums/"+id, bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/albums/"+id, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var album map[string]any
	err = json.Unmarshal(w.Body.Bytes(), &album)
	assert.Nil(t, err)
	expectedResponse := `{"_id": "` + id + `", "title": "Um álbum atualizado", "artist": "Ciclano de Tal", "price": 3.99}`
	assert.JSONEq(t, expectedResponse, string(w.Body.String()))
}

func TestDeleteAlbum(t *testing.T) {
	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)
	var albums []map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &albums)
	assert.Nil(t, err)
	id := albums[0]["_id"].(string)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/albums/"+id, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/albums/"+id, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
	var response map[string]any
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "album not found", response["error"])
}
