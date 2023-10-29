package intel8080

func execute(i8080 *Intel8080, opcode byte) uint64 {

	if i8080.interrupt_delay > 0 {
		i8080.interrupt_delay--
	}

	switch opcode {
	case 0x00, 0x10, 0x20, 0x30, 0x08, 0x18, 0x28, 0x38:
		return 4

	case 0x01:
		i8080.setBC(i8080.nextWord())
		return 10
	case 0x11:
		i8080.setDE(i8080.nextWord())
		return 10
	case 0x21:
		i8080.setHL(i8080.nextWord())
		return 10
	case 0x31:
		i8080.sp = i8080.nextWord()
		return 10

	case 0x02:
		i8080.writeByte(i8080.getBC(), i8080.a)
		return 7
	case 0x12:
		i8080.writeByte(i8080.getDE(), i8080.a)
		return 7
	case 0x22:
		addr := i8080.nextWord()
		i8080.writeByte(addr, i8080.l)
		i8080.writeByte(addr+1, i8080.h)
		return 16
	case 0x32:
		i8080.writeByte(i8080.nextWord(), i8080.a)
		return 13

	case 0x03:
		i8080.setBC(i8080.getBC() + 1)
		return 5
	case 0x13:
		i8080.setDE(i8080.getDE() + 1)
		return 5
	case 0x23:
		i8080.setHL(i8080.getHL() + 1)
		return 5
	case 0x33:
		i8080.sp++
		return 5

	case 0x04:
		inr(i8080, &i8080.b)
		return 5
	case 0x14:
		inr(i8080, &i8080.d)
		return 5
	case 0x24:
		inr(i8080, &i8080.h)
		return 5
	case 0x34:
		m := i8080.getM()
		inr(i8080, &m)
		i8080.setM(m)
		return 10

	case 0x05:
		dcr(i8080, &i8080.b)
		return 5
	case 0x15:
		dcr(i8080, &i8080.d)
		return 5
	case 0x25:
		dcr(i8080, &i8080.h)
		return 5
	case 0x35:
		m := i8080.getM()
		dcr(i8080, &m)
		i8080.setM(m)
		return 10

	case 0x06:
		i8080.b = i8080.nextByte()
		return 7
	case 0x16:
		i8080.d = i8080.nextByte()
		return 7
	case 0x26:
		i8080.h = i8080.nextByte()
		return 7
	case 0x36:
		i8080.setM(i8080.nextByte())
		return 10

	case 0x07:
		rlc(i8080)
		return 4
	case 0x17:
		ral(i8080)
		return 4

	case 0x27:
		daa(i8080)
		return 4

	case 0x37:
		i8080.f[FLAG_CARRY] = true
		return 4

	case 0x09:
		dad(i8080, i8080.getBC())
		return 10
	case 0x19:
		dad(i8080, i8080.getDE())
		return 10
	case 0x29:
		dad(i8080, i8080.getHL())
		return 10
	case 0x39:
		dad(i8080, i8080.sp)
		return 10

	case 0x0A:
		i8080.a = i8080.readByte(i8080.getBC())
		return 7
	case 0x1A:
		i8080.a = i8080.readByte(i8080.getDE())
		return 7

	case 0x2A:
		addr := i8080.nextWord()
		i8080.l = i8080.readByte(addr)
		i8080.h = i8080.readByte(addr + 1)
		return 16

	case 0x3A:
		i8080.a = i8080.readByte(i8080.nextWord())
		return 13

	case 0x0B:
		i8080.setBC(i8080.getBC() - 1)
		return 5
	case 0x1B:
		i8080.setDE(i8080.getDE() - 1)
		return 5
	case 0x2B:
		i8080.setHL(i8080.getHL() - 1)
		return 5
	case 0x3B:
		i8080.sp--
		return 5

	case 0x0C:
		inr(i8080, &i8080.c)
		return 5
	case 0x1C:
		inr(i8080, &i8080.e)
		return 5
	case 0x2C:
		inr(i8080, &i8080.l)
		return 5
	case 0x3C:
		inr(i8080, &i8080.a)
		return 5

	case 0x0D:
		dcr(i8080, &i8080.c)
		return 5
	case 0x1D:
		dcr(i8080, &i8080.e)
		return 5
	case 0x2D:
		dcr(i8080, &i8080.l)
		return 5
	case 0x3D:
		dcr(i8080, &i8080.a)
		return 5

	case 0x0E:
		i8080.c = i8080.nextByte()
		return 7
	case 0x1E:
		i8080.e = i8080.nextByte()
		return 7
	case 0x2E:
		i8080.l = i8080.nextByte()
		return 7
	case 0x3E:
		i8080.a = i8080.nextByte()
		return 7

	case 0x0F:
		rrc(i8080)
		return 4
	case 0x1F:
		rar(i8080)
		return 4

	case 0x2F:
		i8080.a = ^i8080.a
		return 4
	case 0x3F:
		i8080.f[FLAG_CARRY] = !i8080.f[FLAG_CARRY]
		return 4

	case 0x40:
		// i8080.b = i8080.b
		return 5
	case 0x41:
		i8080.b = i8080.c
		return 5
	case 0x42:
		i8080.b = i8080.d
		return 5
	case 0x43:
		i8080.b = i8080.e
		return 5
	case 0x44:
		i8080.b = i8080.h
		return 5
	case 0x45:
		i8080.b = i8080.l
		return 5
	case 0x46:
		i8080.b = i8080.getM()
		return 7
	case 0x47:
		i8080.b = i8080.a
		return 5
	case 0x50:
		i8080.d = i8080.b
		return 5
	case 0x51:
		i8080.d = i8080.c
		return 5
	case 0x52:
		// i8080.d = i8080.d
		return 5
	case 0x53:
		i8080.d = i8080.e
		return 5
	case 0x54:
		i8080.d = i8080.h
		return 5
	case 0x55:
		i8080.d = i8080.l
		return 5
	case 0x56:
		i8080.d = i8080.getM()
		return 7
	case 0x57:
		i8080.d = i8080.a
		return 5
	case 0x60:
		i8080.h = i8080.b
		return 5
	case 0x61:
		i8080.h = i8080.c
		return 5
	case 0x62:
		i8080.h = i8080.d
		return 5
	case 0x63:
		i8080.h = i8080.e
		return 5
	case 0x64:
		// i8080.h = i8080.h
		return 5
	case 0x65:
		i8080.h = i8080.l
		return 5
	case 0x66:
		i8080.h = i8080.getM()
		return 7
	case 0x67:
		i8080.h = i8080.a
		return 5
	case 0x70:
		i8080.setM(i8080.b)
		return 7
	case 0x71:
		i8080.setM(i8080.c)
		return 7
	case 0x72:
		i8080.setM(i8080.d)
		return 7
	case 0x73:
		i8080.setM(i8080.e)
		return 7
	case 0x74:
		i8080.setM(i8080.h)
		return 7
	case 0x75:
		i8080.setM(i8080.l)
		return 7
	case 0x76:
		i8080.halted = true
		return 7
	case 0x77:
		i8080.setM(i8080.a)
		return 7
	case 0x48:
		i8080.c = i8080.b
		return 5
	case 0x49:
		// i8080.c = i8080.c
		return 5
	case 0x4A:
		i8080.c = i8080.d
		return 5
	case 0x4B:
		i8080.c = i8080.e
		return 5
	case 0x4C:
		i8080.c = i8080.h
		return 5
	case 0x4D:
		i8080.c = i8080.l
		return 5
	case 0x4E:
		i8080.c = i8080.getM()
		return 7
	case 0x4F:
		i8080.c = i8080.a
		return 5
	case 0x58:
		i8080.e = i8080.b
		return 5
	case 0x59:
		i8080.e = i8080.c
		return 5
	case 0x5A:
		i8080.e = i8080.d
		return 5
	case 0x5B:
		// i8080.e = i8080.e
		return 5
	case 0x5C:
		i8080.e = i8080.h
		return 5
	case 0x5D:
		i8080.e = i8080.l
		return 5
	case 0x5E:
		i8080.e = i8080.getM()
		return 7
	case 0x5F:
		i8080.e = i8080.a
		return 5
	case 0x68:
		i8080.l = i8080.b
		return 5
	case 0x69:
		i8080.l = i8080.c
		return 5
	case 0x6A:
		i8080.l = i8080.d
		return 5
	case 0x6B:
		i8080.l = i8080.e
		return 5
	case 0x6C:
		i8080.l = i8080.h
		return 5
	case 0x6D:
		// i8080.l = i8080.l
		return 5
	case 0x6E:
		i8080.l = i8080.getM()
		return 7
	case 0x6F:
		i8080.l = i8080.a
		return 5
	case 0x78:
		i8080.a = i8080.b
		return 5
	case 0x79:
		i8080.a = i8080.c
		return 5
	case 0x7A:
		i8080.a = i8080.d
		return 5
	case 0x7B:
		i8080.a = i8080.e
		return 5
	case 0x7C:
		i8080.a = i8080.h
		return 5
	case 0x7D:
		i8080.a = i8080.l
		return 5
	case 0x7E:
		i8080.a = i8080.getM()
		return 7
	case 0x7F:
		// i8080.a = i8080.a
		return 5

	case 0x80:
		add(i8080, i8080.b, 0)
		return 4
	case 0x81:
		add(i8080, i8080.c, 0)
		return 4
	case 0x82:
		add(i8080, i8080.d, 0)
		return 4
	case 0x83:
		add(i8080, i8080.e, 0)
		return 4
	case 0x84:
		add(i8080, i8080.h, 0)
		return 4
	case 0x85:
		add(i8080, i8080.l, 0)
		return 4
	case 0x86:
		add(i8080, i8080.getM(), 0)
		return 7
	case 0x87:
		add(i8080, i8080.a, 0)
		return 4

	case 0x88:
		add(i8080, i8080.b, i8080.valFC())
		return 4
	case 0x89:
		add(i8080, i8080.c, i8080.valFC())
		return 4
	case 0x8A:
		add(i8080, i8080.d, i8080.valFC())
		return 4
	case 0x8B:
		add(i8080, i8080.e, i8080.valFC())
		return 4
	case 0x8C:
		add(i8080, i8080.h, i8080.valFC())
		return 4
	case 0x8D:
		add(i8080, i8080.l, i8080.valFC())
		return 4
	case 0x8E:
		add(i8080, i8080.getM(), i8080.valFC())
		return 7 //, "ADC M", nil
	case 0x8F:
		add(i8080, i8080.a, i8080.valFC())
		return 4

	case 0x90:
		sub(i8080, i8080.b, 0)
		return 4
	case 0x91:
		sub(i8080, i8080.c, 0)
		return 4
	case 0x92:
		sub(i8080, i8080.d, 0)
		return 4
	case 0x93:
		sub(i8080, i8080.e, 0)
		return 4
	case 0x94:
		sub(i8080, i8080.h, 0)
		return 4
	case 0x95:
		sub(i8080, i8080.l, 0)
		return 4
	case 0x96:
		sub(i8080, i8080.getM(), 0)
		return 7
	case 0x97:
		sub(i8080, i8080.a, 0)
		return 4

	case 0x98:
		sub(i8080, i8080.b, i8080.valFC())
		return 4
	case 0x99:
		sub(i8080, i8080.c, i8080.valFC())
		return 4
	case 0x9A:
		sub(i8080, i8080.d, i8080.valFC())
		return 4
	case 0x9B:
		sub(i8080, i8080.e, i8080.valFC())
		return 4
	case 0x9C:
		sub(i8080, i8080.h, i8080.valFC())
		return 4
	case 0x9D:
		sub(i8080, i8080.l, i8080.valFC())
		return 4
	case 0x9E:
		sub(i8080, i8080.getM(), i8080.valFC())
		return 7
	case 0x9F:
		sub(i8080, i8080.a, i8080.valFC())
		return 4

	case 0xA0:
		ana(i8080, i8080.b)
		return 4
	case 0xA1:
		ana(i8080, i8080.c)
		return 4
	case 0xA2:
		ana(i8080, i8080.d)
		return 4
	case 0xA3:
		ana(i8080, i8080.e)
		return 4
	case 0xA4:
		ana(i8080, i8080.h)
		return 4
	case 0xA5:
		ana(i8080, i8080.l)
		return 4
	case 0xA6:
		ana(i8080, i8080.getM())
		return 7
	case 0xA7:
		ana(i8080, i8080.a)
		return 4

	case 0xA8:
		xra(i8080, i8080.b)
		return 4
	case 0xA9:
		xra(i8080, i8080.c)
		return 4
	case 0xAA:
		xra(i8080, i8080.d)
		return 4
	case 0xAB:
		xra(i8080, i8080.e)
		return 4
	case 0xAC:
		xra(i8080, i8080.h)
		return 4
	case 0xAD:
		xra(i8080, i8080.l)
		return 4
	case 0xAE:
		xra(i8080, i8080.getM())
		return 7
	case 0xAF:
		xra(i8080, i8080.a)
		return 4

	case 0xB0:
		ora(i8080, i8080.b)
		return 4
	case 0xB1:
		ora(i8080, i8080.c)
		return 4
	case 0xB2:
		ora(i8080, i8080.d)
		return 4
	case 0xB3:
		ora(i8080, i8080.e)
		return 4
	case 0xB4:
		ora(i8080, i8080.h)
		return 4
	case 0xB5:
		ora(i8080, i8080.l)
		return 4
	case 0xB6:
		ora(i8080, i8080.getM())
		return 7
	case 0xB7:
		ora(i8080, i8080.a)
		return 4

	case 0xB8:
		cmp(i8080, i8080.b)
		return 4
	case 0xB9:
		cmp(i8080, i8080.c)
		return 4
	case 0xBA:
		cmp(i8080, i8080.d)
		return 4
	case 0xBB:
		cmp(i8080, i8080.e)
		return 4
	case 0xBC:
		cmp(i8080, i8080.h)
		return 4
	case 0xBD:
		cmp(i8080, i8080.l)
		return 4
	case 0xBE:
		cmp(i8080, i8080.getM())
		return 7
	case 0xBF:
		cmp(i8080, i8080.a)
		return 4

	case 0xC0:
		cond_ret(i8080, !i8080.f[FLAG_ZERO])
		return 5
	case 0xD0:
		cond_ret(i8080, !i8080.f[FLAG_CARRY])
		return 5
	case 0xE0:
		cond_ret(i8080, !i8080.f[FLAG_PARITY])
		return 5
	case 0xF0:
		cond_ret(i8080, !i8080.f[FLAG_SIGN])
		return 5

	case 0xC1:
		i8080.setBC(pop_stack(i8080))
		return 10
	case 0xD1:
		i8080.setDE(pop_stack(i8080))
		return 10
	case 0xE1:
		i8080.setHL(pop_stack(i8080))
		return 10
	case 0xF1:
		i8080.setPSW(pop_stack(i8080))
		return 10

	case 0xC2:
		cond_jmp(i8080, !i8080.f[FLAG_ZERO])
		return 10
	case 0xD2:
		cond_jmp(i8080, !i8080.f[FLAG_CARRY])
		return 10
	case 0xE2:
		cond_jmp(i8080, !i8080.f[FLAG_PARITY])
		return 10
	case 0xF2:
		cond_jmp(i8080, !i8080.f[FLAG_SIGN])
		return 10

	case 0xC3, 0xCB:
		jmp(i8080, i8080.nextWord())
		return 10

	case 0xD3:
		i8080.OutputDevice[i8080.nextByte()](i8080.a)
		return 10

	case 0xE3:
		xthl(i8080)
		return 18

	case 0xF3:
		i8080.iff = false
		return 4

	case 0xC4:
		cond_call(i8080, !i8080.f[FLAG_ZERO])
		return 11
	case 0xD4:
		cond_call(i8080, !i8080.f[FLAG_CARRY])
		return 11
	case 0xE4:
		cond_call(i8080, !i8080.f[FLAG_PARITY])
		return 11
	case 0xF4:
		cond_call(i8080, !i8080.f[FLAG_SIGN])
		return 11

	case 0xC5:
		push_stack(i8080, i8080.getBC())
		return 11
	case 0xD5:
		push_stack(i8080, i8080.getDE())
		return 11
	case 0xE5:
		push_stack(i8080, i8080.getHL())
		return 11
	case 0xF5:
		push_stack(i8080, i8080.getPSW())
		return 11

	case 0xC6:
		add(i8080, i8080.nextByte(), 0)
		return 7

	case 0xD6:
		sub(i8080, i8080.nextByte(), 0)
		return 7

	case 0xE6:
		ana(i8080, i8080.nextByte())
		return 7

	case 0xF6:
		ora(i8080, i8080.nextByte())
		return 7

	case 0xC7:
		call(i8080, 0x00)
		return 11
	case 0xD7:
		call(i8080, 0x10)
		return 11
	case 0xE7:
		call(i8080, 0x20)
		return 11
	case 0xF7:
		call(i8080, 0x30)
		return 11

	case 0xC8:
		cond_ret(i8080, i8080.f[FLAG_ZERO])
		return 5
	case 0xD8:
		cond_ret(i8080, i8080.f[FLAG_CARRY])
		return 5
	case 0xE8:
		cond_ret(i8080, i8080.f[FLAG_PARITY])
		return 5
	case 0xF8:
		cond_ret(i8080, i8080.f[FLAG_SIGN])
		return 5

	case 0xC9, 0xD9:
		ret(i8080)
		return 10

	case 0xE9:
		i8080.pc = Join(i8080.h, i8080.l)
		return 5

	case 0xF9:
		i8080.sp = i8080.getHL()
		return 5

	case 0xCA:
		cond_jmp(i8080, i8080.f[FLAG_ZERO])
		return 10
	case 0xDA:
		cond_jmp(i8080, i8080.f[FLAG_CARRY])
		return 10
	case 0xEA:
		cond_jmp(i8080, i8080.f[FLAG_PARITY])
		return 10
	case 0xFA:
		cond_jmp(i8080, i8080.f[FLAG_SIGN])
		return 10

	case 0xDB:
		i8080.a = i8080.InputDevice[i8080.nextByte()]()
		return 10

	case 0xEB:
		xchg(i8080)
		return 5

	case 0xFB:
		i8080.iff = true
		i8080.interrupt_delay = 1
		return 4

	case 0xCC:
		cond_call(i8080, i8080.f[FLAG_ZERO])
		return 11
	case 0xDC:
		cond_call(i8080, i8080.f[FLAG_CARRY])
		return 11
	case 0xEC:
		cond_call(i8080, i8080.f[FLAG_PARITY])
		return 11
	case 0xFC:
		cond_call(i8080, i8080.f[FLAG_SIGN])
		return 11

	case 0xCD, 0xDD, 0xED, 0xFD:
		call(i8080, i8080.nextWord())
		return 11

	case 0xCE:
		add(i8080, i8080.nextByte(), i8080.valFC())
		return 7

	case 0xDE:
		sub(i8080, i8080.nextByte(), i8080.valFC())
		return 7

	case 0xEE:
		xra(i8080, i8080.nextByte())
		return 7

	case 0xFE:
		cmp(i8080, i8080.nextByte())
		return 7

	case 0xCF:
		call(i8080, 0x08)
		return 11
	case 0xDF:
		call(i8080, 0x18)
		return 11
	case 0xEF:
		call(i8080, 0x28)
		return 11
	case 0xFF:
		call(i8080, 0x38)
		return 11

	}

	panic("unexpected operation: " + hex(opcode))

}
