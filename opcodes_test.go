package main

import (
	"strconv"
	"strings"
	"testing"
)

func DisplayInstruction(instruction Instruction) string {
	return instruction.Mnemonic + " " + strings.Join(instruction.Operands, " ")
}

func PrettyOutputArray(output []byte) string {
	outStr := "["
	for i, c := range output {
		if i != 0 {
			outStr += ", "
		}
		outStr += strconv.Itoa(int(c))
	}
	outStr += "]"
	return outStr
}

func TryTestInput(t *testing.T, instruction Instruction, expectedOutput []byte) {
	output := OpCodes_GetOutput(instruction, "test", 0)
	if !Utils_ByteSlicesEqual(output, expectedOutput) {
		t.Errorf("Instruction '%s' assembled to %s, should have been %s", DisplayInstruction(instruction), PrettyOutputArray(output), PrettyOutputArray(expectedOutput))
	}
}

func TestControlInstructions(t *testing.T) {
	TryTestInput(t, Instruction{"CALL", []string{"1234"}}, []byte{0xCD, 0xD2, 0x04})
	TryTestInput(t, Instruction{"CALL", []string{"Z", "1234"}}, []byte{0xCC, 0xD2, 0x04})

	TryTestInput(t, Instruction{"JP", []string{"1234"}}, []byte{0xC3, 0xD2, 0x04})
	TryTestInput(t, Instruction{"JP", []string{"HL"}}, []byte{0xE9})
	TryTestInput(t, Instruction{"JP", []string{"[HL]"}}, []byte{0xE9})
	TryTestInput(t, Instruction{"JP", []string{"Z", "1234"}}, []byte{0xCA, 0xD2, 0x04})

	TryTestInput(t, Instruction{"RET", []string{}}, []byte{0xC9})
	TryTestInput(t, Instruction{"RET", []string{"Z"}}, []byte{0xC8})
	TryTestInput(t, Instruction{"RETI", []string{}}, []byte{0xD9})

	TryTestInput(t, Instruction{"RST", []string{"0x00"}}, []byte{0xC7})
	TryTestInput(t, Instruction{"RST", []string{"0x08"}}, []byte{0xCF})
	TryTestInput(t, Instruction{"RST", []string{"0x10"}}, []byte{0xD7})
	TryTestInput(t, Instruction{"RST", []string{"0x18"}}, []byte{0xDF})
	TryTestInput(t, Instruction{"RST", []string{"0x20"}}, []byte{0xE7})
	TryTestInput(t, Instruction{"RST", []string{"0x28"}}, []byte{0xEF})
	TryTestInput(t, Instruction{"RST", []string{"0x30"}}, []byte{0xF7})
	TryTestInput(t, Instruction{"RST", []string{"0x38"}}, []byte{0xFF})
}

func TestBitInstructions(t *testing.T) {
	TryTestInput(t, Instruction{"BIT", []string{"1", "A"}}, []byte{0xCB, 0x4F})
	TryTestInput(t, Instruction{"BIT", []string{"2", "B"}}, []byte{0xCB, 0x50})
	TryTestInput(t, Instruction{"BIT", []string{"3", "[HL]"}}, []byte{0xCB, 0x5E})

	TryTestInput(t, Instruction{"RES", []string{"1", "A"}}, []byte{0xCB, 0x8F})
	TryTestInput(t, Instruction{"RES", []string{"2", "B"}}, []byte{0xCB, 0x90})
	TryTestInput(t, Instruction{"RES", []string{"3", "[HL]"}}, []byte{0xCB, 0x9E})

	TryTestInput(t, Instruction{"SET", []string{"1", "A"}}, []byte{0xCB, 0xCF})
	TryTestInput(t, Instruction{"SET", []string{"2", "B"}}, []byte{0xCB, 0xD0})
	TryTestInput(t, Instruction{"SET", []string{"3", "[HL]"}}, []byte{0xCB, 0xDE})
}

