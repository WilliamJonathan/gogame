package game

import (
	"github.com/WilliamJonathan/gogame/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Laser struct {
	image    *ebiten.Image
	position Vector
}

// Recebe postion da nave pois precisa dar a impressão
// que o Laser esta saindo da nave
func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite

	// Pega tamanho da nave
	bounds := image.Bounds()

	halfW := float64(bounds.Dx()) / 2 // Metade da largura da imgem do laser
	halfH := float64(bounds.Dy()) / 2 // Metade da altura da imagem do laser

	position.x -= halfW
	position.y -= halfH

	return &Laser{
		image:    image,
		position: position,
	}
}

func (l *Laser) Update() {
	speed := 7.0

	l.position.y += -speed
}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Posição X e Y que a imagem sera desenhada na tela
	op.GeoM.Translate(l.position.x, l.position.y)
	// Desenha imagem na tela
	screen.DrawImage(l.image, op)
}
