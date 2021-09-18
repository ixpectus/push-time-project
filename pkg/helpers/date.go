package helpers

import "time"

type UserActivity struct {
	Time        time.Time
	SecondsFrom int64
}

func IsWorkingDay(hourStartUTC, workingDayHours, hourNowUTC int) bool {
	// В случае, если завершение рабочего дня по UTC перехало на утро (например на Камчатке по utc рабочий день начинается в 22)
	// Нужно пересчитать время завершения рабочего дня, которое будет уже утром
	if hourStartUTC+workingDayHours > 23 {
		hourFinishUTC := 24 - hourStartUTC + workingDayHours
		if hourNowUTC < hourFinishUTC {
			return true
		}
	}
	hourFinishUTC := hourStartUTC + workingDayHours
	return hourNowUTC >= hourStartUTC && hourNowUTC < hourFinishUTC
}

func MorningDiff(hourStartUTC, hourNowUTC int) time.Duration {
	if hourNowUTC < hourStartUTC {
		return time.Duration(hourStartUTC-hourNowUTC) * time.Hour
	}
	return time.Duration(24-hourNowUTC+hourStartUTC) * time.Hour
}
