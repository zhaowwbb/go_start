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

type ChristmasLightV2 struct {
	lightCount int
	pos        int
	moveRight  bool
}

func CreateCLV2(count int) *ChristmasLightV2 {
	return &ChristmasLightV2{
		lightCount: count,
		pos:        0,
		moveRight:  true,
	}
}

func (cl2 *ChristmasLightV2) Next() []bool {
	lights := make([]bool, cl2.lightCount)
	left := 0
	right := cl2.lightCount - 1
	lights[left] = true
	lights[right] = true
	if cl2.pos == 0 {
		cl2.pos++
		lights[cl2.pos] = true
		return lights
	}
	if cl2.pos == right-1 && cl2.moveRight {
		cl2.moveRight = false
	}
	if cl2.pos == 1 {
		cl2.moveRight = true
	}

	if cl2.moveRight {
		cl2.pos++
	} else {
		cl2.pos--
	}
	lights[cl2.pos] = true
	return lights
}
