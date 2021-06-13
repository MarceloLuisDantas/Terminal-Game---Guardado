// PACOTE RESPOSNAVEL POR GERENCIAR AS ENTRADAS DO JOGADOR

package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Recebe uma string do jogador
func StrInput(label string) string {
	r := bufio.NewReader(os.Stdin)
	fmt.Print(label)
	v, _, err := r.ReadLine()
	if err != nil {
		panic(err)
	}
	return string(v)
}

// Recebe um inteiro do jogador
func IntInput(label string) (int, error) {
	entrada := StrInput(label)
	valor, err := strconv.Atoi(entrada)
	if err != nil {
		return 0, err
	}
	return valor, nil
}
