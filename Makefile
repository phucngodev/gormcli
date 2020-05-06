build:
	goreleaser

osx:
	go build -o ~/Projects/go/bin/gorm main.go

clean:
	@rm -rf dist
