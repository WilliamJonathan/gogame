package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

// RESPONSAVEL POR ATUALIZAR A LOGICA DO JOGO
func (g *Game) Update() error {
	return nil
}

// RESPONSAVEL POR DESENHAR OBJETOS NA TELA
func (g *Game) Draw(screen *ebiten.Image) {

}

// RESPONSAVEL POR RETORNAR O TAMANHO DA TELA
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {

	g := &Game{}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}

}
