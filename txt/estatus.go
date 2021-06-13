// FUNÇÕES PARA MOSTRAR OS ESTATUS DO JOGADOR E INIMIGOS

package txt

import (
	"fmt"
	"strconv"
	"strings"
	"terminal/mobs"
	"terminal/player"
)

// Mostra os estatus do jogador fora de batalha
func Status(p *player.Player) {
	fmt.Println("Nome: ", p.Nome)
	fmt.Println("HP: ", p.Hp)
	fmt.Println("ATK: ", p.Atk, "DEF: ", p.Def)
}

// Mostra os estatus do jogador e do inimigo em batalha
func StatusBatalha(p *player.Player, m *mobs.Mob) {
	// Calcula a quantidade espaço para encaixar o nome
	espacoNome := strings.Repeat(" ", 50-6-len(p.Nome))

	// Calcula o espaço para enciaxar o HP
	tHp := len(strconv.Itoa(p.Hp))
	espacoHp := strings.Repeat(" ", 50-4-tHp)

	// Calcula o espaço para encaixar o ATK e DEF
	tAtk := len(strconv.Itoa(p.Atk))
	tDef := len(strconv.Itoa(p.Def))
	espacoAtk := strings.Repeat(" ", 50-13-tAtk+tDef)

	fmt.Printf("Nome: %s", p.Nome)
	fmt.Printf("%s", espacoNome)
	fmt.Printf("Nome: %s \n", m.Nome)

	fmt.Printf("HP: %d", p.Hp)
	fmt.Printf("%s", espacoHp)
	fmt.Printf("HP: %d \n", m.Hp)

	fmt.Printf("ATK: %d DEF: %d", p.Atk, p.Def)
	fmt.Printf("%s", espacoAtk)
	fmt.Printf("ATK: %d DEF: %d \n", m.Atk, m.Def)

	print("\n")
}
