package app

import (
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func SetupLogger(debug bool) {
	if debug {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	log.SetFormatter(&log.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return fmt.Sprintf(" %s() ->", f.Func.Name()), ""
		},
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
		PadLevelText:    true,
	})

}
