package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func StrInput(label string) string {
	r := bufio.NewReader(os.Stdin)
	fmt.Printf(label)
	v, _, err := r.ReadLine()
	if err != nil {
		panic(err)
	}
	return string(v)
}

func IntInput(label string) (int, error) {
	entrada := StrInput(label)
	valor, err := strconv.Atoi(entrada)
	if err != nil {
		return 0, err
	}
	return valor, nil
}
