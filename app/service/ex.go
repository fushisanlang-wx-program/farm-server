package service

func GetLevel(Ex int) int {
	Level := 0
	switch {
	case Ex < 200:
		Level = 0
	case Ex >= 200 && Ex < 600:
		Level = 1
	case Ex >= 600 && Ex < 1200:
		Level = 2
	case Ex >= 1200 && Ex < 2000:
		Level = 3
	case Ex >= 2000 && Ex < 3000:
		Level = 4
	case Ex >= 3000 && Ex < 4200:
		Level = 5
	case Ex >= 4200 && Ex < 5600:
		Level = 6
	case Ex >= 5600 && Ex < 7200:
		Level = 7
	case Ex >= 7200 && Ex < 9000:
		Level = 8
	case Ex >= 9000 && Ex < 11000:
		Level = 9
	case Ex >= 11000 && Ex < 13200:
		Level = 10
	case Ex >= 13200 && Ex < 15600:
		Level = 11
	case Ex >= 15600 && Ex < 18200:
		Level = 12
	case Ex >= 18200 && Ex < 21000:
		Level = 13
	case Ex >= 21000 && Ex < 24000:
		Level = 14
	case Ex >= 24000:
		Level = 15

	}
	return Level

}
