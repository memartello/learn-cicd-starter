package auth

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://testendpoint.com", nil)

	req.Header.Set("Authorization", "ApiKey your-token")

	got, err := GetAPIKey(req.Header)
	want := "your-token"

	if err != nil {
		t.Fatalf("got err %v trying to obtain header", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expectedL %v, got: %v", want, got)
	}
}