func TestLoadInstructions(t *testing.T) {
	TryTestInput(t, Instruction{"LD", []string{"A", "A"}}, []byte{0x7f})
	TryTestInput(t, Instruction{"LD", []string{"A", "B"}}, []byte{0x78})
	TryTestInput(t, Instruction{"LD", []string{"A", "C"}}, []byte{0x79})
	TryTestInput(t, Instruction{"LD", []string{"A", "D"}}, []byte{0x7a})
	TryTestInput(t, Instruction{"LD", []string{"A", "E"}}, []byte{0x7b})
	TryTestInput(t, Instruction{"LD", []string{"A", "H"}}, []byte{0x7c})
	TryTestInput(t, Instruction{"LD", []string{"A", "L"}}, []byte{0x7d})
	TryTestInput(t, Instruction{"LD", []string{"B", "66"}}, []byte{0x06, 0x42})
	TryTestInput(t, Instruction{"LD", []string{"BC", "1234"}}, []byte{0x01, 0xD2, 0x04})
	TryTestInput(t, Instruction{"LD", []string{"[HL]", "66"}}, []byte{0x36, 0x42})
	TryTestInput(t, Instruction{"LD", []string{"[BC]", "A"}}, []byte{0x02})
	TryTestInput(t, Instruction{"LD", []string{"[DE]", "A"}}, []byte{0x12})
	TryTestInput(t, Instruction{"LD", []string{"[HL]", "A"}}, []byte{0x77})
	TryTestInput(t, Instruction{"LD", []string{"A", "[HL]"}}, []byte{0x7E})
	TryTestInput(t, Instruction{"LD", []string{"[1234]", "A"}}, []byte{0x32, 0xD2, 0x04})
	TryTestInput(t, Instruction{"LD", []string{"A", "[1234]"}}, []byte{0x3A, 0xD2, 0x04})
}

func TestALUInstructions(t *testing.T) {
	TryTestInput(t, Instruction{"ADD", []string{"A", "66"}}, []byte{0xC6, 0x42})
	TryTestInput(t, Instruction{"ADD", []string{"A", "B"}}, []byte{0x80})
	TryTestInput(t, Instruction{"ADD", []string{"A", "[HL]"}}, []byte{0x86})

	TryTestInput(t, Instruction{"ADD", []string{"HL", "BC"}}, []byte{0x09})
	TryTestInput(t, Instruction{"ADD", []string{"HL", "DE"}}, []byte{0x19})
	TryTestInput(t, Instruction{"ADD", []string{"HL", "HL"}}, []byte{0x29})
	TryTestInput(t, Instruction{"ADD", []string{"HL", "SP"}}, []byte{0x39})

	TryTestInput(t, Instruction{"ADC", []string{"A", "66"}}, []byte{0xCE, 0x42})
	TryTestInput(t, Instruction{"ADC", []string{"A", "B"}}, []byte{0x88})
	TryTestInput(t, Instruction{"ADC", []string{"A", "[HL]"}}, []byte{0x8E})

	TryTestInput(t, Instruction{"SUB", []string{"66"}}, []byte{0xD6, 0x42})
	TryTestInput(t, Instruction{"SUB", []string{"B"}}, []byte{0x90})
	TryTestInput(t, Instruction{"SUB", []string{"[HL]"}}, []byte{0x96})

	TryTestInput(t, Instruction{"SBC", []string{"A", "66"}}, []byte{0xDE, 0x42})
	TryTestInput(t, Instruction{"SBC", []string{"A", "B"}}, []byte{0x98})
	TryTestInput(t, Instruction{"SBC", []string{"A", "[HL]"}}, []byte{0x9E})

	TryTestInput(t, Instruction{"AND", []string{"66"}}, []byte{0xE6, 0x42})
	TryTestInput(t, Instruction{"AND", []string{"B"}}, []byte{0xA0})
	TryTestInput(t, Instruction{"AND", []string{"[HL]"}}, []byte{0xA6})

	TryTestInput(t, Instruction{"XOR", []string{"66"}}, []byte{0xEE, 0x42})
	TryTestInput(t, Instruction{"XOR", []string{"B"}}, []byte{0xA8})
	TryTestInput(t, Instruction{"XOR", []string{"[HL]"}}, []byte{0xAE})

	TryTestInput(t, Instruction{"OR", []string{"66"}}, []byte{0xF6, 0x42})
	TryTestInput(t, Instruction{"OR", []string{"B"}}, []byte{0xB0})
	TryTestInput(t, Instruction{"OR", []string{"[HL]"}}, []byte{0xB6})

	TryTestInput(t, Instruction{"CP", []string{"66"}}, []byte{0xFE, 0x42})
	TryTestInput(t, Instruction{"CP", []string{"A"}}, []byte{0xBF})
	TryTestInput(t, Instruction{"CP", []string{"[HL]"}}, []byte{0xBE})
}

