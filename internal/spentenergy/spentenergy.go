package spentenergy

import (
	"errors"
	"log"
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
	if weight <= 0 || height <= 0 {
		return 0, errors.New("weight or height must be greater than zero")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be greater or equal to zero")
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	return ((averageSpeed * duration.Minutes() * weight) / minInH) * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps must be greater than zero")
	}
	if weight <= 0 || height <= 0 {
		return 0, errors.New("weight or height must be greater than zero")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be greater or equal to zero")
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	return (averageSpeed * duration.Minutes() * weight) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 {
		log.Println("steps must be greater than zero")
		return 0
	}
	if height <= 0 {
		log.Println("height must be greater than zero")
		return 0
	}
	if duration <= 0 {
		log.Println("duration must be greater than zero")
		return 0
	}
	passedDistance := Distance(steps, height)
	return passedDistance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 {
		log.Println("steps must be greater than zero")
		return 0
	}
	if height <= 0 {
		log.Println("height must be greater than zero")
		return 0
	}
	stepDistance := height * stepLengthCoefficient
	return (stepDistance * float64(steps)) / mInKm
}
