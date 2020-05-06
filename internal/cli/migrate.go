package cli

import (
	"fmt"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// InitMigrate init go-migrate
func InitMigrate() (*migrate.Migrate, error) {
	env := viper.GetString("app_env")

	dbConfig := viper.GetStringMap(env)
	var err error
	var dbUri string
	var m *migrate.Migrate
	dbType := dbConfig["dialect"].(string)
	switch dbType {
	case "mysql":
		dbUri = fmt.Sprintf("mysql//%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", dbConfig["user"], dbConfig["password"], dbConfig["host"], dbConfig["port"], dbConfig["database"])
		m, err = migrate.New("file://migrations", dbUri)
		if err != nil {
			return nil, errors.Wrap(err, "failed to created migration object")
		}
	case "postgres":
		dbUri = fmt.Sprintf("postgres//%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", dbConfig["user"], dbConfig["password"], dbConfig["host"], dbConfig["port"], dbConfig["database"])
		m, err = migrate.New("file://migrations", dbUri)
		if err != nil {
			return nil, errors.Wrap(err, "failed to created migration object")
		}
	default:
		return nil, errors.New("Unsupported database")
	}

	return m, nil
}
