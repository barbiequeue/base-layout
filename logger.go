package main

import "github.com/sirupsen/logrus"

var log = logrus.New()

// level string possible values:
//
//	panic • fatal • error • warn | warning • info • debug • trace
func initLogger(levelString string) {

	log.Info("Start logger initialization")
	lvl, err := logrus.ParseLevel(levelString)
	if err != nil {
		log.Errorf("Failed to parse log level `%s`. Will use `debug` level.", levelString)
		lvl = logrus.DebugLevel
	}
	log.SetLevel(lvl)
	log.Infof("Log init complete. Use level: %s", lvl)
}
