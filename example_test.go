package udis86_test

import (
	"fmt"
	"github.com/jroimartin/udis86"
	"os"
)

// This example demonstrates disassembling from buffer
func ExampleBuffer() {
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

// this example demonstrates disassembling from io.Reader
func ExampleReader() {
	f, err := os.Open("testdata/x86.bin")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	defer f.Close()
	d := udis86.NewUDis86()
	d.SetInputReader(f)
	d.SetMode(64)
	d.SetSyntax(udis86.UD_SYN_INTEL)
	d.SetPC(0x402a6f)

	for d.Disassemble() != 0 {
		fmt.Printf("0x%08x %s %d %s\n",
			d.InsnOff(), d.InsnHex(), d.InsnLen(), d.InsnAsm())
	}
	// Output:
	// 0x00402a6f e84cf7ffff 5 call 0x4021c0
	// 0x00402a74 4885c0 3 test rax, rax
	// 0x00402a77 4989c4 3 mov r12, rax
	// 0x00402a7a 7409 2 jz 0x402a85
}
