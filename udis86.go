// Packages udis86 provides access to the udis86 library
package udis86

/*
#cgo LDFLAGS: -ludis86
#cgo freebsd CFLAGS: -I/usr/local/include
#cgo freebsd LDFLAGS: -L/usr/local/lib

#include <udis86.h>

// These wrappers are implemented in wrappers.c
extern void ud_set_syntax_intel(struct ud* ud);
extern void ud_set_syntax_att(struct ud* ud);
extern void ud_set_input_reader(struct ud* ud, void* ptr);
extern int8_t ud_get_lval_sbyte(struct ud_operand *op);
extern uint8_t ud_get_lval_ubyte(struct ud_operand *op);
extern int16_t ud_get_lval_sword(struct ud_operand *op);
extern uint16_t ud_get_lval_uword(struct ud_operand *op);
extern int32_t ud_get_lval_sdword(struct ud_operand *op);
extern uint32_t ud_get_lval_udword(struct ud_operand *op);
extern int64_t ud_get_lval_sqword(struct ud_operand *op);
extern uint64_t ud_get_lval_uqword(struct ud_operand *op);
extern uint16_t ud_get_lval_ptr_seg(struct ud_operand *op);
extern uint32_t ud_get_lval_ptr_off(struct ud_operand *op);
*/
import "C"

import (
	"io"
	"unsafe"
)

// There are two inbuilt translators,
//	UD_SYN_INTEL - for INTEL (NASM-like) syntax.
//	UD_SYN_ATT - for AT&T (GAS-like) syntax.
// If you do not want udis86 to translate, you can pass
// UD_SYN_NONE to SetSyntax. This is particularly useful for
// cases when you only want to identify chunks of code and then
// create the assembly output if needed.
const (
	UD_SYN_NONE = iota
	UD_SYN_INTEL
	UD_SYN_ATT
)

const (
	UD_VENDOR_AMD = iota
	UD_VENDOR_INTEL
)

// UD_EOI is returned when the end of the input is reached
const UD_EOI = -1

// UDis86Ptr holds a value of type seg:off
type UDis86Ptr struct {
	Seg uint16
	Off uint32
}

// UDis86LVal is a representation of a decoded value
type UDis86LVal struct {
	SByte  int8
	UByte  uint8
	SWord  int16
	UWord  uint16
	SDword int32
	UDword uint32
	SQword int64
	UQword uint64
	Ptr    UDis86Ptr
}

// UDis86Operand contains the information decoded from each
// operand of the current instruction
type UDis86Operand struct {
	Type   int
	Size   uint8
	LVal   UDis86LVal
	Base   int
	Index  int
	Offset uint8
	Scale  uint8
}

// The UDis86 struct contais all the information decoded from
// the current instruction.
type UDis86 struct {
	ud       C.struct_ud
	r        io.Reader
	PC       uint64
	Mnemonic int
	Operand  [3]UDis86Operand
	PfxRex   uint8
	PfxSeg   uint8
	PfxOpr   uint8
	PfxAdr   uint8
	PfxLock  uint8
	PfxRep   uint8
	PfxRepe  uint8
	PfxRepne uint8
}

// NewUDis86 returns a new UDis86 object.
func NewUDis86() *UDis86 {
	d := new(UDis86)
	C.ud_init(&d.ud)
	return d
}

// fillOperandData fills the UDis86 Operand field, i is the index
// of the operand.
func (d *UDis86) fillOperandData(i int) {
	d.Operand[i].Type = int(d.ud.operand[i]._type)
	d.Operand[i].Size = uint8(d.ud.operand[i].size)
	d.Operand[i].Base = int(d.ud.operand[i].base)
	d.Operand[i].Index = int(d.ud.operand[i].index)
	d.Operand[i].Offset = uint8(d.ud.operand[i].offset)
	d.Operand[i].Scale = uint8(d.ud.operand[i].scale)
	d.Operand[i].LVal.SByte = int8(C.ud_get_lval_sbyte(&d.ud.operand[i]))
	d.Operand[i].LVal.UByte = uint8(C.ud_get_lval_ubyte(&d.ud.operand[i]))
	d.Operand[i].LVal.SWord = int16(C.ud_get_lval_sword(&d.ud.operand[i]))
	d.Operand[i].LVal.UWord = uint16(C.ud_get_lval_uword(&d.ud.operand[i]))
	d.Operand[i].LVal.SDword = int32(C.ud_get_lval_sdword(&d.ud.operand[i]))
	d.Operand[i].LVal.UDword = uint32(C.ud_get_lval_udword(&d.ud.operand[i]))
	d.Operand[i].LVal.SQword = int64(C.ud_get_lval_sqword(&d.ud.operand[i]))
	d.Operand[i].LVal.UQword = uint64(C.ud_get_lval_uqword(&d.ud.operand[i]))
	d.Operand[i].LVal.Ptr.Seg = uint16(C.ud_get_lval_ptr_seg(&d.ud.operand[i]))
	d.Operand[i].LVal.Ptr.Off = uint32(C.ud_get_lval_ptr_off(&d.ud.operand[i]))
}

