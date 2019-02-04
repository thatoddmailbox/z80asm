# z80asm
A simple Z80 assembler, based heavily on [gbasm](https://github.com/thatoddmailbox/gbasm). Many things, such as the special instructions (`asciz`, `db`, `.incasm`, etc.) are shared between the two, and you should look at gbasm's documentation for more, specifically the [known issues](https://github.com/thatoddmailbox/gbasm#known-issues) and [assembler-specific instructions](https://github.com/thatoddmailbox/gbasm#assembler-instructions). This assembler was written for use with the [Soviet-era microcomputer I built](https://github.com/thatoddmailbox/computer), and is probably missing a few features or instructions. Notable omissions include the ability to create local labels and the relative jump instruction.

If you use Sublime Text, you can use the `gbz80.sublime-syntax` file in the repository for syntax highlighting, but as it was originally intended for use with the Gameboy's LR35902 (a Z80 derivative), it's missing a few instructions.

## Usage
You will need Go installed and set up properly to build the assembler.
```shell
go get https://github.com/thatoddmailbox/z80asm
cd ~/go/src/github.com/thatoddmailbox/z80asm # you might need to change this depending on the location of your GOPATH
go install
# the next command should be run from the directory that contains your source code
~/go/bin/z80asm --weird-mapping
```

The `--weird-mapping` flag enables the modified address decoding, which was necessary to adapt modern-day ROM and RAM chips to the computer when the Soviet parts didn't work.