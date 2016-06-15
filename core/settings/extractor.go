package settings

import (
	"path/filepath"
	"github.com/kittuov/go-django/utils/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type settingModel map[string]interface{}

var settings settingModel

func UpdateFile(filename string) error {
	file_path, err := filepath.Abs(filename)
	if err != nil {
		log.Errorf("Unable to extract absolute path `%s`", file_path)
		log.Error(err)
		return err
	}
	in, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Errorf("Unable to open file with path `%s`", file_path)
		log.Error(err)
		return err
	}
	log.Verbosef("Read Setting file %s \n", file_path)
	Update(in)
	return nil
}

func Update(b []byte) error {
	log.Verbosef("Updating settings with data:\n%s", b)
	err := yaml.Unmarshal(b, &settings)
	if err != nil {
		log.Error("Unable to unmarshall data :\n%s", b)
		return err
	}
	return nil

}