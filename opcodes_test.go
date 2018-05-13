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
	TryTestInput(t, Instruction{"CALL", []string{"NZ", "1234"}}, []byte{0xC4, 0xD2, 0x04})
	TryTestInput(t, Instruction{"CALL", []string{"Z", "1234"}}, []byte{0xCC, 0xD2, 0x04})
	TryTestInput(t, Instruction{"CALL", []string{"NC", "1234"}}, []byte{0xD4, 0xD2, 0x04})
	TryTestInput(t, Instruction{"CALL", []string{"C", "1234"}}, []byte{0xDC, 0xD2, 0x04})
	TryTestInput(t, Instruction{"CALL", []string{"PO", "1234"}}, []byte{0xE4, 0xD2, 0x04})
	TryTestInput(t, Instruction{"CALL", []string{"PE", "1234"}}, []byte{0xEC, 0xD2, 0x04})
	TryTestInput(t, Instruction{"CALL", []string{"P", "1234"}}, []byte{0xF4, 0xD2, 0x04})
	TryTestInput(t, Instruction{"CALL", []string{"M", "1234"}}, []byte{0xFC, 0xD2, 0x04})

	TryTestInput(t, Instruction{"JP", []string{"1234"}}, []byte{0xC3, 0xD2, 0x04})
	TryTestInput(t, Instruction{"JP", []string{"HL"}}, []byte{0xE9})
	TryTestInput(t, Instruction{"JP", []string{"[HL]"}}, []byte{0xE9})
	TryTestInput(t, Instruction{"JP", []string{"NZ", "1234"}}, []byte{0xC2, 0xD2, 0x04})
	TryTestInput(t, Instruction{"JP", []string{"Z", "1234"}}, []byte{0xCA, 0xD2, 0x04})
	TryTestInput(t, Instruction{"JP", []string{"NC", "1234"}}, []byte{0xD2, 0xD2, 0x04})
	TryTestInput(t, Instruction{"JP", []string{"C", "1234"}}, []byte{0xDA, 0xD2, 0x04})
	TryTestInput(t, Instruction{"JP", []string{"PO", "1234"}}, []byte{0xE2, 0xD2, 0x04})
	TryTestInput(t, Instruction{"JP", []string{"PE", "1234"}}, []byte{0xEA, 0xD2, 0x04})
	TryTestInput(t, Instruction{"JP", []string{"P", "1234"}}, []byte{0xF2, 0xD2, 0x04})
	TryTestInput(t, Instruction{"JP", []string{"M", "1234"}}, []byte{0xFA, 0xD2, 0x04})

	TryTestInput(t, Instruction{"RET", []string{}}, []byte{0xC9})
	TryTestInput(t, Instruction{"RET", []string{"NZ"}}, []byte{0xC0})
	TryTestInput(t, Instruction{"RET", []string{"Z"}}, []byte{0xC8})
	TryTestInput(t, Instruction{"RET", []string{"NC"}}, []byte{0xD0})
	TryTestInput(t, Instruction{"RET", []string{"C"}}, []byte{0xD8})
	TryTestInput(t, Instruction{"RET", []string{"PO"}}, []byte{0xE0})
	TryTestInput(t, Instruction{"RET", []string{"PE"}}, []byte{0xE8})
	TryTestInput(t, Instruction{"RET", []string{"P"}}, []byte{0xF0})
	TryTestInput(t, Instruction{"RET", []string{"M"}}, []byte{0xF8})
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
	TryTestInput(t, Instruction{"LD", []string{"A", "A"}}, []byte{0x7F})
	TryTestInput(t, Instruction{"LD", []string{"A", "B"}}, []byte{0x78})
	TryTestInput(t, Instruction{"LD", []string{"A", "C"}}, []byte{0x79})
	TryTestInput(t, Instruction{"LD", []string{"A", "D"}}, []byte{0x7A})
	TryTestInput(t, Instruction{"LD", []string{"A", "E"}}, []byte{0x7B})
	TryTestInput(t, Instruction{"LD", []string{"A", "H"}}, []byte{0x7C})
	TryTestInput(t, Instruction{"LD", []string{"A", "L"}}, []byte{0x7D})

	TryTestInput(t, Instruction{"LD", []string{"B", "A"}}, []byte{0x47})
	TryTestInput(t, Instruction{"LD", []string{"B", "B"}}, []byte{0x40})
	TryTestInput(t, Instruction{"LD", []string{"B", "C"}}, []byte{0x41})
	TryTestInput(t, Instruction{"LD", []string{"B", "D"}}, []byte{0x42})
	TryTestInput(t, Instruction{"LD", []string{"B", "E"}}, []byte{0x43})
	TryTestInput(t, Instruction{"LD", []string{"B", "H"}}, []byte{0x44})
	TryTestInput(t, Instruction{"LD", []string{"B", "L"}}, []byte{0x45})

	TryTestInput(t, Instruction{"LD", []string{"C", "A"}}, []byte{0x4F})
	TryTestInput(t, Instruction{"LD", []string{"C", "B"}}, []byte{0x48})
	TryTestInput(t, Instruction{"LD", []string{"C", "C"}}, []byte{0x49})
	TryTestInput(t, Instruction{"LD", []string{"C", "D"}}, []byte{0x4A})
	TryTestInput(t, Instruction{"LD", []string{"C", "E"}}, []byte{0x4B})
	TryTestInput(t, Instruction{"LD", []string{"C", "H"}}, []byte{0x4C})
	TryTestInput(t, Instruction{"LD", []string{"C", "L"}}, []byte{0x4D})

	TryTestInput(t, Instruction{"LD", []string{"D", "A"}}, []byte{0x57})
	TryTestInput(t, Instruction{"LD", []string{"D", "B"}}, []byte{0x50})
	TryTestInput(t, Instruction{"LD", []string{"D", "C"}}, []byte{0x51})
	TryTestInput(t, Instruction{"LD", []string{"D", "D"}}, []byte{0x52})
	TryTestInput(t, Instruction{"LD", []string{"D", "E"}}, []byte{0x53})
	TryTestInput(t, Instruction{"LD", []string{"D", "H"}}, []byte{0x54})
	TryTestInput(t, Instruction{"LD", []string{"D", "L"}}, []byte{0x55})

	TryTestInput(t, Instruction{"LD", []string{"E", "A"}}, []byte{0x5F})
	TryTestInput(t, Instruction{"LD", []string{"E", "B"}}, []byte{0x58})
	TryTestInput(t, Instruction{"LD", []string{"E", "C"}}, []byte{0x59})
	TryTestInput(t, Instruction{"LD", []string{"E", "D"}}, []byte{0x5A})
	TryTestInput(t, Instruction{"LD", []string{"E", "E"}}, []byte{0x5B})
	TryTestInput(t, Instruction{"LD", []string{"E", "H"}}, []byte{0x5C})
	TryTestInput(t, Instruction{"LD", []string{"E", "L"}}, []byte{0x5D})

	TryTestInput(t, Instruction{"LD", []string{"H", "A"}}, []byte{0x67})
	TryTestInput(t, Instruction{"LD", []string{"H", "B"}}, []byte{0x60})
	TryTestInput(t, Instruction{"LD", []string{"H", "C"}}, []byte{0x61})
	TryTestInput(t, Instruction{"LD", []string{"H", "D"}}, []byte{0x62})
	TryTestInput(t, Instruction{"LD", []string{"H", "E"}}, []byte{0x63})
	TryTestInput(t, Instruction{"LD", []string{"H", "H"}}, []byte{0x64})
	TryTestInput(t, Instruction{"LD", []string{"H", "L"}}, []byte{0x65})

	TryTestInput(t, Instruction{"LD", []string{"L", "A"}}, []byte{0x6F})
	TryTestInput(t, Instruction{"LD", []string{"L", "B"}}, []byte{0x68})
	TryTestInput(t, Instruction{"LD", []string{"L", "C"}}, []byte{0x69})
	TryTestInput(t, Instruction{"LD", []string{"L", "D"}}, []byte{0x6A})
	TryTestInput(t, Instruction{"LD", []string{"L", "E"}}, []byte{0x6B})
	TryTestInput(t, Instruction{"LD", []string{"L", "H"}}, []byte{0x6C})
	TryTestInput(t, Instruction{"LD", []string{"L", "L"}}, []byte{0x6D})

	TryTestInput(t, Instruction{"LD", []string{"A", "66"}}, []byte{0x3E, 0x42})
	TryTestInput(t, Instruction{"LD", []string{"B", "66"}}, []byte{0x06, 0x42})
	TryTestInput(t, Instruction{"LD", []string{"C", "66"}}, []byte{0x0E, 0x42})
	TryTestInput(t, Instruction{"LD", []string{"D", "66"}}, []byte{0x16, 0x42})
	TryTestInput(t, Instruction{"LD", []string{"E", "66"}}, []byte{0x1E, 0x42})
	TryTestInput(t, Instruction{"LD", []string{"H", "66"}}, []byte{0x26, 0x42})
	TryTestInput(t, Instruction{"LD", []string{"L", "66"}}, []byte{0x2E, 0x42})

	TryTestInput(t, Instruction{"LD", []string{"BC", "1234"}}, []byte{0x01, 0xD2, 0x04})
	TryTestInput(t, Instruction{"LD", []string{"DE", "1234"}}, []byte{0x11, 0xD2, 0x04})
	TryTestInput(t, Instruction{"LD", []string{"HL", "1234"}}, []byte{0x21, 0xD2, 0x04})
	TryTestInput(t, Instruction{"LD", []string{"SP", "1234"}}, []byte{0x31, 0xD2, 0x04})

	TryTestInput(t, Instruction{"LD", []string{"SP", "HL"}}, []byte{0xF9})

	TryTestInput(t, Instruction{"LD", []string{"[HL]", "66"}}, []byte{0x36, 0x42})

	TryTestInput(t, Instruction{"LD", []string{"[BC]", "A"}}, []byte{0x02})
	TryTestInput(t, Instruction{"LD", []string{"[DE]", "A"}}, []byte{0x12})

	TryTestInput(t, Instruction{"LD", []string{"[HL]", "A"}}, []byte{0x77})
	TryTestInput(t, Instruction{"LD", []string{"[HL]", "B"}}, []byte{0x70})
	TryTestInput(t, Instruction{"LD", []string{"[HL]", "C"}}, []byte{0x71})
	TryTestInput(t, Instruction{"LD", []string{"[HL]", "D"}}, []byte{0x72})
	TryTestInput(t, Instruction{"LD", []string{"[HL]", "E"}}, []byte{0x73})
	TryTestInput(t, Instruction{"LD", []string{"[HL]", "H"}}, []byte{0x74})
	TryTestInput(t, Instruction{"LD", []string{"[HL]", "L"}}, []byte{0x75})

	TryTestInput(t, Instruction{"LD", []string{"A", "[BC]"}}, []byte{0x0A})
	TryTestInput(t, Instruction{"LD", []string{"A", "[DE]"}}, []byte{0x1A})

	TryTestInput(t, Instruction{"LD", []string{"A", "[HL]"}}, []byte{0x7E})
	TryTestInput(t, Instruction{"LD", []string{"B", "[HL]"}}, []byte{0x46})
	TryTestInput(t, Instruction{"LD", []string{"C", "[HL]"}}, []byte{0x4E})
	TryTestInput(t, Instruction{"LD", []string{"D", "[HL]"}}, []byte{0x56})
	TryTestInput(t, Instruction{"LD", []string{"E", "[HL]"}}, []byte{0x5E})
	TryTestInput(t, Instruction{"LD", []string{"H", "[HL]"}}, []byte{0x66})
	TryTestInput(t, Instruction{"LD", []string{"L", "[HL]"}}, []byte{0x6E})

	TryTestInput(t, Instruction{"LD", []string{"[1234]", "A"}}, []byte{0x32, 0xD2, 0x04})
	TryTestInput(t, Instruction{"LD", []string{"A", "[1234]"}}, []byte{0x3A, 0xD2, 0x04})
	TryTestInput(t, Instruction{"LD", []string{"[1234]", "HL"}}, []byte{0x22, 0xD2, 0x04})
	TryTestInput(t, Instruction{"LD", []string{"HL", "[1234]"}}, []byte{0x2A, 0xD2, 0x04})
	TryTestInput(t, Instruction{"LD", []string{"[1234]", "BC"}}, []byte{0xED, 0x43, 0xD2, 0x04})
	TryTestInput(t, Instruction{"LD", []string{"BC", "[1234]"}}, []byte{0xED, 0x4B, 0xD2, 0x04})
}

