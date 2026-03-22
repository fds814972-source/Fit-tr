package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	//lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

// преобразование строки
func parseTraining(data string) (int, string, time.Duration, error) {

	sliceStr := strings.Split(data, ",")

	if len(sliceStr) != 3 {

		return 0, "", 0, fmt.Errorf("неверная длина слайса")
	}

	steps, err := strconv.Atoi(sliceStr[0])

	if err != nil {

		return 0, "0", 0, err
	}

	training := strings.TrimSpace(sliceStr[1])

	duration, err := time.ParseDuration(sliceStr[2])

	if err != nil {

		return 0, "0", 0, err
	}

	return steps, training, duration, nil
}

// расчет дистанции
func distance(steps int, height float64) float64 {

	stepLength := height * stepLengthCoefficient

	distance := float64(steps) * stepLength

	distance = distance / mInKm

	return distance

}

// расчет средней скорости
func meanSpeed(steps int, height float64, duration time.Duration) float64 {

	if duration < 0 {
		return 0
	}

	meanSpeed := distance(steps, height) / duration.Hours()

	return meanSpeed

}

// расчет и вывод данных для каждой тренировки
func TrainingInfo(data string, weight, height float64) (string, error) {

	steps, training, duration, err := parseTraining(data)

	if err != nil {
		log.Println(err)
		// TODO: реализовать функцию
	}

	var calories float64
	var Err error

	switch training {
	case "Бег":
		calories, Err = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		calories, Err = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", training)

	}

	if Err != nil {
		log.Println(Err)
		return "", Err
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)

	return fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n", training, duration.Hours(), dist, speed, calories), nil

}

// расчет потраченых калорий при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	switch {
	case steps < 0:
		return 0, fmt.Errorf("количество шагов неверно")
	case weight < 0:
		return 0, fmt.Errorf("вес неверный")
	case height < 0:
		return 0, fmt.Errorf("рост неверный")
	case duration < 0:
		return 0, fmt.Errorf("продолжительность неверная")
	}

	mSpeed := meanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	spentCalories := (durationInMinutes * weight * mSpeed) / minInH

	return spentCalories, nil
}

// расчет потраченых калрий при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	switch {
	case steps < 0:
		return 0, fmt.Errorf("количество шагов неверно")
	case weight < 0:
		return 0, fmt.Errorf("вес неверный")
	case height < 0:
		return 0, fmt.Errorf("рост неверный")
	case duration < 0:
		return 0, fmt.Errorf("продолжительность неверная")
	}

	mSpeed := meanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	spentCalories := (durationInMinutes * weight * mSpeed) / minInH

	spentCalories = spentCalories * walkingCaloriesCoefficient

	return spentCalories, nil

}
