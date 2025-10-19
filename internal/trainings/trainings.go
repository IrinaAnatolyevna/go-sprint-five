package trainings


import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)


type Training struct {
    Steps        int
    TrainingType string
    Duration     time.Duration
    personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
    parts := strings.Split(datastring, ",")
    if len(parts) != 3 {
        return errors.New("invalid data format: expected 3 elements")
    }
    stepStr := parts[0]
    if strings.ContainsAny(stepStr, " \t") {
        return fmt.Errorf("steps contain spaces")
    }
    // шаги
    steps, err := strconv.Atoi(stepStr)
    if err != nil {
        return fmt.Errorf("failed to parse steps: %w", err)
    }
    if steps <= 0 {
        return errors.New("steps must be positive")
    }
    t.Steps = steps

    // тип тренировки
    t.TrainingType = strings.TrimSpace(parts[1])

    // длительность
    duration, err := time.ParseDuration(strings.TrimSpace(parts[2]))
    if err != nil {
        return fmt.Errorf("failed to parse duration: %w", err)
    }
    if duration <= 0 {
        return errors.New("duration must be positive")
    }
    t.Duration = duration

    return nil
}

func (t Training) ActionInfo() (string, error) {
    dist := spentenergy.Distance(t.Steps, t.Height)
    speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

    var calories float64
    var err error

    switch strings.ToLower(t.TrainingType) {
    case "бег":
	calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
    case "ходьба":
	calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
    default:
	return "", errors.New("неизвестный тип тренировки")
    }

    if err != nil {
	return "", err
    }

    result := fmt.Sprintf(
	"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
	strings.Title(t.TrainingType),
	t.Duration.Hours(),
	dist,
	speed,
	calories,
     )

     return result, nil
}
