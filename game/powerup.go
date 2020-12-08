package game

//"Powerup" that randomly selects a color from a list and assigns it to the snake.
var powerUpList = []rune{
	'🌈',
}

type PowerUp struct {
	char  rune
	coord Coordinate
}

func PowerUpAt(c Coordinate) PowerUp {
	return PowerUp{
		char:  CharPowerUp(),
		coord: c,
	}
}

func CharPowerUp() rune {
	if RuneSupport() {
		return powerUpList[0]
	}
	return '?'
}
