package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

// преобразование строки
func parsePackage(data string) (int, time.Duration, error) {

	sliceStr := strings.Split(data, ",")

	if len(sliceStr) != 2 {

		return 0, 0, fmt.Errorf("неверная длина слайса")
	} // проверка длины слайса

	steps, err := strconv.Atoi(sliceStr[0])

	if err != nil {

		return 0, 0, fmt.Errorf("конвертация количества шагов не выполнена")
	} // проверка конвертации шагов

	if steps <= 0 {

		return 0, 0, fmt.Errorf("неверное количество шагов")
	}

	duration, err := time.ParseDuration(sliceStr[1])

	if err != nil {
		return 0, 0, fmt.Errorf("ошибка парсинга продолжительности")
	} // проверка парсинга продолжительности прогулки

	if duration <= 0 {
		return 0, 0, fmt.Errorf("проверка продолжительности")
	}

	return steps, duration, nil
}

// вычисление дистанции и калорий
func DayActionInfo(data string, weight, height float64) string {

	steps, duration, err := parsePackage(data)

	if err != nil {
		log.Println(err)
		return ""
	}

	if duration <= 0 {
		log.Println("ноль продолжительности")
		return ""
	}

	distance := float64(steps) * stepLength

	distKm := distance / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	if err != nil {
		log.Println(err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distKm, calories)

}
