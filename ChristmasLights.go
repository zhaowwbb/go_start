package main

type ChristmasLights struct {
	lightCount int
	pos        int
	moveRight  bool
}

func NewChristmasLights(count int) *ChristmasLights {
	return &ChristmasLights{
		lightCount: count,
		pos:        0,
		moveRight:  true,
	}
}

func (cl *ChristmasLights) Next() []bool {
	lights := make([]bool, cl.lightCount)

	left, right := 0, cl.lightCount-1
	lights[left] = true
	lights[right] = true

	if cl.pos == 0 {
		cl.pos += 1
		lights[cl.pos] = true
		return lights
	}

	if cl.pos == right-1 {
		cl.moveRight = false
	}
	if cl.pos == 1 && cl.moveRight == false {
		cl.moveRight = true
	}
	if cl.moveRight {
		cl.pos++
	} else {
		cl.pos--
	}
	lights[cl.pos] = true

	return lights
}
