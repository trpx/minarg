package errors

import "testing"

func TestErrors(t *testing.T) {

	err := NewParseError(
		1,
		"msg %s",
		"abc",
	)

	expectedCode := 1
	gotCode := err.Code()
	if expectedCode != gotCode {
		t.Errorf("expected error code %d, got %d", expectedCode, gotCode)
	}

	expectedMsg := "msg abc"
	gotMsg := err.Error()
	if gotMsg != expectedMsg {
		t.Errorf("expected '%s' from err.Error(), got '%s'", expectedMsg, gotMsg)
	}
}
