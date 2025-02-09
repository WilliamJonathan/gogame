package game

import (
	"github.com/WilliamJonathan/gogame/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image    *ebiten.Image
	position Vector
	game     *Game
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite

	// Pega tamanho da nave
	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := Vector{
		x: (screenWidth / 2) - halfW,
		y: 500,
	}

	return &Player{
		image:    image,
		game:     game,
		position: position,
	}
}

func (p *Player) Update() {

	speed := 6.0

	// Movimenta nave de uma lado para o outro
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.x -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.x += speed
	}

	// Dispara o laser
	if ebiten.IsKeyPressed(ebiten.KeySpace) {

		bounds := p.image.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		spawnPos := Vector{
			p.position.x + halfW,
			p.position.y + halfH/2,
		}

		laser := NewLaser(spawnPos)
		p.game.AddLasers(laser)
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Posição X e Y que a imagem sera desenhada na tela
	op.GeoM.Translate(p.position.x, p.position.y)
	// Desenha imagem na tela
	screen.DrawImage(p.image, op)

}
