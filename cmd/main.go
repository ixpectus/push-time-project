package main

import (
	"fmt"

	"github.com/ixpectus/push-time-project/pkg/time_clarifier"
)

func main() {
	c := time_clarifier.New()
	c.SetDelayIntervalWhenUserOnline(1000)
	c.SetOnlineSecondsInternval(300)
	pushDelayTime := c.GetPushDelayTime(1, 10)
	fmt.Printf("\n>>> push delay time %v <<< debug\n", pushDelayTime)
}
