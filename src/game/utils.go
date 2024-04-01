package game

type Vector struct {
	X float64
	Y float64
}

type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func NewRect(x, y, width, height float64) Rect {
	return Rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (r Rect) Intersects(other Rect) bool {
	return r.X <= other.maxX() &&
		other.X <= r.maxX() &&
		r.Y <= other.maxY() &&
		other.Y <= r.maxY()
}

func (r Rect) maxX() float64 {
	return r.X + r.Width
}

func (r Rect) maxY() float64 {
	return r.Y + r.Height
}
