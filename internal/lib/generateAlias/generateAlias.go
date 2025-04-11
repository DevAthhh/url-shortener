package generateAlias

import (
	"math/rand"
	"strings"
)

func GenerateStr(size int) string {
	symbols := "QWERTYUIOPASDFGHJKLZXCVBNM" +
		"qwertyuiopasdfghjklzxcvbnm" +
		"0123456789"

	sign := strings.Builder{}

	for i := 0; i < size; i++ {
		idx := rand.Intn(len(symbols))
		sign.WriteByte(symbols[idx])
	}

	return sign.String()
}
