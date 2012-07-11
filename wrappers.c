#include <udis86.h>
#include "_cgo_export.h"

// These wrapers are used to set a go Reader as input.
int ud_read_go_reader(struct ud* ud) {
	void* ptr = ud_get_user_opaque_data(ud);
	return goRead(ptr);
}

void ud_set_input_reader(struct ud* ud, void* ptr) {
	ud_set_user_opaque_data(ud, ptr);
	ud_set_input_hook(ud, &ud_read_go_reader);
}

// These wrappers are used because we need to pass a function
// pointer as parameter of ud_set_syntax.
void ud_set_syntax_intel(struct ud* ud) {
	ud_set_syntax(ud, UD_SYN_INTEL);
}

void ud_set_syntax_att(struct ud* ud) {
	ud_set_syntax(ud, UD_SYN_ATT);
}

// These getters are necessary because there's no way to
// express the union type as a go type.
int8_t ud_get_lval_sbyte(struct ud_operand *op) {
	return op->lval.sbyte;
}

uint8_t ud_get_lval_ubyte(struct ud_operand *op) {
	return op->lval.ubyte;
}

int16_t ud_get_lval_sword(struct ud_operand *op) {
	return op->lval.sword;
}

uint16_t ud_get_lval_uword(struct ud_operand *op) {
	return op->lval.uword;
}

int32_t ud_get_lval_sdword(struct ud_operand *op) {
	return op->lval.sdword;
}

uint32_t ud_get_lval_udword(struct ud_operand *op) {
	return op->lval.udword;
}

int64_t ud_get_lval_sqword(struct ud_operand *op) {
	return op->lval.sqword;
}

uint64_t ud_get_lval_uqword(struct ud_operand *op) {
	return op->lval.uqword;
}

uint16_t ud_get_lval_ptr_seg(struct ud_operand *op) {
	return op->lval.ptr.seg;
}

uint32_t ud_get_lval_ptr_off(struct ud_operand *op) {
	return op->lval.ptr.off;
}
