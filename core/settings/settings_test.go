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
		log.Erro(err)
		t.Fail()
	}
	if set["title"] != "Krishna Chaitanya" {
		log.Erro("wrong Title extracted")
		t.Fail()
	}
	apps, found := set["apps"].([]interface{})
	if !found {
		log.Erro("`apps` is not parsed as a list")
		t.Fail()
	}
	if len(apps) != 3 {
		log.Erro("`apps` parsed but got wrong length")
	}
	if apps[2] != "data" {
		log.Erro("wrong `apps` array extracted")
		t.Fail()
	}

}

func TestUpdateFile(t *testing.T) {
	testFile := "test_file.yaml"
	err := UpdateFile(testFile)
	if err != nil {
		log.Erro(err)
		t.Fail()
	}

}