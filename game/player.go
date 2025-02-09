package game

import (
	"github.com/WilliamJonathan/gogame/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image    *ebiten.Image
	position Vector
}

func NewPlayer() *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := Vector{
		x: (screenWidth / 2) - halfW,
		y: 500,
	}

	return &Player{
		image:    image,
		position: position,
	}
}

func (p *Player) Update() {

	speed := 6.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.x -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.x += speed
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Posição X e Y que a imagem sera desenhada na tela
	op.GeoM.Translate(p.position.x, p.position.y)
	// Desenha imagem na tela
	screen.DrawImage(p.image, op)

}