func TestMathInstructions(t *testing.T) {
	TryTestInput(t, Instruction{"DEC", []string{"A"}}, []byte{0x3D})
	TryTestInput(t, Instruction{"DEC", []string{"HL"}}, []byte{0x2B})
	TryTestInput(t, Instruction{"INC", []string{"A"}}, []byte{0x3C})
	TryTestInput(t, Instruction{"INC", []string{"BC"}}, []byte{0x03})
	TryTestInput(t, Instruction{"INC", []string{"HL"}}, []byte{0x23})
	TryTestInput(t, Instruction{"INC", []string{"[HL]"}}, []byte{0x34})
}

func TestROTInstructions(t *testing.T) {
	TryTestInput(t, Instruction{"RLC", []string{"A"}}, []byte{0xCB, 0x07})
	TryTestInput(t, Instruction{"RLC", []string{"B"}}, []byte{0xCB, 0x00})
	TryTestInput(t, Instruction{"RLC", []string{"[HL]"}}, []byte{0xCB, 0x06})

	TryTestInput(t, Instruction{"RRC", []string{"A"}}, []byte{0xCB, 0x0F})
	TryTestInput(t, Instruction{"RRC", []string{"B"}}, []byte{0xCB, 0x08})
	TryTestInput(t, Instruction{"RRC", []string{"[HL]"}}, []byte{0xCB, 0x0E})

	TryTestInput(t, Instruction{"RL", []string{"A"}}, []byte{0xCB, 0x17})
	TryTestInput(t, Instruction{"RL", []string{"B"}}, []byte{0xCB, 0x10})
	TryTestInput(t, Instruction{"RL", []string{"[HL]"}}, []byte{0xCB, 0x16})

	TryTestInput(t, Instruction{"RR", []string{"A"}}, []byte{0xCB, 0x1F})
	TryTestInput(t, Instruction{"RR", []string{"B"}}, []byte{0xCB, 0x18})
	TryTestInput(t, Instruction{"RR", []string{"[HL]"}}, []byte{0xCB, 0x1E})

	TryTestInput(t, Instruction{"SLA", []string{"A"}}, []byte{0xCB, 0x27})
	TryTestInput(t, Instruction{"SLA", []string{"B"}}, []byte{0xCB, 0x20})
	TryTestInput(t, Instruction{"SLA", []string{"[HL]"}}, []byte{0xCB, 0x26})

	TryTestInput(t, Instruction{"SRA", []string{"A"}}, []byte{0xCB, 0x2F})
	TryTestInput(t, Instruction{"SRA", []string{"B"}}, []byte{0xCB, 0x28})
	TryTestInput(t, Instruction{"SRA", []string{"[HL]"}}, []byte{0xCB, 0x2E})

	TryTestInput(t, Instruction{"SWAP", []string{"A"}}, []byte{0xCB, 0x37})
	TryTestInput(t, Instruction{"SWAP", []string{"B"}}, []byte{0xCB, 0x30})
	TryTestInput(t, Instruction{"SWAP", []string{"[HL]"}}, []byte{0xCB, 0x36})

	TryTestInput(t, Instruction{"SRL", []string{"A"}}, []byte{0xCB, 0x3F})
	TryTestInput(t, Instruction{"SRL", []string{"B"}}, []byte{0xCB, 0x38})
	TryTestInput(t, Instruction{"SRL", []string{"[HL]"}}, []byte{0xCB, 0x3E})
}

func TestMiscALUInstructions(t *testing.T) {
	TryTestInput(t, Instruction{"RLCA", []string{}}, []byte{0x07})
	TryTestInput(t, Instruction{"RRCA", []string{}}, []byte{0x0F})
	TryTestInput(t, Instruction{"RLA", []string{}}, []byte{0x17})
	TryTestInput(t, Instruction{"RRA", []string{}}, []byte{0x1F})
	TryTestInput(t, Instruction{"DAA", []string{}}, []byte{0x27})
	TryTestInput(t, Instruction{"CPL", []string{}}, []byte{0x2F})
	TryTestInput(t, Instruction{"SCF", []string{}}, []byte{0x37})
	TryTestInput(t, Instruction{"CCF", []string{}}, []byte{0x3F})
}

func TestExchangeInstructions(t *testing.T) {
	TryTestInput(t, Instruction{"EX", []string{"AF", "AF'"}}, []byte{0x08})
	TryTestInput(t, Instruction{"EX", []string{"[SP]", "HL"}}, []byte{0xE3})
	TryTestInput(t, Instruction{"EX", []string{"DE", "HL"}}, []byte{0xEB})

	TryTestInput(t, Instruction{"EXX", []string{}}, []byte{0xD9})
}

