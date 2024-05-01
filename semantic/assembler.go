package semantic

import "os"

type assembler struct {
	asm string
}

var ASM = createAsmGenerator()

func createAsmGenerator() *assembler {
	a := new(assembler)
	a.asm = asmHeader
	return a
}

func (a *assembler) append(s string) {
	a.asm += s + "\n"
}

func (a *assembler) WriteToFile(filename string) {
	a.asm += asmFooter
	file, err := os.Create(filename)
	if err != nil {
		errorf("Error creating file %s", filename)
	}

	defer file.Close()

	_, err = file.WriteString(a.asm)
	if err != nil {
		errorf("Error writing to file %s", filename)
	}
}

const asmHeader = `
SYS_EXIT equ 1
SYS_READ equ 3
SYS_WRITE equ 4
STDIN equ 0
STDOUT equ 1
True equ 1
False equ 0

segment .data

formatin: db "%d", 0
formatout: db "%d", 10, 0 ; newline, nul terminator
scanint: times 4 db 0 ; 32-bits integer = 4 bytes

segment .bss  ; variaveis
res RESB 1

section .text
global main ; linux
extern scanf ; linux
extern printf ; linux
extern fflush ; linux
extern stdout ; linux

; subrotinas if/while
binop_je:
JE binop_true
JMP binop_false

binop_jg:
JG binop_true
JMP binop_false

binop_jl:
JL binop_true
JMP binop_false

binop_false:
MOV EAX, False  
JMP binop_exit
binop_true:
MOV EAX, True
binop_exit:
RET

main:

PUSH EBP ; guarda o base pointer
MOV EBP, ESP ; estabelece um novo base pointer
`

const asmFooter = `
PUSH DWORD [stdout]
CALL fflush
ADD ESP, 4

MOV ESP, EBP
POP EBP

MOV EAX, 1
XOR EBX, EBX
INT 0x80
`
