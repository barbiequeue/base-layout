package main

import (
	"fmt"
	"os"
	"strconv"
)

type configParameter string

func (cp configParameter) Int() int {
	n, err := strconv.ParseInt(string(cp), 10, 32)
	if err != nil {
		return 0
	}

	return int(n)
}

func (cp configParameter) Bool() bool {
	t, err := strconv.ParseBool(string(cp))
	if err != nil {
		return false
	}

	return t
}

func (cp configParameter) Float64() float64 {
	f, err := strconv.ParseFloat(string(cp), 64)
	if err != nil {
		return 0.0
	}

	return f
}

func (cp configParameter) String() string {
	return string(cp)
}

func setFromEnv(v *configParameter, env string, required bool, warn bool) error {

	ev, exists := os.LookupEnv(env)
	if !exists {
		if required {
			return fmt.Errorf("conf error: failed to load required env variable '%s'", env)
		}
		if warn {
			log.Errorf("conf warn: failed to load env variable '%s'", env)
		}
	}

	*v = configParameter(ev)

	return nil
}
