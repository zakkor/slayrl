package main

type builder struct {
	w          *World
	marked     [][]bool
	negateNext bool
}

func newBuilder(w *World) *builder {
	marked := make([][]bool, w.SizeX)
	for x := 0; x < w.SizeX; x++ {
		marked[x] = make([]bool, w.SizeY)
	}
	return &builder{
		w:      w,
		marked: marked,
	}
}

func (b *builder) All() *builder {
	defer b.afterCall()

	for x := 0; x < b.w.SizeX; x++ {
		for y := 0; y < b.w.SizeY; y++ {
			b.mark(x, y)
		}
	}
	return b
}

func (b *builder) Point(x, y int) *builder {
	defer b.afterCall()

	b.mark(x, y)
	return b
}

func (b *builder) Rect(x1, y1, x2, y2 int) *builder {
	defer b.afterCall()

	return b.Line(x1, y1, x2, y1).
		Line(x2, y1, x2, y2).
		Line(x1, y2, x2, y2).
		Line(x1, y1, x1, y2)
}

func (b *builder) RectFill(x1, y1, x2, y2 int) *builder {
	defer b.afterCall()

	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			b.mark(x, y)
		}
	}
	return b
}

func (b *builder) Circle(x, y, r int) *builder {
	defer b.afterCall()

	points := CirclePoints(x, y, r)
	for _, point := range points {
		b.mark(point.X, point.Y)
	}

	return b
}

func (b *builder) CircleThick(x, y, r int) *builder {
	defer b.afterCall()

	points := CircleThickPoints(x, y, r)
	for _, point := range points {
		b.mark(point.X, point.Y)
	}

	return b
}

func (b *builder) Line(x1, y1, x2, y2 int) *builder {
	defer b.afterCall()

	points := LinePoints(x1, y1, x2, y2)
	for _, point := range points {
		b.mark(point.X, point.Y)
	}

	return b
}

func (b *builder) Except() *builder {
	b.negateNext = true
	return b
}

func (b *builder) Do(fn func(*Entity)) {
	for x := 0; x < b.w.SizeX; x++ {
		for y := 0; y < b.w.SizeY; y++ {
			if !b.marked[x][y] {
				continue
			}
			fn(b.w.At(x, y))
		}
	}
}

func (b *builder) mark(x, y int) {
	val := true
	if b.negateNext {
		val = false
	}
	b.marked[x][y] = val
}

func (b *builder) afterCall() {
	b.negateNext = false
}
