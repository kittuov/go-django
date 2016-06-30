package settings

import (
	"errors"
	"github.com/kittuov/go-django/utils/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type settingsModel map[string]interface{}

func (t settingsModel) String() string {
	out, err := yaml.Marshal(t)
	if err != nil {
		log.Warn("Settings can't be marshalled")
		return ""
	}
	return string(out)
}

var settings = make(settingsModel)

// UpdateFromFile is same as Update but reads from filename.
// in case of unsuccessful attempt, returns error. filename
// attribute has to be absolute file path
func UpdateFromFile(filename string) error {
	in, err := readFromFile(filename)
	if err != nil {
		log.Warnf("Not updating settings from file %s", filename)
		return err
	}
	log.Verbosef("Read Setting file %s", filename)
	err = Update(in)
	return err
}

// Update takes in a byteArray (as read from a file) and parses it as yml
// and updates the settings. in case it couldn't update the settings,
// returns error
func Update(yml []byte) error {
	log.Verbosef("Updating settings")
	data, err := getSettingsModel(yml)
	if err != nil {
		log.Warn("Not Updating settings")
		return err
	}
	for key, value := range data {
		settings[key] = value
	}
	log.Verbose("Successfully updated settings")
	log.Verbosef("New Settings are:\n%s", settings)

	return nil
}

// Sets a value to the settings if only the field is not already set
// returns error if couldn't do the job and if succeeded, returns
// nil
func SetDefault(yml []byte) error {
	log.Verbosef("Setting defaults")
	data, err := getSettingsModel(yml)
	if err != nil {
		log.Warn("Not Setting Defaults")
		return err
	}
	for key, value := range data {
		_, found := settings[key]
		if !found {
			settings[key] = value
		}
	}
	log.Verbose("Successfully set Defaults")
	log.Verbosef("New Settings are:\n %s", settings)
	return nil
}

// Same as SetDefault function but reads from file. it will either set defaults
// for all of the file or return error.
func SetDefaultFromFile(filename string) error {
	in, err := readFromFile(filename)
	if err != nil {
		log.Warn("unable to read from file")
		return err
	}
	err = SetDefault(in)
	return err
}

func readFromFile(filename string) (in []byte, err error) {
	is_abs := filepath.IsAbs(filename)
	if !is_abs {
		log.Errorf("Filename is not Absolute `%s`", filename)
		return nil, errors.New("Filename is not Absolute")
	}
	in, err = ioutil.ReadFile(filename)
	if err != nil {
		log.Errorf("Unable to open file with path `%s`", filename)
		return
	}
	return
}

func getSettingsModel(b []byte) (data settingsModel, err error) {
	data = make(settingsModel)
	log.Verbosef("Parsing yml data :\n%s", b)
	err = yaml.Unmarshal(b, &data)
	if err != nil {
		log.Warn("Unable to parse Data")
		return
	}
	log.Verbose("Successfully parsed yml")
	return
}
