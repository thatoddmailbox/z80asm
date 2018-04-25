package main

import (
	"log"
	"os"
	"strconv"
)

const KiB = 1024

func main() {
	log.Println("z80asm")

	workingDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ReadConfigFile(workingDirectory)

	ROM_Create(workingDirectory)

	log.Println()
	log.Println("Constant listing:")
	for name, val := range CurrentROM.Definitions {
		log.Println(" *", name, val, "0x"+strconv.FormatInt(int64(val), 16))
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
