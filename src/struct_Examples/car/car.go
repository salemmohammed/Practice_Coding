package car

import (
	"errors"
	"fmt"
)

type Car struct {
	Brand string
	Year  int
}

func Newcar() Car {
	return Car{
		Brand: "TOYOTA",
		Year:  2031,
	}
}
func (c Car) Display() (string, error) {

	if c.Year > 2030 {
		return "", errors.New("Wrong")
	} else {
		m := fmt.Sprintf("%s,%d", c.Brand, c.Year)
		return m, nil
	}

}
