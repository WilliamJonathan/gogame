package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *Player
}

func NewGame() *Game {
	player := NewPlayer()
	return &Game{
		player: player,
	}
}

// RESPONSAVEL POR ATUALIZAR A LOGICA DO JOGO
func (g *Game) Update() error {
	g.player.Update()
	return nil
}

// RESPONSAVEL POR DESENHAR OBJETOS NA TELA
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

// RESPONSAVEL POR RETORNAR O TAMANHO DA TELA
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
