devrun:
	go run cmd/main.go
release:
	go build -o ./dist/prog cmd/main.go
prodrun:
	./dist/prog
