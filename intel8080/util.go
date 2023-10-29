package intel8080

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type unsigned interface {
	uint8 | uint16
}

func hex[T unsigned](value T) string {
	maxT := ^T(0)
	n := len(fmt.Sprintf("%X", maxT))
	format := "0x%0" + fmt.Sprint(n) + "X"
	return fmt.Sprintf(format, value)
}

func hexs[T unsigned](values []T) string {
	s := make([]string, len(values))
	for i, value := range values {
		s[i] = hex(value)
	}
	return "[" + strings.Join(s, " ") + "]"
}

func Join(hi, lo byte) word {
	return word(lo) | word(hi)<<8
}

func Split(addr word) (hi, lo byte) {
	return byte(addr >> 8), byte(addr)
}

func carry(a, b, cy byte) (carry, half_carry bool) {
	sum := word(a) + word(b) + word(cy)
	carries := sum ^ word(a) ^ word(b)
	return carries&(1<<8) > 0, carries&(1<<4) > 0
}

func parity(val byte) bool {
	val = val ^ (val >> 4)
	val = val ^ (val >> 2)
	val = val ^ (val >> 1)
	return (val & 1) == 0
}

func Load(path string) ([]byte, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("load file error: %s", err)
	}
	defer f.Close()

	prog, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("load file error: %s", err)
	}

	return prog, nil
}
