install-binares:
	go install github.com/mitranim/gow@latest

run:
	gow -c -e=go,mod,gohtml run main.go