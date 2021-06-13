// PACOTE RESPONSAVEL POR CRIAR E GERENCIAR JOGADORES

package player

import (
	"fmt"
	"math/rand"
	"strings"
	"terminal/input"
)

type Player struct {
	Nome string  // Nome
	MHp  int     // vida total
	Hp   int     // vida atual
	Atk  int     // atauqe
	Def  int     // Defesa
	Nvl  int     // Nivel
	Xp   float64 // Total de XP
	Next int     // Total de XP para o proximo nivel
}

// Calcula o XP ganho apos uma batalha
func calculaXp(pn, ml, mxp int) float64 {
	return float64(mxp*ml) + (float64(pn/2) + 1.5)
}

// Metodo para Atacar
func (p *Player) Atacar(def int) (int, bool) {
	critico := rand.Intn(100)
	if critico <= 5 {
		return (p.Atk / ((def + 1) / 2)) * 2, true
	}
	return (p.Atk / ((def + 1) / 2)), false
}

// Metodo para Defender
func (p *Player) Defender(dano int) (int, bool) {
	sorte := rand.Intn(100)
	if sorte <= 25 {
		return 0, true
	}
	return (dano / 2) + 1, false
}

// Metodo par aupar de nivel
func (p *Player) Upar(ml, mxp int) {
	p.Xp = p.Xp + calculaXp(p.Nvl, ml, mxp)
	if p.Xp >= float64(p.Next) {
		p.Nvl++
		p.Xp = 0
		p.Next = p.Next * 2
		p.MHp = p.MHp + p.MHp/2
		p.Hp = p.MHp
		p.Atk = p.Atk + p.Atk/2
		p.Def = p.Def + p.Def/2
	}
}

// Função para criar um jogador novo
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
