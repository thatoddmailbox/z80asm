package main

import (
	"os"
	"path"
)

type ROM struct {
	Info                 ROMInfo
	Output               [8 * KiB]byte
	Definitions          map[string]int
	UnpointedDefinitions []string
}

var CurrentROM ROM

func ROM_Create(basePath string) {
	CurrentROM.Definitions = map[string]int{}

	outputFileName := path.Join(basePath, "out.bin")

	// output the actual data
	Assembler_ParseFile(path.Join(basePath, "main.s"), 0x0000, 8*KiB)

	// output the actual file
	outputFile, err := os.OpenFile(outputFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	_, err = outputFile.Write(CurrentROM.Output[:])
	if err != nil {
		panic(err)
	}
}