func TestALUInstructions(t *testing.T) {
	// add
	TryTestInput(t, Instruction{"ADD", []string{"A", "66"}}, []byte{0xC6, 0x42})

	TryTestInput(t, Instruction{"ADD", []string{"A", "A"}}, []byte{0x87})
	TryTestInput(t, Instruction{"ADD", []string{"A", "B"}}, []byte{0x80})
	TryTestInput(t, Instruction{"ADD", []string{"A", "C"}}, []byte{0x81})
	TryTestInput(t, Instruction{"ADD", []string{"A", "D"}}, []byte{0x82})
	TryTestInput(t, Instruction{"ADD", []string{"A", "E"}}, []byte{0x83})
	TryTestInput(t, Instruction{"ADD", []string{"A", "H"}}, []byte{0x84})
	TryTestInput(t, Instruction{"ADD", []string{"A", "L"}}, []byte{0x85})

	TryTestInput(t, Instruction{"ADD", []string{"A", "[HL]"}}, []byte{0x86})

	TryTestInput(t, Instruction{"ADD", []string{"HL", "BC"}}, []byte{0x09})
	TryTestInput(t, Instruction{"ADD", []string{"HL", "DE"}}, []byte{0x19})
	TryTestInput(t, Instruction{"ADD", []string{"HL", "HL"}}, []byte{0x29})
	TryTestInput(t, Instruction{"ADD", []string{"HL", "SP"}}, []byte{0x39})

	// adc
	TryTestInput(t, Instruction{"ADC", []string{"A", "66"}}, []byte{0xCE, 0x42})

	TryTestInput(t, Instruction{"ADC", []string{"A", "A"}}, []byte{0x8F})
	TryTestInput(t, Instruction{"ADC", []string{"A", "B"}}, []byte{0x88})
	TryTestInput(t, Instruction{"ADC", []string{"A", "C"}}, []byte{0x89})
	TryTestInput(t, Instruction{"ADC", []string{"A", "D"}}, []byte{0x8A})
	TryTestInput(t, Instruction{"ADC", []string{"A", "E"}}, []byte{0x8B})
	TryTestInput(t, Instruction{"ADC", []string{"A", "H"}}, []byte{0x8C})
	TryTestInput(t, Instruction{"ADC", []string{"A", "L"}}, []byte{0x8D})

	TryTestInput(t, Instruction{"ADC", []string{"A", "[HL]"}}, []byte{0x8E})

	// sub
	TryTestInput(t, Instruction{"SUB", []string{"66"}}, []byte{0xD6, 0x42})

	TryTestInput(t, Instruction{"SUB", []string{"A"}}, []byte{0x97})
	TryTestInput(t, Instruction{"SUB", []string{"B"}}, []byte{0x90})
	TryTestInput(t, Instruction{"SUB", []string{"C"}}, []byte{0x91})
	TryTestInput(t, Instruction{"SUB", []string{"D"}}, []byte{0x92})
	TryTestInput(t, Instruction{"SUB", []string{"E"}}, []byte{0x93})
	TryTestInput(t, Instruction{"SUB", []string{"H"}}, []byte{0x94})
	TryTestInput(t, Instruction{"SUB", []string{"L"}}, []byte{0x95})

	TryTestInput(t, Instruction{"SUB", []string{"[HL]"}}, []byte{0x96})

	// sbc
	TryTestInput(t, Instruction{"SBC", []string{"A", "66"}}, []byte{0xDE, 0x42})

	TryTestInput(t, Instruction{"SBC", []string{"A", "A"}}, []byte{0x9F})
	TryTestInput(t, Instruction{"SBC", []string{"A", "B"}}, []byte{0x98})
	TryTestInput(t, Instruction{"SBC", []string{"A", "C"}}, []byte{0x99})
	TryTestInput(t, Instruction{"SBC", []string{"A", "D"}}, []byte{0x9A})
	TryTestInput(t, Instruction{"SBC", []string{"A", "E"}}, []byte{0x9B})
	TryTestInput(t, Instruction{"SBC", []string{"A", "H"}}, []byte{0x9C})
	TryTestInput(t, Instruction{"SBC", []string{"A", "L"}}, []byte{0x9D})

	TryTestInput(t, Instruction{"SBC", []string{"A", "[HL]"}}, []byte{0x9E})
	TryTestInput(t, Instruction{"SBC", []string{"HL", "BC"}}, []byte{0xED, 0x42})

	// and
	TryTestInput(t, Instruction{"AND", []string{"66"}}, []byte{0xE6, 0x42})

	TryTestInput(t, Instruction{"AND", []string{"A"}}, []byte{0xA7})
	TryTestInput(t, Instruction{"AND", []string{"B"}}, []byte{0xA0})
	TryTestInput(t, Instruction{"AND", []string{"C"}}, []byte{0xA1})
	TryTestInput(t, Instruction{"AND", []string{"D"}}, []byte{0xA2})
	TryTestInput(t, Instruction{"AND", []string{"E"}}, []byte{0xA3})
	TryTestInput(t, Instruction{"AND", []string{"H"}}, []byte{0xA4})
	TryTestInput(t, Instruction{"AND", []string{"L"}}, []byte{0xA5})

	TryTestInput(t, Instruction{"AND", []string{"[HL]"}}, []byte{0xA6})

	// xor
	TryTestInput(t, Instruction{"XOR", []string{"66"}}, []byte{0xEE, 0x42})

	TryTestInput(t, Instruction{"XOR", []string{"A"}}, []byte{0xAF})
	TryTestInput(t, Instruction{"XOR", []string{"B"}}, []byte{0xA8})
	TryTestInput(t, Instruction{"XOR", []string{"C"}}, []byte{0xA9})
	TryTestInput(t, Instruction{"XOR", []string{"D"}}, []byte{0xAA})
	TryTestInput(t, Instruction{"XOR", []string{"E"}}, []byte{0xAB})
	TryTestInput(t, Instruction{"XOR", []string{"H"}}, []byte{0xAC})
	TryTestInput(t, Instruction{"XOR", []string{"L"}}, []byte{0xAD})

	TryTestInput(t, Instruction{"XOR", []string{"[HL]"}}, []byte{0xAE})

	// or
	TryTestInput(t, Instruction{"OR", []string{"66"}}, []byte{0xF6, 0x42})

	TryTestInput(t, Instruction{"OR", []string{"A"}}, []byte{0xB7})
	TryTestInput(t, Instruction{"OR", []string{"B"}}, []byte{0xB0})
	TryTestInput(t, Instruction{"OR", []string{"C"}}, []byte{0xB1})
	TryTestInput(t, Instruction{"OR", []string{"D"}}, []byte{0xB2})
	TryTestInput(t, Instruction{"OR", []string{"E"}}, []byte{0xB3})
	TryTestInput(t, Instruction{"OR", []string{"H"}}, []byte{0xB4})
	TryTestInput(t, Instruction{"OR", []string{"L"}}, []byte{0xB5})

	TryTestInput(t, Instruction{"OR", []string{"[HL]"}}, []byte{0xB6})

	// cp
	TryTestInput(t, Instruction{"CP", []string{"66"}}, []byte{0xFE, 0x42})

	TryTestInput(t, Instruction{"CP", []string{"A"}}, []byte{0xBF})
	TryTestInput(t, Instruction{"CP", []string{"B"}}, []byte{0xB8})
	TryTestInput(t, Instruction{"CP", []string{"C"}}, []byte{0xB9})
	TryTestInput(t, Instruction{"CP", []string{"D"}}, []byte{0xBA})
	TryTestInput(t, Instruction{"CP", []string{"E"}}, []byte{0xBB})
	TryTestInput(t, Instruction{"CP", []string{"H"}}, []byte{0xBC})
	TryTestInput(t, Instruction{"CP", []string{"L"}}, []byte{0xBD})

	TryTestInput(t, Instruction{"CP", []string{"[HL]"}}, []byte{0xBE})
}

