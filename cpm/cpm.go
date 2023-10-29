package cpm

import (
	"fmt"
	"time"

	"github.com/Krawabbel/go-8080/intel8080"
)

const WARM_BOOT_MSG = "RE-ENTRY TO CP/M WARM BOOT"

func Run(path string) error {

	prog, err := intel8080.Load(path)
	if err != nil {
		return err
	}

	orig := word(0x100)

	memory := make([]byte, 0x10000)
	memory[0x0005] = 0xC9
	copy(memory[orig:], prog)

	bus := Bus(memory)

	i8080 := intel8080.NewIntel8080(&bus, orig)

	tStart := time.Now()
	for steps := 0; true; steps++ {

		i8080.Step()

		if err := call_bdos(i8080); err != nil {
			switch err.Error() {
			case WARM_BOOT_MSG:
				duration := time.Since(tStart)
				fmt.Printf("\nduration: %v\n", duration)
				fmt.Printf("CPU clock: %v cycles\n", i8080.Clock())
				fmt.Print(i8080.DebugSpeed(duration))
				return nil
			default:
				return err
			}
		}
	}
	return nil
}
