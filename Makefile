build:
	go build -ldflags="-X 'cabytes/notezine/wp.Version=1.1.1'" -o bin/wordpost main.go
	go build -buildmode=plugin -o bin/theme-light themes/light/main.go