func TestMathInstructions(t *testing.T) {
	TryTestInput(t, Instruction{"DEC", []string{"A"}}, []byte{0x3D})
	TryTestInput(t, Instruction{"DEC", []string{"B"}}, []byte{0x05})
	TryTestInput(t, Instruction{"DEC", []string{"C"}}, []byte{0x0D})
	TryTestInput(t, Instruction{"DEC", []string{"D"}}, []byte{0x15})
	TryTestInput(t, Instruction{"DEC", []string{"E"}}, []byte{0x1D})
	TryTestInput(t, Instruction{"DEC", []string{"H"}}, []byte{0x25})
	TryTestInput(t, Instruction{"DEC", []string{"L"}}, []byte{0x2D})

	TryTestInput(t, Instruction{"DEC", []string{"BC"}}, []byte{0x0B})
	TryTestInput(t, Instruction{"DEC", []string{"DE"}}, []byte{0x1B})
	TryTestInput(t, Instruction{"DEC", []string{"HL"}}, []byte{0x2B})
	TryTestInput(t, Instruction{"DEC", []string{"SP"}}, []byte{0x3B})

	TryTestInput(t, Instruction{"DEC", []string{"[HL]"}}, []byte{0x35})

	TryTestInput(t, Instruction{"INC", []string{"A"}}, []byte{0x3C})
	TryTestInput(t, Instruction{"INC", []string{"B"}}, []byte{0x04})
	TryTestInput(t, Instruction{"INC", []string{"C"}}, []byte{0x0C})
	TryTestInput(t, Instruction{"INC", []string{"D"}}, []byte{0x14})
	TryTestInput(t, Instruction{"INC", []string{"E"}}, []byte{0x1C})
	TryTestInput(t, Instruction{"INC", []string{"H"}}, []byte{0x24})
	TryTestInput(t, Instruction{"INC", []string{"L"}}, []byte{0x2C})

	TryTestInput(t, Instruction{"INC", []string{"BC"}}, []byte{0x03})
	TryTestInput(t, Instruction{"INC", []string{"DE"}}, []byte{0x13})
	TryTestInput(t, Instruction{"INC", []string{"HL"}}, []byte{0x23})
	TryTestInput(t, Instruction{"INC", []string{"SP"}}, []byte{0x33})

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
