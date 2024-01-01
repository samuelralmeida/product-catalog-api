install-binares:
	go install github.com/mitranim/gow@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

run:
	gow -c -e=go,mod,gohtml run cmd/app/main.go