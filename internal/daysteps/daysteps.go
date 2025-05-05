package daysteps

import (
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"

	"fmt"
	"errors"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	list := strings.Split(datastring, ",")
	if len(list) != 2 {
		return errors.New("invalid data")
	}
	ds.Steps, err = strconv.Atoi(list[0])
	if err != nil {
		return err
	}
	if ds.Steps <= 0 {
		return errors.New("steps must be greater than zero")
	}
	ds.Duration, err = time.ParseDuration(list[1])
	if err != nil {
		return err
	}
	if ds.Duration <= 0 {
		return errors.New("duration must be greater than zero")
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	/*
		Количество шагов: 792.
		Дистанция составила 0.51 км.
		Вы сожгли 221.33 ккал.
	*/
	if ds.Duration <= 0 {
		return "", errors.New("duration must be greater than zero")
	}
	if ds.Steps <= 0 {
		return "", errors.New("steps must be greater than zero")
	}
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	if distance <= 0 {
		return "", errors.New("distance must be greater than zero")
	}
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, spentCalories)

	return result, nil
}
