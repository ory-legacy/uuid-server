package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"regexp"
)

func TestCreate(t *testing.T) {
	mockUuidCreate(t)
}

func BenchmarkTestCreate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = createUuid()
	}
}

func TestCreateUnique(t *testing.T) {
	id1 := mockUuidCreate(t)
	id2 := mockUuidCreate(t)
	assert.NotEqual(t, id1.Id, id2.Id)
}

func mockUuidCreate(t *testing.T) Id {
	m := mux.NewRouter()
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "http://example.com/", nil)
	assert.Nil(t, err)
	m.HandleFunc("/", createHandler).Methods("POST")
	m.ServeHTTP(recorder, req)

	assert.Nil(t, err)
	assert.Equal(t, 200, recorder.Code)
	assert.NotEmpty(t, recorder.Body)

	var e Id
	decoder := json.NewDecoder(recorder.Body)
	err = decoder.Decode(&e)

	valid, err := regexp.Match("[0-9a-z]{8}-[0-9a-z]{4}-4[0-9a-z]{3}-[8|9|a|b][0-9a-z]{3}-[0-9a-z]{12}", []byte(e.Id.String()))
	assert.Nil(t, err)
	assert.NotEmpty(t, e.Id)
	assert.True(t, valid)

	return e
}
