package spentcalories

import (
	"errors"
	"fmt"
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
	ErrInvalidArguments    = errors.New("invalid arguments")
	ErrInvalidFormat       = errors.New("invalid format")
	ErrUnknownTrainingType = errors.New("неизвестный тип тренировки")
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

	if steps <= 0 || duration <= 0 {
		return 0, "", time.Duration(0), ErrInvalidArguments
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
	var trainingDistance float64
	var averageSpeed float64
	var spentCalories float64
	var trainingInfo string

	steps, activityType, duration, err := parseTraining(data)
	if err != nil {
		return "", ErrInvalidArguments
	}

	switch activityType {
	case "Бег":
		trainingDistance = distance(steps, height)
		averageSpeed = meanSpeed(steps, height, duration)
		spentCalories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		trainingDistance = distance(steps, height)
		averageSpeed = meanSpeed(steps, height, duration)
		spentCalories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	default:
		return "", ErrUnknownTrainingType
	}

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60

	durationFormatted := fmt.Sprintf("%d.%02d", hours, minutes)

	trainingInfo = fmt.Sprintf("Тип тренировки: %s\nДлительность: %s ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activityType, durationFormatted, trainingDistance, averageSpeed, spentCalories)
	return trainingInfo, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	var averageSpeed float64
	var spentCalories float64

	isDataCorrect := steps > 0 && weight > 0 && height > 0 && duration > 0
	if !isDataCorrect {
		return 0, ErrInvalidArguments
	}

	averageSpeed = meanSpeed(steps, height, duration)
	spentCalories = weight * averageSpeed * duration.Minutes() / minInH

	return spentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	var averageSpeed float64
	var spentCalories float64

	isDataCorrect := steps > 0 && weight > 0 && height > 0 && duration > 0
	if !isDataCorrect {
		return 0, ErrInvalidArguments
	}

	averageSpeed = meanSpeed(steps, height, duration)
	spentCalories = (weight * averageSpeed * duration.Minutes() / minInH) * walkingCaloriesCoefficient

	return spentCalories, nil
}
