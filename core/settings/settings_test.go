package settings

import (
	"testing"
	"github.com/kittuov/go-django/utils/log"
	"path/filepath"
)

const test_data = `
title: Krishna Chaitanya
apps:
  - admin
  - sessions
  - data
`
const test_data2 = `
title: Kornepati
data : testing the program
`

func TestUpdate(t *testing.T) {
	resetSettings()
	settings = make(settingsModel)
	err := Update([]byte(test_data))
	if err != nil {
		t.Fail()
	}
	if settings["title"] != "Krishna Chaitanya" {
		log.Error("wrong Title extracted")
		t.Fail()
	}
	apps, found := settings["apps"].([]interface{})
	if !found {
		log.Error("`apps` is not parsed as a list")
		t.Fail()
	}
	if len(apps) != 3 {
		log.Error("`apps` parsed but got wrong length")
	}
	if apps[2] != "data" {
		log.Error("wrong `apps` array extracted")
		t.Fail()
	}
}

func TestUpdateFile(t *testing.T) {
	resetSettings()
	testFile := "test_file.yaml"
	testFile, _ = filepath.Abs(testFile)
	err := UpdateFromFile(testFile)
	if err != nil {
		t.Fail()
	}
	config, found := settings["config"].(string)

	if !found {
		log.Error("config property Not set")
		t.Fail()
	}
	if config != "default_config" {
		log.Error("Wrong String was parsed and updated")
		t.Fail()
	}
}

func TestUpdate2(t *testing.T) {
	resetSettings()
	err := Update([]byte(test_data))
	if err != nil {
		t.Fail()
	}
	err = Update([]byte(test_data2))
	if err != nil {
		t.Fail()
	}
	title, found := settings["title"].(string)
	if !found {
		log.Error("title was unset")
		t.Fail()
	}
	if title != "Kornepati" {
		log.Error("Title not updated")
		t.Fail()
	}
	_, found = settings["data"].(string)
	if !found {
		log.Error("New params in updated settings are not set")
		t.Fail()
	}
}

func TestSetDefault(t *testing.T) {
	resetSettings()
	SetDefault([]byte(test_data))
	SetDefault([]byte(test_data2))
	if settings["title"] != "Krishna Chaitanya" {
		log.Error("title was also updated")
		t.Fail()
	}
	if _, found := settings["data"]; !found {
		log.Error("Unable to find 'data' key")
		t.Fail()
	}
}

func TestSetDefaultFromFile(t *testing.T) {
	resetSettings()
	testFile := "test_file.yaml"
	testFile, _ = filepath.Abs(testFile)
	err := UpdateFromFile(testFile)
	if err != nil {
		t.Fail()
	}
	config, found := settings["config"].(string)

	if !found {
		log.Error("config property Not set")
		t.Fail()
	}
	if config != "default_config" {
		log.Error("Wrong String was parsed and updated")
		t.Fail()
	}
}

func resetSettings() {
	log.Debug("Reset Settings")
	settings = make(settingsModel)
}