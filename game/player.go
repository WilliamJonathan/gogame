package game

import (
	"github.com/WilliamJonathan/gogame/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image *ebiten.Image
}

func NewPlayer() *Player {
	image := assets.PlayerSprite
	return &Player{
		image: image,
	}
}

func (p *Player) Update() {

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Posição X e Y que a imagem sera desenhada na tela
	op.GeoM.Translate(100, 100)
	// Desenha imagem na tela
	screen.DrawImage(p.image, op)

}
