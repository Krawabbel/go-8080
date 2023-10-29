package intel8080

import (
	"fmt"
	"time"
)

func (i8080 Intel8080) DebugState() string {
	return fmt.Sprintf(
		"PC: %s, AF: %s, BC: %s, DE: %s, HL: %s, SP: %s, CYC: %v",
		hex(i8080.pc),
		hex(i8080.getAF()),
		hex(i8080.getBC()),
		hex(i8080.getDE()),
		hex(i8080.getHL()),
		hex(i8080.sp),
		i8080.clock)
}

func (i8080 Intel8080) DebugProg() string {
	next := make([]byte, 5)
	for i := range next {
		next[i] = i8080.readByte(i8080.pc + word(i))
	}
	return fmt.Sprintf("%s: %-15s >> %s", hex(i8080.pc), hexs(next), mnemonics[i8080.readByte(i8080.pc)])
}

func (i8080 Intel8080) DebugSpeed(duration time.Duration) string {
	freq := float64(i8080.clock) / float64(duration.Microseconds())
	speedup := freq / 2
	return fmt.Sprintf("clock rate: %0.2f MHz (%0.2fx original Intel 8080 frequency)\n", freq, speedup)
}
