package cpm

import (
	"fmt"

	"github.com/Krawabbel/go-8080/intel8080"
)

type word = uint16

func call_bdos(i8080 *intel8080.Intel8080) error {

	pc, _ := i8080.Register16(intel8080.REG_PC)
	switch pc {
	case 0x0005:
		return write(i8080)
	case 0x0000:
		return fmt.Errorf(WARM_BOOT_MSG)
	default:
		if pc < 0x100 {
			return fmt.Errorf("unexpected BDOS access at address 0x%04X", pc)
		}
	}
	return nil
}

func write(i8080 *intel8080.Intel8080) error {
	reg_c, _ := i8080.Register8(intel8080.REG_C)
	switch reg_c {
	case 0x09:
		output := ""
		do_continue := true
		addr, _ := i8080.Register16(intel8080.REG_DE)
		for ; do_continue; addr++ {
			if c := i8080.ReadBus(addr); c != '$' {
				output += string(c)
			} else {
				do_continue = false
			}
		}
		fmt.Print(output)

	case 0x02, 0x05:
		reg_e, _ := i8080.Register8(intel8080.REG_E)
		fmt.Print(string(rune(reg_e)))
	default:
		return fmt.Errorf("unexpected BDOS function %v called", reg_c)
	}

	return nil
}
