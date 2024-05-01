run: build
	./compiler $(file).lua
	nasm -f elf -o program.o $(file).asm
	gcc -m32 -no-pie -o program program.o
	./program

test:
	go test compiler/tokenizer compiler/preprocessor -count=1

build:
	go build -o compiler main.go

clean:
	rm -f compiler
	rm -f program
	rm -f program.o
	rm -f *.asm
