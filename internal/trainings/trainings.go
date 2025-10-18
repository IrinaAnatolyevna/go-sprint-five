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

func (t *Training) Parse(datastring string) error {
    parts := strings.Split(datastring, ",")
    if len(parts) != 3 {
        return errors.New("некорректный формат данных: ожидалось 3 элемента")
    }

    // шаги
    steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
    if err != nil {
        return fmt.Errorf("ошибка парсинга количества шагов: %w", err)
    }
    if steps <= 0 {
        return errors.New("шаги должны быть положительными")
    }
    t.Steps = steps

    // тип тренировки
    rawType := strings.TrimSpace(parts[1])
    normalizedType := strings.ToLower(rawType)

    switch normalizedType {
    case "бег":
        t.TrainingType = "Бег"
    case "ходьба":
        t.TrainingType = "Ходьба"
    case "плавание":
        t.TrainingType = "Плавание"
    default:
        return fmt.Errorf("неизвестный тип тренировки: %s", rawType)
    }

    // длительность
    duration, err := time.ParseDuration(strings.TrimSpace(parts[2]))
    if err != nil {
        return fmt.Errorf("ошибка парсинга длительности: %w", err)
    }
    if duration <= 0 {
        return errors.New("длительность должна быть положительной")
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
