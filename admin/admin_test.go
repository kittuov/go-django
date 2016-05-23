package admin

import "testing"

func TestRegApp(t *testing.T) {
	app1, err := RegApp("app1", "vapp1")
	if err != nil {
		t.Errorf("Register app failed reason: \n%s", err)
	}
	if app1.name != "app1" || app1.vname!="vapp1"{
		t.Error("Not correct app registrations")
	}
}
