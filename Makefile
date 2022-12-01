devrun:
	go run cmd/main.go
release:
	go build -mod=vendor -o ./dist/prog cmd/main.go
prodrun:
	./dist/prog
