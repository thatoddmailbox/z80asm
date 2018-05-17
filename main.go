package main

import (
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
)

const KiB = 1024

var WeirdMapping bool

func main() {
	log.Println("z80asm")

	flag.BoolVar(&WeirdMapping, "weird-mapping", false, "Enables the weird mapping, with two modern ROMs in ROM0 and ROM1, and a modern RAM chip in ROM3.")

	flag.Parse()

	workingDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ReadConfigFile(workingDirectory)

	ROM_Create(workingDirectory)

	log.Println()
	log.Println("Constant listing:")

	definitionKeys := make([]string, len(CurrentROM.Definitions))
	i := 0
	for name, _ := range CurrentROM.Definitions {
		definitionKeys[i] = name
		i += 1
	}

	sort.Strings(definitionKeys)

	for _, name := range definitionKeys {
		value := CurrentROM.Definitions[name]
		log.Println(" *", name, value, "0x"+strconv.FormatInt(int64(value), 16))
	}

	log.Println()
	log.Println("Base addresses:")
	for i := 0; i < ROM_BankCount; i++ {
		log.Printf(" * Bank %d: 0x%x", i, ROM_CalculateAbsoluteAddress(0, i))
	}

	log.Println()
	log.Println("Usage:")
	for i := 0; i < ROM_BankCount; i++ {
		log.Printf(" * Bank %d: %d out of 2048 bytes", i, CurrentROM.UsedByteCount[i])
	}
}
