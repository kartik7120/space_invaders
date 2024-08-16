package utils

type Rect struct {
	X, Y, Width, Height float64
}

func NewRect(x, y, w, h float64) Rect {
	return Rect{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
	}
}

func (r Rect) MaxX() float64 {
	return r.X + r.Width
}

func (r Rect) MaxY() float64 {
	return r.Y + r.Height
}

func (r Rect) Intersects(other Rect) bool {
	return r.X < other.MaxX() &&
		r.MaxX() > other.X &&
		r.Y < other.MaxY() &&
		r.MaxY() > other.Y
}
