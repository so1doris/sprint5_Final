package trainings

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"

	"errors"
	"strconv"
	"strings"
	"time"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	list := strings.Split(datastring, ",")
	if len(list) != 3 {
		return errors.New("invalid data")
	}
	t.Steps, err = strconv.Atoi(list[0])
	if err != nil {
		return err
	}
	if t.Steps <= 0 {
		return errors.New("steps must be greater than zero")
	}
	t.Duration, err = time.ParseDuration(list[2])
	if err != nil {
		return err
	}
	if t.Duration <= 0 {
		return errors.New("duration must be greater than zero")
	}
	t.TrainingType = list[1]
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	if distance <= 0 {
		return "", errors.New("distance must be greater than zero")
	}
	averageSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	if averageSpeed <= 0 {
		return "", errors.New("average speed must be greater than zero")
	}
	var spentCalories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	/*
		Тип тренировки: Бег
		Длительность: 0.75 ч.
		Дистанция: 10.00 км.
		Скорость: 13.34 км/ч
		Сожгли калорий: 18621.75
	*/
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		averageSpeed,
		spentCalories)
	return result, nil
}
