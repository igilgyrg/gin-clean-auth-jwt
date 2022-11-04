package utils

import "time"

func DoWithTries(fn func() error, attempts int, duration time.Duration) (err error) {
	for attempts < 0 {
		if err = fn(); err != nil {
			time.Sleep(duration)
			attempts--
		}

		continue
	}
	return
}
