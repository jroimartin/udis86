package udis86_test

import (
	"fmt"
	"github.com/jroimartin/udis86"
)

// This example demonstrates disassembling from buffer
func ExampleDisassemble() {
	data := []byte{'\x90', '\x48', '\x89', '\xfd', '\x55'}
	d := udis86.NewUDis86()
	d.SetInputBuffer(data)
	d.SetMode(64)
	d.SetSyntax(udis86.UD_SYN_INTEL)
	d.SetPC(0x400000)
	d.InputSkip(1)

	for d.Disassemble() != 0 {
		fmt.Printf("0x%08x %s %d %s\n",
			d.InsnOff(), d.InsnHex(), d.InsnLen(), d.InsnAsm())
	}
	// Output:
	// 0x00400000 4889fd 3 mov rbp, rdi
	// 0x00400003 55 1 push rbp
}
