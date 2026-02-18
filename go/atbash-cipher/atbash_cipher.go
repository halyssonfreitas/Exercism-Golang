package atbash

import (
	"strings"
)

func Atbash(s string) string {
	s = strings.ToLower(s)
	var builder strings.Builder
	count := 0 // Contador para controlar o agrupamento de 5 em 5

	for _, r := range s {
		var charToAppend rune

		if r >= 'a' && r <= 'z' {
			// Regra 1: Espelhar letras
			charToAppend = 'z' - (r - 'a')
		} else if r >= '0' && r <= '9' {
			// Regra 2: Manter números
			charToAppend = r
		} else {
			// Regra 3: Ignorar espaços, pontos, vírgulas, etc.
			continue
		}

		// Lógica de agrupamento (adicionar espaço a cada 5 caracteres)
		if count > 0 && count%5 == 0 {
			builder.WriteByte(' ')
		}

		builder.WriteRune(charToAppend)
		count++
	}

	return builder.String()
}
