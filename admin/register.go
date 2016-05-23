package admin

import (
	"reflect"
)

type Model struct {
	name      string
	vname     string
	tname     string
	modelType reflect.Type
}
type App struct {
	name   string
	vname  string
	models map[string]*Model
}


// Use this interface to override the "kind" of the model (table name in datastore)
type tabler interface {
	TableName() string
}

// use this to set a verbose name for the model
type verboser interface {
	VerboseName() string
}

func RegApp(name string, args ... interface{}) (*App, error) {
	// Initializing _apps to an initial value
	if _apps == nil {
		_apps = make(map[string]*App)
	}

	// Extracting name, vname and making App
	var vname string
	if tem, found := args[0].(string); found {
		vname = tem
	}else {
		vname = name
	}
	app := &App{
		name: name,
		vname: vname,
		models: make(map[string]*Model),
	}

	// checking if name already exists to raise an error
	_, found := _apps[name]
	if found {
		return nil, AlreadyRegistered{"App" ,name}
	}

	//registering the app
	_apps[name] = app

	return app, nil
}

func GetApp(name string)(*App, bool){
	app, found := _apps[name]
	return app, found
}

func (a *App) RegModel(dat interface{}, args ...interface{}) error {
	// extracting name, table name, verbose name
	dtype := reflect.TypeOf(dat)
	name := dtype.Name()
	var tname string
	var vname string

	cusTname, found := dat.(tabler)
	if found {
		tname = cusTname.TableName()
	}else {
		tname = name
	}
	cusVname, found := dat.(verboser)
	if found {
		vname = cusVname.VerboseName()
	}else {
		vname = name
	}

	// generate a Model Object
	mod := &Model{
		name: name,
		vname: vname,
		tname: tname,
		modelType: dtype,
	}

	// store the model object or raise an error if already exists.
	_, found = a.models[name]
	if found{
		return AlreadyRegistered{
			"Model",
			name,
		}
	}else{
		a.models[name] = mod
	}
	return nil
}

