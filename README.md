# z80asm
A simple Z80 assembler, based heavily on [gbasm](https://github.com/thatoddmailbox/gbasm). Many things, such as the special instructions (`asciz`, `db`, `.incasm`, etc.) are shared between the two, and you should look at gbasm's documentation for more. This assembler was written for use with the Soviet-era microcomputer I built. See [this repository](https://github.com/thatoddmailbox/computer) for more information.

## Usage
You will need Go installed and set up properly to build the emulator.
```shell
go get https://github.com/thatoddmailbox/z80asm
cd ~/go/src/github.com/thatoddmailbox/z80asm # you might need to change this depending on the location of your GOPATH
go install
# the next command should be run from the directory that contains your source code
~/go/bin/z80asm --weird-mapping
```

The `--weird-mapping` flag enables the modified address decoding, which was necessary to adapt modern-day ROM and RAM chips to the computer when the Soviet parts were found to be defective.