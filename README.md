Golang GORM migration tools


## Install
- Download binary package from [release page](https://github.com/phucngodev/gormcli/releases) and put it in PATH
- run gorm --help for more information

## Sample
```bash
gorm --help
gorm commandline tools

Usage:
  gorm [flags]
  gorm [command]

Available Commands:
  help        Help about any command
  init        Initialize project using gorm in current directory
  migrate     Run database migration

Flags:
      --config string   /Users/phuc/Projects/gormcli/config.yaml
  -h, --help            help for gorm
  -t, --toggle          Help message for toggle

Use "gorm [command] --help" for more information about a command.
Run database migration

Usage:
  gorm migrate [flags]
  gorm migrate [command]

Available Commands:
  create      Create migration file
  rollback    Rollback to previous version

Flags:
  -h, --help           help for migrate
  -v, --version uint   migrate database up/down to version

Global Flags:
      --config string   /Users/phuc/Projects/gormcli/config.yaml


```
