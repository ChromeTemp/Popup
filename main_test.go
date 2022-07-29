package Popup

import (
	"testing"
)

// Tests must run manually on a Windows machine with UI access (cannot be tested with GitHub Actions).
func TestMain(t *testing.T) {
	Alert("Title", "Message")
	if Dialog("Title", "Message") {
		t.Log("Dialog answer: Yes")
	} else {
		t.Log("Dialog answer: No")
	}
}
