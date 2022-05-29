.PHONY: build

build: 
	go build -o ino ./src/main.go

clean:
	rm -v ino