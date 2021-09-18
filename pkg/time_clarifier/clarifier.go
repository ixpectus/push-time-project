package time_clarifier

import (
	"strings"
	"time"

	"github.com/ixpectus/push-time-project/pkg/helpers"
	"github.com/ixpectus/push-time-project/pkg/user_online"
)

type PushTimeClarifier struct {
	userOnlineClient            user_online.Client
	onlineSecondsInterval       int
	delayIntervalWhenUserOnline time.Duration
}

func New() *PushTimeClarifier {
	userOnlineClient := user_online.New()
	return &PushTimeClarifier{
		userOnlineClient: userOnlineClient,
	}
}

func (p PushTimeClarifier) SetOnlineSecondsInternval(i int) {
	p.onlineSecondsInterval = i
}

func (p PushTimeClarifier) SetDelayIntervalWhenUserOnline(t time.Duration) {
	p.delayIntervalWhenUserOnline = t
}

func (d *PushTimeClarifier) GetPushDelayTime(userID, dayStartWork int) time.Duration {
	res := delayForStartBotActivity(
		userID,
		dayStartWork,
		time.Now().UTC().Hour(),
		d.onlineSecondsInterval,
		d.delayIntervalWhenUserOnline,
		d.userOnlineClient.Get,
	)
	return res
}

func delayForStartBotActivity(
	userID int,
	hourStartUTC, hoursUTCNow, onlineSecondsInterval int,
	delayIntervalWhenUserOnline time.Duration,
	loadUserActivity func(int64) (error, *helpers.UserActivity),
) time.Duration {
	// Нужно проверить а ночь ли сейчас
	// Если сейчас не рабочий день(т.е. ночь), то нужно сдвинуть на утро
	if !helpers.IsWorkingDay(hourStartUTC, 13, hoursUTCNow) {
		return helpers.MorningDiff(hourStartUTC, hoursUTCNow)
	}
	err, lastActivity := loadUserActivity(int64(userID))
	strings.Contains(err.Error(), "no data")
	if err != nil && !strings.Contains(err.Error(), "no data") {
		// в случае ошибки попробуем завтра
		return helpers.MorningDiff(hourStartUTC, hoursUTCNow)
	}
	// если данных нет, значит пользователь давно не был онлайн, отправляем сразу
	if strings.Contains(err.Error(), "no data") {
		return 0
	}
	// Если пользователь недавно был онлайн, то отправим сообщение немного позже
	if lastActivity.SecondsFrom > int64(onlineSecondsInterval) {
		return 0
	}
	return delayIntervalWhenUserOnline
}
