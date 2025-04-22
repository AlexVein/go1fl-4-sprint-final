package spentcalories

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

var (
	ErrInvalidArguments = errors.New("invalid arguments")
	ErrInvalidFormat    = errors.New("invalid format")
)

func parseTraining(data string) (int, string, time.Duration, error) {
	var steps int
	var activityType string
	var duration time.Duration

	trainingData := strings.Split(data, ",")
	if len(trainingData) != 3 {
		return 0, "", time.Duration(0), ErrInvalidArguments
	}

	steps, err := strconv.Atoi(trainingData[0])
	if err != nil {
		return 0, "", time.Duration(0), ErrInvalidFormat
	}

	activityType = trainingData[1]

	duration, err = time.ParseDuration(trainingData[2])
	if err != nil {
		return 0, "", time.Duration(0), ErrInvalidFormat
	}

	return steps, activityType, duration, nil
}

func distance(steps int, height float64) float64 {
	return height * stepLengthCoefficient * float64(steps) / float64(mInKm)
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	averageSpeed := distance(steps, height) / duration.Hours()
	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}
