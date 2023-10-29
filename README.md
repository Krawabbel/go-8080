# go-8080
an Intel 8080 emulator written in Go (Golang)

(also contains a minimal implementation of CP/M to run CPU tests)

* It passes (almost) all CPU tests that I could find. Run 
```
go run . /path/to/cputest/file.COM
```
to see test output. (Warning: 8080EXM.COM takes quite long to finish and tells me I still have a bug somewhere.)

* It can be used to emulate the Space Invaders Arcade Cabinet.

## Resources

* opcode table by Pastraiser: https://pastraiser.com/cpu/i8080/i8080_opcodes.html
* opcode table by emulator101: http://www.emulator101.com/8080-by-opcode.html
* Intel 8080 Wikipedia page: https://en.wikipedia.org/wiki/Intel_8080#Programming_model
* EmuDev Tutorial q00.si: https://emudev.de/q00-si/a-short-fun-project/
* CPU tests: https://altairclone.com/downloads/cpu_tests/
* VM tests: https://altairclone.com/downloads/
* Intel 8080 Assembly Programmer Manual: https://altairclone.com/downloads/manuals/8080%20Programmers%20Manual.pdf
* 8080/8085 Assembly Language Programming Manual: http://bitsavers.org/components/intel/MCS80/9800301D_8080_8085_Assembly_Language_Programming_Manual_May81.pdf
* Introduction to CPM: https://obsolescence.wixsite.com/obsolescence/introduction-to-cpm
* Comprehensive CP/M reference: https://www.seasip.info/Cpm/bdos.html
* Intel 8080 emulation tutorial: http://www.emulator101.com/
* Intel 8080 data sheet: https://deramp.com/downloads/intel/8080%20Data%20Sheet.pdf

## Other Implementations
* https://github.com/superzazu/8080/blob/master/i8080.c
* https://github.com/tobiasvl/lua-8080/