// fillInsnData fills the UDis86 object with the decoded data.
func (d *UDis86) fillInsnData() {
	d.PC = uint64(d.ud.pc)
	d.Mnemonic = int(d.ud.mnemonic)
	d.PfxRex = uint8(d.ud.pfx_rex)
	d.PfxSeg = uint8(d.ud.pfx_seg)
	d.PfxOpr = uint8(d.ud.pfx_opr)
	d.PfxAdr = uint8(d.ud.pfx_adr)
	d.PfxLock = uint8(d.ud.pfx_lock)
	d.PfxRep = uint8(d.ud.pfx_rep)
	d.PfxRepe = uint8(d.ud.pfx_repe)
	d.PfxRepne = uint8(d.ud.pfx_repne)
	for i := range d.Operand {
		d.fillOperandData(i)
	}
}

// Disassemble disassembles one instruction and returns the
// number of bytes disassembled. A Zero means end of
// disassembly.
func (d *UDis86) Disassemble() uint {
	r := uint(C.ud_disassemble(&d.ud))
	d.fillInsnData()
	return r
}

// InputSkip skips n input bytes.
func (d *UDis86) InputSkip(n uint) {
	C.ud_input_skip(&d.ud, C.size_t(n))
}

// InsnAsm returns a string with the disassembled instruction.
func (d *UDis86) InsnAsm() string {
	return C.GoString(C.ud_insn_asm(&d.ud))
}

// InsnHex returns a string with the hex form of the
// disassembled instruction.
func (d *UDis86) InsnHex() string {
	return C.GoString(C.ud_insn_hex(&d.ud))
}

// InsnLen returns the number of bytes disassembled.
func (d *UDis86) InsnLen() uint {
	return uint(C.ud_insn_len(&d.ud))
}

// InsnOff Returns the starting offset of the disassembled
// instruction relative to the program counter value specified
// initially.
func (d *UDis86) InsnOff() uint64 {
	return uint64(C.ud_insn_off(&d.ud))
}

// goRead is called from ud_read_go_reader to read a single byte
// from the io.Reader.
//export goRead
func goRead(ptr unsafe.Pointer) int {
	b := make([]byte, 1)
	d := (*UDis86)(ptr)
	_, err := d.r.Read(b)
	if err == io.EOF {
		return UD_EOI
	}
	return int(b[0])
}

// SetInputReader sets an io.Reader as input.
func (d *UDis86) SetInputReader(r io.Reader) {
	d.r = r
	C.ud_set_input_reader(&d.ud, unsafe.Pointer(d))
}

// SetInputBuffer sets a byte slice as input.
func (d *UDis86) SetInputBuffer(b []byte) {
	C.ud_set_input_buffer(&d.ud, (*C.uint8_t)(&b[0]),
		C.size_t(len(b)))
}

// SetMode sets disassembly mode. Valid values are 16, 32
// and 64. By default, the library works in 32bit mode.
func (d *UDis86) SetMode(m uint8) {
	C.ud_set_mode(&d.ud, C.uint8_t(m))
}

// SetPC sets the program counter (EIP/RIP).
func (d *UDis86) SetPC(pc uint64) {
	C.ud_set_pc(&d.ud, C.uint64_t(pc))
}

// SetSyntax sets the output syntax.
func (d *UDis86) SetSyntax(s int) {
	switch s {
	case UD_SYN_NONE:
		C.ud_set_syntax(&d.ud, nil)
	case UD_SYN_INTEL:
		C.ud_set_syntax_intel(&d.ud)
	case UD_SYN_ATT:
		C.ud_set_syntax_att(&d.ud)
	}
}

// SetVendor sets the vendor of whose instruction to choose
// from. This is only useful for selecting the VMX or SVM
// instruction sets at which point INTEL and AMD have diverged
// significantly.
func (d *UDis86) SetVendor(v int) {
	switch v {
	case UD_VENDOR_INTEL:
		C.ud_set_vendor(&d.ud, C.UD_VENDOR_INTEL)
	case UD_VENDOR_AMD:
		C.ud_set_vendor(&d.ud, C.UD_VENDOR_AMD)
	}
}
