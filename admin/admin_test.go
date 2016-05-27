package admin

import "testing"

// For testing app reg
func TestRegApp(t *testing.T) {
	app1, err := RegApp("app1", "vapp1")
	if err != nil {
		t.Errorf("Register app failed reason: \n%s", err)
	}
	app, found := GetApp("app1")
	if app1 != app || !found {
		t.Errorf("Register app failed reason: \n%s", err)
	}

	app1, err = RegApp("app1", "vapp1")
	if err == nil {
		t.Errorf("Register app failed reason: \n%s", err)
	}
}
