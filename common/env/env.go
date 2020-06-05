package env

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

type Environtment struct {
	data interface{}
}

func InitEnv() *Environtment {

	//default profile
	var profile = ""

	//check if there is input flag, if exits set flag to profile
	if len(os.Args) > 1 {
		profile = "-" + os.Args[1]
	}

	//open file config
	environtmentFile, err := os.Open("config" + profile + ".json")
	if err != nil {
		panic("Error, config" + profile + ".json Is Missing In Directory, " + err.Error())
	}

	//defer close file
	defer func() {
		if err := environtmentFile.Close(); err != nil {
			panic(err.Error())
		}
	}()

	//decode json file
	var temp interface{}
	jsonParser := json.NewDecoder(environtmentFile)
	err = jsonParser.Decode(&temp)
	if err != nil {
		panic(err.Error())
	}

	//create e
	return &Environtment{
		data: temp,
	}
}

func InitPath(path string) *Environtment {

	//open file config
	environtmentFile, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}

	//defer close file
	defer func() {
		if err := environtmentFile.Close(); err != nil {
			panic(err.Error())
		}
	}()

	//decode json file
	var temp interface{}
	jsonParser := json.NewDecoder(environtmentFile)
	err = jsonParser.Decode(&temp)
	if err != nil {
		panic(err.Error())
	}

	return &Environtment{
		data: temp,
	}
}

//get environtment value with str data type in linear
func (c *Environtment) Get(key string) string {
	m, err := c.doMapify()
	if err == nil {
		if val, ok := m[key]; ok {
			c := &Environtment{val}
			if s, ok := c.data.(string); ok {
				return s
			}
			panic("Error Conversion, Field Is Not String")
		}
	}
	return ""
}

//map environtment to map str interface
func (c *Environtment) doMapify() (map[string]interface{}, error) {
	if m, ok := c.data.(map[string]interface{}); ok {
		return m, nil
	}
	return nil, errors.New("can't type assert with map[str]interface{}")
}

func GetEnvPath() string {
	_, filename, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(filename, "../../"+string(filepath.Separator))))
	return apppath
}
