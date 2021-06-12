package battle

import (
	"fmt"
	"terminal/input"
	"terminal/mobs"
	"terminal/player"
	"terminal/txt"
	"time"
)

func win(hpp, hpm int) int {
	if hpp <= 0 && hpm > 0 { // Morreu
		return -1
	} else if hpp <= 0 && hpm <= 0 { // Empate
		return 0
	} else if hpp > 0 && hpm <= 0 { // Ganhou
		return 1
	}
	return 2 // Ainda n acabou
}

func escolha() int {
	escolha, err := input.IntInput("Você ira atacar ou defender? [1/2]: ")
	if err != nil {
		panic(err)
	}
	return escolha
}

func confontro(p *player.Player, m *mobs.Mob, escolha, inimigo int) []string {
	danoPlayer, pCritico := p.Atacar(m.Def)
	danoMob, mCritico := m.Atacar(m.Def)

	var mensagem []string
	if escolha == 1 && inimigo == 1 {
		p.Hp = p.Hp - danoMob
		m.Hp = m.Hp - danoPlayer
		if pCritico {
			mensagem = append(mensagem, fmt.Sprintln("Voçê acertou um dano critico, causando ", danoPlayer, " pontos de dano"))
		} else {
			mensagem = append(mensagem, fmt.Sprintln("Você causou ", danoPlayer, " pontos de dano"))
		}
		if mCritico {
			mensagem = append(mensagem, fmt.Sprintln("O inimigo acertou um dano critico, causando ", danoMob, " pontos de dano"))
		} else {
			mensagem = append(mensagem, fmt.Sprintln("O inimigo causou ", danoMob, " pontos de dano"))
		}
	} else if escolha == 1 && inimigo == 2 {
		dano, sorte := m.Defender(danoPlayer)
		m.Hp = m.Hp - dano
		if sorte {
			mensagem = append(mensagem, fmt.Sprintln("O inimigo bloquiou seu ataque, fazendo você errou o golpe"))
		} else {
			mensagem = append(mensagem, fmt.Sprintln("O inimigo defendou o seu golpe, recebendo apenas ", dano, " de dano"))
		}
	} else if escolha == 2 && inimigo == 1 {
		dano, sorte := m.Defender(danoMob)
		p.Hp = p.Hp - dano
		if sorte {
			mensagem = append(mensagem, fmt.Sprintln("Você bloquiou o ataque inimigo, fazendo ele errar o golpe"))
		} else {
			mensagem = append(mensagem, fmt.Sprintln("Você se defende do golpe inimigo, recebendo apenas ", dano, " de dano"))
		}
	}
	return mensagem
}

func Batalha(p *player.Player, m *mobs.Mob) int {
	txt.CallClear()
	txt.Top2()
	txt.StatusBatalha(p, m)
	input.StrInput(fmt.Sprintf("Um %s apareceu e inicio uma batalha: ", m.Nome))
	var re []string
	for {
		txt.CallClear()
		txt.Top2()
		txt.StatusBatalha(p, m)

		for _, v := range re {
			fmt.Print(v)
		}

		escolha := escolha()
		inimigo := 1
		re = confontro(p, m, escolha, inimigo)

		time.Sleep(1000)

		txt.CallClear()
		txt.Top2()
		txt.StatusBatalha(p, m)

		time.Sleep(1000)

		ganhou := win(p.Hp, m.Hp)
		if ganhou == -1 {
			return -1
		} else if ganhou == 0 {
			p.Upar(m.Nvl, m.XPBase)
			return 0
		} else if ganhou == 1 {
			p.Upar(m.Nvl, m.XPBase)
			return 1
		}
	}
}
