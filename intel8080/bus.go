package intel8080

type Bus interface {
	Read(word) byte
	Write(word, byte)
}
