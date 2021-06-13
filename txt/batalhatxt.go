package txt

import (
	"fmt"
	"terminal/player"
)

// Mostra o texto de ganho de XP e se o jogador upou
func FinalBatalha(xp int, upou bool, p, old player.Player) {
	fmt.Println("Você ganhou", xp, "pontos de experiencia")
	if upou {
		fmt.Printf("Você upou %d -> %d \n", old.Nvl, p.Nvl)
		fmt.Printf("Hp  %d -> %d \n", old.MHp, p.MHp)
		fmt.Printf("Atk %d -> %d \n", old.Atk, p.Atk)
		fmt.Printf("Def %d -> %d \n", old.Def, p.Def)
		fmt.Printf("Xp  %d -> %d \n", old.Next, p.Next)
	}
}
