package main

import (
	"os"
	"path"
	"strconv"
)

const ROM_BankCount = 4

type ROM struct {
	Info                 ROMInfo
	CurrentBank          int
	Output               [ROM_BankCount][2 * KiB]byte
	UsedByteCount        [ROM_BankCount]int
	Definitions          map[string]int
	UnpointedDefinitions []string
}

var CurrentROM ROM

func ROM_Create(basePath string) {
	CurrentROM.Definitions = map[string]int{}

	// output the actual data
	Assembler_ParseFile(path.Join(basePath, "main.s"), 0x0000, 8*KiB)

	// output the actual file
	if !WeirdMapping {
		for i := 0; i < ROM_BankCount; i++ {
			outputFileName := path.Join(basePath, "bank"+strconv.Itoa(i)+".bin")
			outputFile, err := os.OpenFile(outputFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
			if err != nil {
				panic(err)
			}
			defer outputFile.Close()
			_, err = outputFile.Write(CurrentROM.Output[i][:])
			if err != nil {
				panic(err)
			}
		}
	} else {
		// output the patched files
		var rom0 [128 * 1024]byte
		var rom1 [32 * 1024]byte

		copy(rom0[0x2800:], CurrentROM.Output[0][:])
		copy(rom0[0x3800:], CurrentROM.Output[1][:])
		copy(rom1[0x2800:], CurrentROM.Output[2][:])
		copy(rom1[0x3800:], CurrentROM.Output[3][:])

		rom0File, err := os.OpenFile(path.Join(basePath, "rom0.bin"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			panic(err)
		}
		defer rom0File.Close()
		_, err = rom0File.Write(rom0[:])
		if err != nil {
			panic(err)
		}

		rom1File, err := os.OpenFile(path.Join(basePath, "rom1.bin"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			panic(err)
		}
		defer rom1File.Close()
		_, err = rom1File.Write(rom1[:])
		if err != nil {
			panic(err)
		}
	}
}

func ROM_CalculateAbsoluteAddress(address uint16, bank int) uint16 {
	var base uint16
	if !WeirdMapping {
		if bank == 0 {
			base = 0x0000
		} else if bank == 1 {
			base = 0x1000
		} else if bank == 2 {
			base = 0x2000
		} else if bank == 3 {
			base = 0x3000
		}
	} else {
		if bank == 0 {
			base = 0x0000
		} else if bank == 1 {
			base = 0x0800
		} else if bank == 2 {
			base = 0x1000
		} else if bank == 3 {
			base = 0x1800
		}
	}
	return base + address
}
