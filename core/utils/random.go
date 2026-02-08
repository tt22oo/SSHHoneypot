package utils

import (
	"honeypot/core/configs"
	"math/rand/v2"
	"time"
)

func RandomSleep() {
	time.Sleep(time.Duration(rand.IntN(configs.Cfg.Config.MaxDelay)) * time.Millisecond)
}
