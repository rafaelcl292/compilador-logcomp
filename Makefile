test:
	go test compiler/tokenizer compiler/preprocessor compiler -count=1

build:
	go build -o compiler main.go

run: build
	./compiler

clean:
	rm -f compiler
