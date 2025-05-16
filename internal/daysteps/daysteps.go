package daysteps

import (
	"time"
	"log"
	"fmt"
	"strconv"
	"strings"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
	
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	// 1. Разделить строку на слайс строк
	parts := strings.Split(data, ",")

	// 2. Проверить длину слайса — должно быть ровно 2 части
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("ожидался формат '<шаги>,<длительность>', получено: %q", data)
	}

	// 3. Преобразовать первый элемент в int (кол-во шагов)
	stepsStr := parts[0]
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования шагов: %v", err)
	}

	// 4. Проверить, что количество шагов > 0
	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	// 5. Преобразовать второй элемент в time.Duration
	durationStr := strings.TrimSpace(parts[1])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования длительности: %v", err)
	}
	if duration <= 0 {
	return 0, 0, fmt.Errorf("длительность должна быть положительной, получено: %v", duration)
	}

	// 6. Вернуть результат, если всё прошло успешно
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println("Ошибка:", err) // ← логируем правильно
		return ""
	}

	if steps <= 0 {
		log.Println("Ошибка: количество шагов должно быть положительным")
		return ""
	}

	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println("Ошибка:", err)
		return ""
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps,
		distanceKm,
		calories,
	)
	return result
}
