package intel8080

import (
	"fmt"
	"log"
)

const (
	FLAG_CARRY = iota
	set_1
	FLAG_PARITY
	unset_2
	FLAG_HALF_CARRY
	unset_3
	FLAG_ZERO
	FLAG_SIGN
)

const (
	REG_A = iota
	REG_B
	REG_C
	REG_D
	REG_E
	REG_F
	REG_H
	REG_L
	REG_BC
	REG_DE
	REG_AF
	REG_HL
	REG_PC
	REG_SP
)

type word = uint16

type Intel8080 struct {
	sp, pc              word    // stack pointer, program counter
	a, b, c, d, e, h, l byte    // registers
	f                   [8]bool // zero/parity/sign/carry/half-carry flags
	iff                 bool    // interrupt flip-flop

	clock uint64 // cpu clock (in cycles)

	halted bool // has program been halted?

	interrupt_pending bool // is an interrupt pending?
	interrupt_vector  byte // command to be executed after interrupt
	interrupt_delay   byte // interrupts are only serviced 1 instruction after calling "DI"

	bus Bus // ROM, RAM, VRAM, etc.

	InputDevice  [0x100]func() byte // input ports
	OutputDevice [0x100]func(byte)  // output ports

}

func NewIntel8080(bus Bus, pc word) *Intel8080 {
	i8080 := new(Intel8080)
	for i := 0; i <= 0xFF; i++ {
		port := byte(i)
		i8080.OutputDevice[port] = func(b byte) {
			log.Fatalf("unexpected data sent to port %v of output device: %v\n", hex(port), hex(b))
		}
		i8080.InputDevice[port] = func() byte {
			log.Fatalf("unexpected input requested on port %v of input device\n", hex(port))
			return 0
		}
	}
	i8080.bus = bus
	i8080.pc = pc

	return i8080
}

func (i8080 *Intel8080) Step() {
	if i8080.interrupt_pending && i8080.iff && i8080.interrupt_delay == 0 {
		i8080.interrupt_pending = false
		i8080.iff = false
		i8080.halted = false
		i8080.clock += execute(i8080, i8080.interrupt_vector)
	} else if !i8080.halted {
		i8080.clock += execute(i8080, i8080.nextByte())
	}
}

func (i8080 *Intel8080) Interrupt(code byte) {
	i8080.interrupt_pending = true
	i8080.interrupt_vector = code
}

func (i8080 Intel8080) Clock() uint64 {
	return i8080.clock
}

func (i8080 Intel8080) ReadBus(addr word) byte {
	return i8080.bus.Read(addr)
}

func (i8080 Intel8080) Register8(reg int) (byte, error) {
	switch reg {
	case REG_A:
		return i8080.a, nil
	case REG_B:
		return i8080.b, nil
	case REG_C:
		return i8080.c, nil
	case REG_D:
		return i8080.d, nil
	case REG_E:
		return i8080.e, nil
	case REG_F:
		return i8080.getF(), nil
	case REG_H:
		return i8080.h, nil
	case REG_L:
		return i8080.l, nil
	default:
		return 0, fmt.Errorf("unknown 8-bit register %d", reg)
	}
}

func (i8080 Intel8080) Register16(reg_pair int) (word, error) {
	switch reg_pair {
	case REG_AF:
		return i8080.getAF(), nil
	case REG_BC:
		return i8080.getBC(), nil
	case REG_DE:
		return i8080.getDE(), nil
	case REG_HL:
		return i8080.getHL(), nil
	case REG_PC:
		return i8080.pc, nil
	case REG_SP:
		return i8080.sp, nil
	default:
		return 0, fmt.Errorf("unknown 16-bit register %d", reg_pair)
	}
}

func (i8080 Intel8080) Flag(flag int) (bool, error) {
	switch flag {
	case FLAG_CARRY, FLAG_HALF_CARRY, FLAG_PARITY, FLAG_SIGN, FLAG_ZERO:
		return i8080.f[flag], nil
	default:
		return false, fmt.Errorf("unknown flag %d", flag)
	}
}
