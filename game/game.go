package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *Player
	lasers []*Laser
}

func NewGame() *Game {
	g := &Game{}
	player := NewPlayer(g)
	g.player = player
	return g
}

// RESPONSAVEL POR ATUALIZAR A LOGICA DO JOGO
func (g *Game) Update() error {
	// Atualiza nave na tela
	g.player.Update()
	// Itera e atualiza laser na tela
	for _, l := range g.lasers {
		l.Update()
	}

	return nil
}

// RESPONSAVEL POR DESENHAR OBJETOS NA TELA
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, l := range g.lasers {
		l.Draw(screen)
	}
}

// RESPONSAVEL POR RETORNAR O TAMANHO DA TELA
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}
