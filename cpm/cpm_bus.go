package cpm

type Bus []byte

func (bus *Bus) Write(addr word, val byte) {
	(*bus)[addr] = val
}

func (mem Bus) Read(addr word) byte {
	return mem[addr]
}
