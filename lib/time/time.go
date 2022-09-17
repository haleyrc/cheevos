package time

import "time"

type timeFn func() time.Time

var nowFn timeFn = time.Now

var (
	Hour       = Duration{d: time.Hour}
	Nanosecond = Duration{d: time.Nanosecond}
	Second     = Duration{d: time.Second}
)

type Duration struct {
	d time.Duration
}

type Time struct {
	t time.Time
}

func (t Time) Add(d Duration) Time {
	return Time{t: t.t.Add(d.d)}
}

func (t Time) Before(other Time) bool {
	return t.t.Before(other.t)
}

func (t Time) Equals(other Time) bool {
	return t.EqualsWithResolution(other, Nanosecond)
}

func (t Time) EqualsWithResolution(other Time, res Duration) bool {
	return t.t.Round(time.Second) == other.t.Round(time.Second)
}

func (t Time) IsZero() bool {
	return t.t.IsZero()
}

func (t Time) MarshalJSON() ([]byte, error) {
	return t.t.MarshalJSON()
}

func (t Time) String() string {
	return t.t.String()
}

func (t Time) Sub(d Duration) Time {
	return Time{t: t.t.Add(-d.d)}
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
	return Time{t: nowFn()}
}
