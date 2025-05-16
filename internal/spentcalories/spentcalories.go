package spentcalories

import (
	"time"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	// 1. Разделить строку на слайс строк
	parts := strings.Split(data, ",")
	// 2. Проверить длину слайса — должно быть ровно 3 части
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("ожидался формат '<шаги>,<длительность>', получено: %q", data)
	}
	// 3. Преобразовать первый элемент в int (кол-во шагов)
	stepsStr := strings.TrimSpace(parts[0])
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования шагов: %v", err)
	}
	// 4. Извлечь вид активности
	activity := strings.TrimSpace(parts[1])

	// 5. Преобразовать длительность
	durationStr := strings.TrimSpace(parts[2])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования длительности: %v", err)
	}

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	// 1. Вычисляем длину шага в метрах
	stepLength := height * stepLengthCoefficient
	// 2. Преобразуем steps в float64, чтобы можно было умножать
	totalDistanceMeters := float64(steps) * stepLength
	// 3. Переводим метры в километры
	distanceKm := totalDistanceMeters / mInKm
	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	// 1. Проверка: продолжительность должна быть больше 0
	if duration <= 0 {
		return 0
	}

	// 2. Вычислить расстояние в километрах
	dist := distance(steps, height)

	// 3. Перевести длительность в часы (как float64)
	hours := duration.Hours()

	// 4. Рассчитать среднюю скорость
	speed := dist / hours

	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	// 1. Разбор строки
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return "", fmt.Errorf("ошибка разбора данных: %v", err)
	}

	// 2. Проверка и расчёты по типу тренировки
	var (
		dist  float64
		speed float64
		cal   float64
	)

	dist = distance(steps, height)
	speed = meanSpeed(steps, height, duration)

	switch strings.ToLower(activity) {
	case "бег":
		cal, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

	case "ходьба":
		cal, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	// 3. Формирование строки результата
	result := fmt.Sprintf("Тип: %s; Длительность: %.0f мин; Расстояние: %.2f км; Ср. скорость: %.2f км/ч; Калории: %.2f",
		activity,
		duration.Minutes(),
		dist,
		speed,
		cal,
	)

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// 1. Проверка входных данных
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть больше 0")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0, errors.New("длительность должна быть больше 0")
	}

	// 2. Рассчитать среднюю скорость в км/ч
	speed := meanSpeed(steps, height, duration)

	// 3. Преобразовать длительность в минуты
	minutes := duration.Minutes()

	// 4. Вычислить количество калорий (примерная формула)
	calories := (weight * speed * minutes) / 60.0

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// 1. Проверка входных данных
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть больше 0")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0, errors.New("длительность должна быть больше 0")
	}

	// 2. Рассчитать среднюю скорость в км/ч
	speed := meanSpeed(steps, height, duration)

	// 3. Преобразовать длительность в минуты
	minutes := duration.Minutes()

	// 4. Вычислить количество калорий (примерная формула)
	calories := (weight * speed * minutes) / 60.0

	// 5. Умножение на поправочный коэффициент
	adjustedCalories := calories * walkingCaloriesCoefficient
		
		return adjustedCalories, nil
}
