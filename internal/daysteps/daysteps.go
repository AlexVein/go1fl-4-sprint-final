package daysteps

import (
	"errors"
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

var (
	ErrInvalidArguments      = errors.New("invalid arguments")
	ErrInvalidArgumentsCount = errors.New("invalid arguments count")
	ErrInvalidFormat         = errors.New("invalid format")
	ErrZeroOrNegativeValue   = errors.New("zero or negative value")
)

func parsePackage(data string) (int, time.Duration, error) {
	var steps int
	var duration time.Duration

	trainingRecord := strings.Split(data, ",")
	if len(trainingRecord) != 2 {
		return 0, 0, fmt.Errorf("%w: %s", ErrInvalidArgumentsCount, data)
	}

	steps, err := strconv.Atoi(trainingRecord[0])
	if err != nil {
		return 0, 0, fmt.Errorf("%w: %s", ErrInvalidFormat, trainingRecord[0])
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("%w: %d", ErrZeroOrNegativeValue, steps)
	}

	duration, err = time.ParseDuration(trainingRecord[1])
	if err != nil {
		return 0, 0, fmt.Errorf("%w: %s", ErrInvalidFormat, trainingRecord[1])
	}

	if duration <= 0 {
		return 0, 0, fmt.Errorf("%w: %s", ErrZeroOrNegativeValue, duration)
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	var distance float64
	var distanceInKm float64
	var spentCalories float64

	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	if steps <= 0 {
		log.Println(ErrZeroOrNegativeValue)
		return ""
	}

	if duration <= 0 {
		log.Println(ErrZeroOrNegativeValue)
		return ""
	}

	distance = float64(steps) * stepLength
	distanceInKm = distance / mInKm

	spentCalories, err = spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(ErrInvalidArguments)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distanceInKm, spentCalories)
}
