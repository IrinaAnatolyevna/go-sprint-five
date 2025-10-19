package daysteps

import (
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
        return fmt.Errorf("expected 2 parameters, got %d", len(parts))
    }

    stepStr := parts[0]
    durStr := parts[1]

    // Парсим шаги
    steps, err := strconv.Atoi(stepStr)
    if err != nil {
      return fmt.Errorf("failed to parse steps: %w", err)
    }
    if steps <= 0 {
      return fmt.Errorf("steps must be greater than zero")
    }
    ds.Steps = steps


    // Парсим длительность
    duration, err := time.ParseDuration(durStr)
    if err != nil {
      return fmt.Errorf("failed to parse duration: %w", err)
    }
    if duration <= 0 {
      return fmt.Errorf("duration must be greater than zero")
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

