package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/pkg/errors"
)

// Config config data structure
type Config struct {
	AppName string
	DBType  string
	DBPort  string
}

var configTpl = `
---
app_name: "{{ .AppName }}"
app_env: "development"

development:
  dialect: "{{ .DBType }}"
  database: "{{ .AppName }}_development"
  host: "localhost"
  port: "{{ .DBPort }}"
  user: "dev"
  password: "dev12345"


# production:
#  dialect: "{{ .DBType }}"
#  database: "{{ .AppName }}_development"
#  host: "localhost"
#  port: "{{ .DBPort }}"
#  user: "dev"
#  password: "dev12345"
`

func createMigrationDirectory() error {
	pwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "[ConfigTemplate] failed to get working directory")
	}

	err = os.Mkdir(fmt.Sprintf("%s/migrations", pwd), os.ModePerm)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return errors.Wrap(err, "failed to create migrations directory")
	}

	return nil
}

// createConfigFile create database config file
func createConfigFile(dbType string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "[ConfigTemplate] failed to get app name")
	}

	var dbPort string
	if dbType == "mysql" {
		dbPort = "3306"
	} else if dbType == "postgres" {
		dbPort = "5432"
	}

	appConfig := Config{
		AppName: filepath.Base(pwd),
		DBType:  dbType,
		DBPort:  dbPort,
	}

	f, err := os.Create("config.yaml")
	if err != nil {
		return errors.Wrap(err, "[ConfigTemplate] failed to open config file")
	}
	defer f.Close()

	t := template.Must(template.New("config").Parse(configTpl))
	err = t.Execute(f, appConfig)
	if err != nil {
		return errors.Wrap(err, "[ConfigTemplate] failed to write config file")
	}

	return nil
}
