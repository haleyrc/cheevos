package time

import "time"

type timeFn func() time.Time

var nowFn timeFn = time.Now

type Time struct {
	time.Time
}

func Freeze() {
	nowFn = func() time.Time {
		t, err := time.Parse(time.RFC3339, "2022-09-16T15:02:04-04:00")
		if err != nil {
			panic(err)
		}
		return t
	}
}

func Unfreeze() {
	nowFn = time.Now
}

func Now() Time {
	return Time{Time: nowFn()}
}
