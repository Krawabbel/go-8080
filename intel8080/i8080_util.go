package intel8080

func (i8080 Intel8080) valFC() byte {
	if i8080.f[FLAG_CARRY] {
		return 1
	}
	return 0
}

func (i8080 Intel8080) getF() byte {
	F := byte(0)
	for i := 0; i < 8; i++ {
		switch {
		case i == unset_2, i == unset_3:
			// do nothing
		case i == set_1:
			F |= (1 << i)
		case i8080.f[i]:
			F |= (1 << i)
		}
	}
	return F
}

func (i8080 *Intel8080) setF(F byte) {
	for i := 0; i < 8; i++ {
		i8080.f[i] = F&(1<<i) > 0
	}
	i8080.f[set_1] = true
	i8080.f[unset_2] = false
	i8080.f[unset_3] = false
}

func (i8080 *Intel8080) setZSP(val byte) {
	i8080.f[FLAG_ZERO] = val == 0
	i8080.f[FLAG_SIGN] = val&(1<<7) > 0
	i8080.f[FLAG_PARITY] = parity(val)
}

func (i8080 Intel8080) getAF() word {
	return Join(i8080.a, i8080.getF())
}

func (i8080 *Intel8080) setAF(af word) {
	a, f := Split(af)
	i8080.a = a
	i8080.setF(f)
}

func (i8080 Intel8080) getPSW() word {
	return i8080.getAF()
}

func (i8080 *Intel8080) setPSW(psw word) {
	i8080.setAF(psw)
}

func (i8080 Intel8080) getBC() word {
	return Join(i8080.b, i8080.c)
}

func (i8080 *Intel8080) setBC(bc word) {
	i8080.b, i8080.c = Split(bc)
}

func (i8080 Intel8080) getDE() word {
	return Join(i8080.d, i8080.e)
}

func (i8080 *Intel8080) setDE(de word) {
	i8080.d, i8080.e = Split(de)
}

func (i8080 Intel8080) getHL() word {
	return Join(i8080.h, i8080.l)
}

func (i8080 *Intel8080) setHL(hl word) {
	i8080.h, i8080.l = Split(hl)
}

func (i8080 Intel8080) readByte(addr word) byte {
	return i8080.bus.Read(addr)
}

func (i8080 *Intel8080) writeByte(addr word, val byte) {
	i8080.bus.Write(addr, val)
}

func (i8080 Intel8080) getM() byte {
	addr := Join(i8080.h, i8080.l)
	return i8080.readByte(addr)
}

func (i8080 *Intel8080) setM(val byte) {
	addr := Join(i8080.h, i8080.l)
	i8080.writeByte(addr, val)
}

func (i8080 *Intel8080) nextByte() byte {
	val := i8080.readByte(i8080.pc)
	i8080.pc++
	return val
}
func (i8080 *Intel8080) nextWord() word {
	lo := i8080.nextByte()
	hi := i8080.nextByte()
	return Join(hi, lo)
}