func TestDataInstructions(t *testing.T) {
	TryTestInput(t, Instruction{"IN", []string{"[C]"}}, []byte{0xED, 0x70})
	TryTestInput(t, Instruction{"IN", []string{"C", "[C]"}}, []byte{0xED, 0x48})
	TryTestInput(t, Instruction{"IN", []string{"A", "[66]"}}, []byte{0xDB, 0x42})

	TryTestInput(t, Instruction{"OUT", []string{"[C]", "B"}}, []byte{0xED, 0x41})
	TryTestInput(t, Instruction{"OUT", []string{"[C]", "0"}}, []byte{0xED, 0x71})
	TryTestInput(t, Instruction{"OUT", []string{"[66]", "A"}}, []byte{0xD3, 0x42})
}

func TestBlockInstructions(t *testing.T) {
	TryTestInput(t, Instruction{"LDI", []string{}}, []byte{0xED, 0xA0})
	TryTestInput(t, Instruction{"CPI", []string{}}, []byte{0xED, 0xA1})
	TryTestInput(t, Instruction{"INI", []string{}}, []byte{0xED, 0xA2})
	TryTestInput(t, Instruction{"OUTI", []string{}}, []byte{0xED, 0xA3})

	TryTestInput(t, Instruction{"LDD", []string{}}, []byte{0xED, 0xA8})
	TryTestInput(t, Instruction{"CPD", []string{}}, []byte{0xED, 0xA9})
	TryTestInput(t, Instruction{"IND", []string{}}, []byte{0xED, 0xAA})
	TryTestInput(t, Instruction{"OUTD", []string{}}, []byte{0xED, 0xAB})

	TryTestInput(t, Instruction{"LDIR", []string{}}, []byte{0xED, 0xB0})
	TryTestInput(t, Instruction{"CPIR", []string{}}, []byte{0xED, 0xB1})
	TryTestInput(t, Instruction{"INIR", []string{}}, []byte{0xED, 0xB2})
	TryTestInput(t, Instruction{"OTIR", []string{}}, []byte{0xED, 0xB3})

	TryTestInput(t, Instruction{"LDDR", []string{}}, []byte{0xED, 0xB8})
	TryTestInput(t, Instruction{"CPDR", []string{}}, []byte{0xED, 0xB9})
	TryTestInput(t, Instruction{"INDR", []string{}}, []byte{0xED, 0xBA})
	TryTestInput(t, Instruction{"OTDR", []string{}}, []byte{0xED, 0xBB})
}

func TestMiscInstructions(t *testing.T) {
	TryTestInput(t, Instruction{"DB", []string{"66"}}, []byte{0x42})
	TryTestInput(t, Instruction{"DB", []string{"66", "66"}}, []byte{0x42, 0x42})
	TryTestInput(t, Instruction{"DB", []string{"66", "66", "66"}}, []byte{0x42, 0x42, 0x42})
	TryTestInput(t, Instruction{"DW", []string{"1234"}}, []byte{0xD2, 0x04})

	TryTestInput(t, Instruction{"ASCII", []string{"\"hello\""}}, []byte{0x68, 0x65, 0x6C, 0x6C, 0x6F})
	TryTestInput(t, Instruction{"ASCIZ", []string{"\"hello\""}}, []byte{0x68, 0x65, 0x6C, 0x6C, 0x6F, 0x00})

	TryTestInput(t, Instruction{"DI", []string{}}, []byte{0xF3})
	TryTestInput(t, Instruction{"EI", []string{}}, []byte{0xFB})
	TryTestInput(t, Instruction{"HALT", []string{}}, []byte{0x76})
	TryTestInput(t, Instruction{"NOP", []string{}}, []byte{0x00})
	TryTestInput(t, Instruction{"PUSH", []string{"BC"}}, []byte{0xC5})
	TryTestInput(t, Instruction{"PUSH", []string{"DE"}}, []byte{0xD5})
	TryTestInput(t, Instruction{"PUSH", []string{"HL"}}, []byte{0xE5})
	TryTestInput(t, Instruction{"PUSH", []string{"AF"}}, []byte{0xF5})
	TryTestInput(t, Instruction{"POP", []string{"BC"}}, []byte{0xC1})
	TryTestInput(t, Instruction{"POP", []string{"DE"}}, []byte{0xD1})
	TryTestInput(t, Instruction{"POP", []string{"HL"}}, []byte{0xE1})
	TryTestInput(t, Instruction{"POP", []string{"AF"}}, []byte{0xF1})
}
