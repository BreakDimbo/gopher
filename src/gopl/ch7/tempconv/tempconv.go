package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64

type Fahrenheit float64

type celsiusFlag struct{ Celsius }

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (c *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		c.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		c.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("ivalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
