package settings

import (
	"path/filepath"
	"github.com/kittuov/go-django/utils/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type setting map[string]interface{}

var set setting

func UpdateFile(filename string) error {
	file_path, err := filepath.Abs(filename)
	if err != nil {
		log.Errof("Unable to extract absolute path `%s`", file_path)
		log.Erro(err)
		return err
	}
	in, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Errof("Unable to open file with path `%s`", file_path)
		log.Erro(err)
		return err
	}
	log.Verbf("Read Setting file %s \n", file_path)
	Update(in)
	return nil
}

func Update(b []byte) error {
	log.Verbf("Updating settings with data:\n%s", b)
	err := yaml.Unmarshal(b, &set)
	if err != nil {
		log.Erro("Unable to unmarshall data :\n%s", b)
		return err
	}
	return nil

}