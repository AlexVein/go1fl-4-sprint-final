package daysteps

import (
	"errors"
	"fmt"
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
	ErrInvalidArguments = errors.New("invalid arguments")
	ErrInvalidFormat    = errors.New("invalid format")
)

func parsePackage(data string) (int, time.Duration, error) {
	var steps int
	var duration time.Duration

	trainingRecord := strings.Split(data, ",")
	if len(trainingRecord) != 2 {
		return 0, 0, ErrInvalidArguments
	}

	steps, err := strconv.Atoi(trainingRecord[0])
	if err != nil {
		return 0, 0, ErrInvalidFormat
	}

	duration, err = time.ParseDuration(trainingRecord[1])
	if err != nil {
		return 0, 0, ErrInvalidFormat
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	var distance float64
	var distanceInKm float64
	var spentCalories float64

	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distance = float64(steps) * stepLength
	distanceInKm = distance / mInKm

	spentCalories, err = spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		steps, distanceInKm, spentCalories)
}
