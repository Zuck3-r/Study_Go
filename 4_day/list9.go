package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	logrus.Info("succeeded")
	logrus.Warn("not correct")
	logrus.Error("something error")
	logrus.Fatal("panic")
}
