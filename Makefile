.PHONY: a1 a2

a1:
	GOOS=linux go build -o test cmd/assignment-1/main.go
	zip test.zip test
	rm test
a2:
	GOOS=linux go build -o test2 cmd/assignment-2/main.go
	zip test.zip test2
	rm test2