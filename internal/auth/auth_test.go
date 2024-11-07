package auth

import (
	"net/http/httptest"
	"testing"
)

func GetAPIKeyTest(test *testing.T) {
	header := httptest.NewRecorder().Header()
	res, err := GetAPIKey(header)

	if err == nil {
		test.Errorf("expected to have no data on Authorization Header, has %s", res)
	}

}
