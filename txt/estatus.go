package txt

import (
	"fmt"
	"strconv"
	"strings"
	"terminal/mobs"
	"terminal/player"
)

func Status(p *player.Player) {
	fmt.Println("Nome: ", p.Nome)
	fmt.Println("HP: ", p.Hp)
	fmt.Println("ATK: ", p.Atk, "DEF: ", p.Def)
}

func StatusBatalha(p *player.Player, m *mobs.Mob) {
	espacoNome := strings.Repeat(" ", 50-6-len(p.Nome))

	tHp := len(strconv.Itoa(p.Hp))
	espacoHp := strings.Repeat(" ", 50-4-tHp)

	tAtk := len(strconv.Itoa(p.Atk))
	tDef := len(strconv.Itoa(p.Def))
	espacoAtk := strings.Repeat(" ", 50-13-tAtk+tDef)

	fmt.Printf("Nome: %s", p.Nome)
	fmt.Printf("%s", espacoNome)
	fmt.Printf("Nome: %s \n", m.Nome)

	fmt.Printf("HP: %d", p.Hp)
	fmt.Printf("%s", espacoHp)
	fmt.Printf("HP: %d \n", p.Hp)

	fmt.Printf("ATK: %d DEF: %d", p.Atk, p.Def)
	fmt.Printf("%s", espacoAtk)
	fmt.Printf("ATK: %d DEF: %d \n", m.Atk, m.Def)

	print("\n")
}
