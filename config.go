package main

type configuration struct {
	logLevel configParameter
}

func (c configuration) loadFromEnv() error {

	if err := setFromEnv(&conf.logLevel, "LOG_LEVEL", false, true); err != nil {
		return err
	}

	return nil
}

var conf configuration
