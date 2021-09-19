package servo

import (
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var (
	// default RaspberryPi
	adaptor = raspi.NewAdaptor()
)

type Servo struct {
	Min    uint8
	Max    uint8
	cur    uint8
	driver *gpio.ServoDriver
}

func NewServo(pin string, min, max uint8) *Servo {
	s := &Servo{
		Min: min,
		Max: max,
	}
	s.driver = gpio.NewServoDriver(adaptor, pin)
	s.cur = (max-min)/2 + min
	s.driver.Move(s.cur)
	return s
}

func (s *Servo) Reduce(value ...uint8) {
	if s.cur <= s.Min {
		return
	}
	if value[0] != 0 {
		s.cur -= value[0]
	} else {
		s.cur -= 1
	}
	if s.cur < s.Min {
		s.cur = s.Min
	}
	s.driver.Move(s.cur)
}

func (s *Servo) Add(value ...uint8) {
	if s.cur == s.Max {
		return
	}
	if value[0] != 0 {
		s.cur += value[0]
	} else {
		s.cur += 1
	}
	if s.cur > s.Max {
		s.cur = s.Max
	}
	s.driver.Move(s.cur)
}

var (
	X *Servo
	Y *Servo
)

func Init(xpin, ypin string, xmin, xmax, ymin, ymax uint8) {
	X = NewServo(xpin, xmin, xmax)
	Y = NewServo(ypin, ymin, ymax)
}
