package mobs

import "math/rand"

type Mob struct {
	Nome   string
	Hp     int
	Atk    int
	Def    int
	Nvl    int
	XPBase int
}

func (m Mob) Atacar(def int) (int, bool) {
	critico := rand.Intn(100)
	if critico <= 5 {
		return (m.Atk / ((def + 1) / 2)) * 2, true
	}
	return (m.Atk / ((def + 1) / 2)), false
}

func (p *Mob) Defender(dano int) (int, bool) {
	sorte := rand.Intn(100)
	if sorte <= 10 {
		return 0, true
	}
	return (dano / 2) + 1, false
}
