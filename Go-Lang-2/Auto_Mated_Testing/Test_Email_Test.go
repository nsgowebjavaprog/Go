package Auto_Mated_Testing

import "testing"

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("Hello is not an email")
	}
}
