.PHONEY: clean

build/server: main.go
	go build -o build/server

clean:
	rm -r build/
