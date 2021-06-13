// PACOTE RESPONSAVEL POR GERENCIAR OS EVENTOS EM BATALHAS

package battle

import (
	"fmt"
	"terminal/input"
	"terminal/mobs"
	"terminal/player"
	"terminal/txt"
)

// Verifica se o jogaodr ganhou perdeu ou empatou
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

// Recebe a escolha do jogoador se ele quer Atacar ou Defender
func escolha() int {
	escolha, err := input.IntInput("Você ira atacar ou defender? [1/2]: ")
	if err != nil {
		panic(err)
	}
	return escolha
}

// Realiza o combate entre o jogador e o inimigo
func confontro(p *player.Player, m *mobs.Mob, escolha, inimigo int) []string {
	// Calcula o dano que o jogador causa no mob e se deu critico
	danoPlayer, pCritico := p.Atacar(m.Def)
	// Calcula o dano que o mob causa no jogador e se deu critico
	danoMob, mCritico := m.Atacar(m.Def)

	var mensagem []string
	if escolha == 1 && inimigo == 1 { // Caso ambos descidam atacar
		// Aplica o dano
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
	} else if escolha == 1 && inimigo == 2 { // Caso so o jogador queira bater e o inimigo tenha defendido
		// Calcula o dano que o inimigo recebe se defendendo e aplica o dano
		dano, sorte := m.Defender(danoPlayer) // E se o jogaodr errou o golpe
		m.Hp = m.Hp - dano
		if sorte {
			mensagem = append(mensagem, fmt.Sprintln("O inimigo bloquiou seu ataque, fazendo você errou o golpe"))
		} else {
			mensagem = append(mensagem, fmt.Sprintln("O inimigo defendou o seu golpe, recebendo apenas ", dano, " de dano"))
		}
	} else if escolha == 2 && inimigo == 1 { // Caso o jogador defenda, porem o inimigo ataque
		// Calcula o dano que o jogador recebe se defendendo e aplica o dano
		dano, sorte := m.Defender(danoMob) // E se o inimigo errou o golpe
		p.Hp = p.Hp - dano
		if sorte {
			mensagem = append(mensagem, fmt.Sprintln("Você bloquiou o ataque inimigo, fazendo ele errar o golpe"))
		} else {
			mensagem = append(mensagem, fmt.Sprintln("Você se defende do golpe inimigo, recebendo apenas ", dano, " de dano"))
		}
	} else if escolha == 2 && inimigo == 2 { // Caso ambos tenham defendido
		mensagem = append(mensagem, fmt.Sprintln("Albos se prepararam para defender um golpe, nada aconteceu"))
	}
	return mensagem
}

// Função que gerencia a batalha
func Batalha(p *player.Player, m *mobs.Mob) int {
	txt.CallClear()
	txt.Top2()
	txt.StatusBatalha(p, m)
	input.StrInput(fmt.Sprintf("Um %s apareceu e iniciou uma batalha ", m.Nome))
	// Mensagem que aparecem a baixo da barra de status
	var re []string
	for {
		txt.CallClear()
		txt.Top2()
		txt.StatusBatalha(p, m)

		for _, v := range re {
			fmt.Print(v)
		}

		// BATALHA --------------
		escolha := escolha()
		inimigo := 1
		re = confontro(p, m, escolha, inimigo)
		// ----------------------

		// Verifica se o jogador ganhou ou perdeu
		ganhou := win(p.Hp, m.Hp)
		if ganhou == -1 { // Pedeu
			return -1
		} else if ganhou == 0 { // Empate
			// Calcula o XP ganho em batalha
			p.Upar(m.Nvl, m.XPBase)
			return 0
		} else if ganhou == 1 { // Ganhou
			// Calcula o XP ganho em batalha
			p.Upar(m.Nvl, m.XPBase)
			return 1
		}
	}
}
