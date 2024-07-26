package main

type BatteryLevel = int

func CheckBatteryLevel(b BatteryLevel) string {
	if b > 75 {
		return "Высокий уровень заряда"
	} else if b >= 35 && b <= 75 {
		return "Средний уровень заряда"
	}
	return "Низкий уровень заряда"
}
