package utils

import (
	"os"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() error {
	if Log == nil {

		pathMap := lfshook.PathMap{
			logrus.InfoLevel:  "./info.log",
			logrus.ErrorLevel: "./info.log",
		}
		Log = logrus.New()
		Log.Hooks.Add(lfshook.NewHook(
			pathMap,
			&logrus.JSONFormatter{},
		))

		Log.SetFormatter(&logrus.JSONFormatter{})

		// Output to stdout instead of the default stderr
		// Can be any io.Writer, see below for File example
		Log.SetOutput(os.Stdout)

		//logrus.Fields

		// Only log the warning severity or above.
		//Log.SetLevel(log.WarnLevel)

	}

	return nil
}
