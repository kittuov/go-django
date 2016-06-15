package settings

import (
	"testing"
	"github.com/kittuov/go-django/utils/log"
)

const test_data = `
title: Krishna Chaitanya
apps:
  - admin
  - sessions
  - data
`

func TestUpdate(t *testing.T) {
	err := Update([]byte(test_data))
	if err != nil {
		log.Error(err)
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
	testFile := "test_file.yaml"
	err := UpdateFile(testFile)
	if err != nil {
		log.Error(err)
		t.Fail()
	}

}