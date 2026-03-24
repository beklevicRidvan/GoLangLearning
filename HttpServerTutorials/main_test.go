package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRoot(t *testing.T) {

	w := httptest.NewRecorder()

	handleRoot(w, nil)

	desiredCode := http.StatusOK

	assertResponse(t, w, desiredCode, "Welcome to our Homepage!\n")

}

func TestHandleGoodbye(t *testing.T) {

	w := httptest.NewRecorder()
	desiredCode := http.StatusOK

	handleGoodbye(w, nil)
	assertResponse(t, w, desiredCode, "Goodbye!")
}

func TestHandleHello(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?user=TestMan", nil)

	w := httptest.NewRecorder()

	handleHelloParameterized(w, req)

	desiredCode := http.StatusOK

	assertResponse(t, w, desiredCode, "Hello,Testman!\n")

}

func TestHandleHelloNoParameter(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello/", nil)

	w := httptest.NewRecorder()

	handleHelloParameterized(w, req)

	desiredCode := http.StatusOK

	assertResponse(t, w, desiredCode, "Hello, User!\n")

}
func TestHandleHelloWrongParamater(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?foo=bar", nil)

	w := httptest.NewRecorder()

	handleHelloParameterized(w, req)

	desiredCode := http.StatusOK

	assertResponse(t, w, desiredCode, "Hello, User!\n")

}

func TestHandleUserResponsesHello(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/responses/Testman/hello/", nil)

	req.SetPathValue("user", "TestMan")

	handleUserResponsesHello(w, req)

	desiredCode := http.StatusOK

	assertResponse(t, w, desiredCode, "Hello, TestMan!\n")

}

func TestHandleHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user/hello", nil)
	req.Header.Set("user", "Test Man")

	w := httptest.NewRecorder()

	handleHelloHeader(w, req)

	desiredCode := http.StatusOK

	assertResponse(t, w, desiredCode, "Hello, Test Man!\n")
}

func TestHandleNoHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user/hello", nil)
	//req.Header.Set("user", "Test Man")

	w := httptest.NewRecorder()

	handleHelloHeader(w, req)

	desiredCode := http.StatusBadRequest
	assertResponse(t, w, desiredCode, "invalid username provider\n")
}
func TestHandleJSON(t *testing.T) {
	testRequest := UserData{
		Name: "Test Man",
	}

	marshalledRequestBody, err := json.Marshal(testRequest)
	if err != nil {
		t.Fatalf("error marshalling test data: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/json", bytes.NewBuffer(marshalledRequestBody))
	w := httptest.NewRecorder()

	handleJson(w, req)

	desiredCode := http.StatusOK
	assertResponse(t, w, desiredCode, "Hello, Test Man!\n")
}

func TestHandleJSONWithEmptyBody(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/json", nil)
	w := httptest.NewRecorder()

	handleJson(w, req)

	desiredCode := http.StatusOK
	assertResponse(t, w, desiredCode, "Hello, Test Man!\n")
}
func TestHandleJSONWithEmptyName(t *testing.T) {
	testRequest := UserData{
		Name: "",
	}

	marshalledRequestBody, err := json.Marshal(testRequest)
	if err != nil {
		t.Fatalf("error marshalling test data: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/json", bytes.NewBuffer(marshalledRequestBody))
	w := httptest.NewRecorder()

	handleJson(w, req)

	desiredCode := http.StatusOK
	assertResponse(t, w, desiredCode, "HEllo, Test Man!\n")
}

func assertResponse(t *testing.T, w *httptest.ResponseRecorder, desiredCode int, expectedBody string) {
	t.Helper()

	if w.Code != desiredCode {
		t.Errorf("bad response code, expected %v but got %v\nbody: %s\n", desiredCode, w.Code, w.Body.String())
	}

	if w.Body.String() != expectedBody {
		t.Errorf("bad return, got %q, expected %q", w.Body.String(), expectedBody)
	}
}
