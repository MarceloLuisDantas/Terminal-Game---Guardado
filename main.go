package main

import (
	"fmt"
	"strings"
	"terminal/battle"
	"terminal/input"
	"terminal/mobs"
	"terminal/player"
	"terminal/txt"
)

func main() {
	txt.CallClear()
	morcego := mobs.Mob{
		Nome:   "Morcego",
		Hp:     3,
		Atk:    1,
		Def:    1,
		Nvl:    1,
		XPBase: 2,
	}
	txt.Top()
	jogador := player.NewPlayer()
	for {
		txt.CallClear()
		txt.Top2()
		resultado := battle.Batalha(&jogador, &morcego)
		if resultado == -1 {
			txt.CallClear()
			fmt.Println("Você perdeu")
			break
		} else if resultado == 0 {
			txt.CallClear()
			fmt.Println("Batalha empatada")
			break
		} else {
			txt.CallClear()
			fmt.Println("Você venceu")
		}
		cont := input.StrInput("Deseja continua? [s/n]: ")
		if strings.ToLower(cont) != "s" {
			break
		}
	}
}
