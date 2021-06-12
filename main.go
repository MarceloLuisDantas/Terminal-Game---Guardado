package main

import (
	"terminal/battle"
	"terminal/input"
	"terminal/mobs"
	"terminal/player"
	"terminal/txt"
)

func main() {
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
		battle.Batalha(&jogador, &morcego)
		_ = input.StrInput("")
	}
}
