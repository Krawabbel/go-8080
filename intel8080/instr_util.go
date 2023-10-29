package intel8080

func push_stack(i8080 *Intel8080, addr word) {
	hi, lo := Split(addr)
	i8080.writeByte(i8080.sp-1, hi)
	i8080.writeByte(i8080.sp-2, lo)
	i8080.sp -= 2
}

func pop_stack(i8080 *Intel8080) word {
	lo := i8080.readByte(i8080.sp)
	hi := i8080.readByte(i8080.sp + 1)
	i8080.sp += 2
	return Join(hi, lo)
}

func add(i8080 *Intel8080, data byte, cy byte) {
	result := i8080.a + data + cy
	i8080.setZSP(result)
	i8080.f[FLAG_CARRY], i8080.f[FLAG_HALF_CARRY] = carry(i8080.a, data, cy)
	i8080.a = result
}

func sub(i8080 *Intel8080, data byte, cy byte) {
	add(i8080, ^data, cy^1)
	i8080.f[FLAG_CARRY] = !i8080.f[FLAG_CARRY]
}

func dad(i8080 *Intel8080, add word) {
	hl := i8080.getHL()
	i8080.f[FLAG_CARRY] = 0xFFFF-add < hl
	i8080.setHL(hl + add)
}

func inr(i8080 *Intel8080, val *byte) {
	res := *val + 1
	i8080.setZSP(res)
	i8080.f[FLAG_HALF_CARRY] = (*val & 0xF) == 0xF
	*val = res
}

func dcr(i8080 *Intel8080, val *byte) {
	res := *val - 1
	i8080.setZSP(res)
	i8080.f[FLAG_HALF_CARRY] = !((res & 0xF) == 0xF)
	*val = res
}

func ana(i8080 *Intel8080, val byte) {
	res := i8080.a & val
	i8080.setZSP(res)
	i8080.f[FLAG_HALF_CARRY] = (i8080.a|val)&0x08 > 0
	i8080.f[FLAG_CARRY] = false
	i8080.a = res
}

func xra(i8080 *Intel8080, val byte) {
	i8080.a = i8080.a ^ val
	i8080.setZSP(i8080.a)
	i8080.f[FLAG_HALF_CARRY] = false
	i8080.f[FLAG_CARRY] = false
}

func ora(i8080 *Intel8080, val byte) {
	i8080.a = i8080.a | val
	i8080.setZSP(i8080.a)
	i8080.f[FLAG_HALF_CARRY] = false
	i8080.f[FLAG_CARRY] = false
}

func cmp(i8080 *Intel8080, val byte) {
	a := i8080.a
	sub(i8080, val, 0)
	i8080.a = a
}

func jmp(i8080 *Intel8080, addr word) {
	i8080.pc = addr
}

func cond_jmp(i8080 *Intel8080, cond bool) {
	addr := i8080.nextWord()
	if cond {
		jmp(i8080, addr)
	}
}

func call(i8080 *Intel8080, addr word) {
	push_stack(i8080, i8080.pc)
	jmp(i8080, addr)
}

func cond_call(i8080 *Intel8080, cond bool) {
	addr := i8080.nextWord()
	if cond {
		call(i8080, addr)
		i8080.clock += 6
	}
}

func ret(i8080 *Intel8080) {
	i8080.pc = pop_stack(i8080)
}

func cond_ret(i8080 *Intel8080, cond bool) {
	if cond {
		ret(i8080)
		i8080.clock += 6
	}
}

func rlc(i8080 *Intel8080) {
	cy := (i8080.a >> 7)
	i8080.f[FLAG_CARRY] = cy > 0
	i8080.a = (i8080.a << 1) | cy
}

func rrc(i8080 *Intel8080) {
	cy := i8080.a & 1
	i8080.f[FLAG_CARRY] = cy > 0
	i8080.a = (i8080.a >> 1) | (cy << 7)
}

func ral(i8080 *Intel8080) {
	cy := i8080.valFC()
	i8080.f[FLAG_CARRY] = (i8080.a >> 7) > 0
	i8080.a = (i8080.a << 1) | (cy << 7)
}

func rar(i8080 *Intel8080) {
	cy := i8080.valFC()
	i8080.f[FLAG_CARRY] = (i8080.a & 1) > 0
	i8080.a = (i8080.a >> 1) | (cy << 7)
}

func daa(i8080 *Intel8080) {

	if i8080.a&0x0F > 0x09 || i8080.f[FLAG_HALF_CARRY] {
		i8080.f[FLAG_HALF_CARRY] = (i8080.a & 0x0F) > (0x0F - 0x06)
		i8080.a += 0x06
	} else {
		i8080.f[FLAG_HALF_CARRY] = false
	}

	if (i8080.a&0xF0) > 0x90 || i8080.f[FLAG_CARRY] {
		i8080.f[FLAG_CARRY] = (i8080.a & 0xF0) > (0xF0 - 0x60)
		i8080.a += 0x60
	} else {
		i8080.f[FLAG_CARRY] = false
	}

	i8080.setZSP(i8080.a)
}

func xchg(i8080 *Intel8080) {
	i8080.h, i8080.l, i8080.d, i8080.e = i8080.d, i8080.e, i8080.h, i8080.l
}

func xthl(i8080 *Intel8080) {
	l := i8080.l
	h := i8080.h

	i8080.l = i8080.readByte(i8080.sp)
	i8080.h = i8080.readByte(i8080.sp + 1)

	i8080.writeByte(i8080.sp, l)
	i8080.writeByte(i8080.sp+1, h)
}
