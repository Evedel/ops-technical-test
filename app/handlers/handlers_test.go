package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Evedel/MYOBTestService/healthcheck"
	"github.com/Evedel/MYOBTestService/storage"
)

func TestGetRoot(t *testing.T) {
	expectedCode := http.StatusOK
	expectedBody := "hello world\n"
	endpointTest(t, root, "GET", "/", "handler.root()", expectedCode, expectedBody)
}

func TestGetHealthCheck(t *testing.T) {
	expectedCode := healthcheck.Run()
	expectedBody := ""
	endpointTest(t, healthCheck, "GET", "/healthcheck", "handler.healthCheck()", expectedCode, expectedBody)
}

func TestGetAppMetadata(t *testing.T) {
	expectedCode := http.StatusOK
	metadataInBytes, _ := json.Marshal(storage.GetAppMetadata())
	expectedBody := string(metadataInBytes) + "\n"
	endpointTest(t, appMetadata, "GET", "/metadata", "handler.appMetadata()", expectedCode, expectedBody)
}

func endpointTest(t *testing.T,
	handlerFunc func(w http.ResponseWriter, r *http.Request),
	method, path, handlerName string,
	expectedCode int, expectedBody string) {

	request, err := http.NewRequest(method, path, nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	http.HandlerFunc(handlerFunc).ServeHTTP(recorder, request)
	if recorder.Code != expectedCode {
		t.Errorf("%v received %v, expected %v", handlerName, recorder.Code, expectedCode)
	}

	if recorder.Body.String() != expectedBody {
		t.Errorf("%v received %v, expected %v", handlerName, recorder.Body.String(), expectedBody)
	}
}
