package daysteps

import (
	//"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
    parts := strings.Split(datastring, ",")
    if len(parts) != 2 {
        return fmt.Errorf("ожидалось 2 параметра, получено %d", len(parts))
    }

    stepStr := parts[0]
    durStr := parts[1]

    // Проверка пробелов в начале/конце
    if stepStr != strings.TrimSpace(stepStr) {
        return fmt.Errorf("шаги не должны содержать пробелов в начале или конце")
    }
    if durStr != strings.TrimSpace(durStr) {
        return fmt.Errorf("продолжительность не должна содержать пробелов в начале или конце")
    }

    // Парсим шаги
    steps, err := strconv.Atoi(stepStr)
    if err != nil || steps <= 0 {
        return fmt.Errorf("некорректное количество шагов: %s", stepStr)
    }
    ds.Steps = steps

    // Парсим длительность
    duration, err := time.ParseDuration(durStr)
    if err != nil || duration <= 0 {
        return fmt.Errorf("некорректная продолжительность: %s", durStr)
    }
    ds.Duration = duration

    return nil
}

func (ds DaySteps) ActionInfo() (string, error) {

     dist := spentenergy.Distance(ds.Steps, ds.Height)

     calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
     if err != nil {
	return "", err
     }

     result := fmt.Sprintf(
	"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
	ds.Steps,
	dist,
	calories,
      )

      return result, nil
}

