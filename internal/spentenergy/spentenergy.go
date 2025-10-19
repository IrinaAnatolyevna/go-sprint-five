package spentenergy

import (
        "errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
   if steps <= 0 {
        return 0, errors.New("steps must be greater than zero")
    }
    if weight <= 0 {
        return 0, errors.New("weight must be greater than zero")
    }
    if height <= 0 {
        return 0, errors.New("height must be greater than zero")
    }
    if duration <= 0 {
        return 0, errors.New("duration must be greater than zero")
    }

    speed := MeanSpeed(steps, height, duration)
    durationMinutes := duration.Minutes()

    calories := (weight * speed * durationMinutes) / minInH
    calories *= walkingCaloriesCoefficient // корректирующий коэффициент
    return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
    if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
	return 0, errors.New("invalid input parameters")
    }

    speed := MeanSpeed(steps, height, duration)
    durationMinutes := duration.Minutes()

    calories := (weight * speed * durationMinutes) / minInH
    return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
    if steps <= 0 || height <= 0 || duration <= 0 {
	return 0
    }
    distance := Distance(steps, height)
    hours := duration.Hours()
    if hours <= 0 {
	return 0
    }
    return distance / hours
}

func Distance(steps int, height float64) float64 {
    if steps <= 0 || height <= 0 {
	return 0
    }
    stepLength := height * stepLengthCoefficient        // длина шага в метрах
    distance := float64(steps) * stepLength / mInKm     // переводим в километры
    return distance
}
