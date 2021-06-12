package player

import (
	"fmt"
	"strings"
	"terminal/input"
)

type Player struct {
	Nome string  // Nome
	Hp   int     // vida
	Atk  int     // atauqe
	Def  int     // Defesa
	Nvl  int     // Nivel
	Xp   float64 // Total de XP
	Next int     // Total de XP para o proximo nivel
}

func calculaXp(pn, ml, mxp int) float64 {
	return float64(mxp*ml) + (float64(pn/2) + 1.5)
}

func (p *Player) Upar(ml, mxp int) {
	p.Xp = p.Xp + calculaXp(p.Nvl, ml, mxp)
	if p.Xp >= float64(p.Next) {
		p.Nvl++
		p.Xp = 0
		p.Next = p.Next * 2
		p.Hp = p.Hp + p.Hp/2
		p.Atk = p.Atk + p.Atk/2
		p.Def = p.Def + p.Def/2
	}
}

func NewPlayer() Player {
	var nome string
	for {
		nome = input.StrInput("Digite seu nome: ")
		ok := input.StrInput(fmt.Sprintf("Tem certeza que %s sera seu nome? [S/N]", nome))
		if strings.ToLower(ok) != "n" {
			break
		}
	}
	return Player{
		Nome: string(nome),
		Hp:   10,
		Atk:  2,
		Def:  4,
		Nvl:  1,
		Xp:   0,
	}
}
