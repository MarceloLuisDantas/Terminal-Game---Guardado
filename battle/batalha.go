package battle

import (
	"terminal/mobs"
	"terminal/player"
	"terminal/txt"
)

func Batalha(p *player.Player, m *mobs.Mob) {
	for {
		txt.CallClear()
		txt.StatusBatalha(p, m)

	}
}
