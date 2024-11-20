package configs

import (
	"fmt"
	"runtime"
	"strings"
	"whatisyourtime/internal/constants"

	"github.com/sirupsen/logrus"
)

func initLogger() {
	logrus.SetFormatter(
		&logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			DisableColors:   false,
			TimestampFormat: constants.TimeLayout,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcname := s[len(s)-1]
				return fmt.Sprintf("%s()", funcname), fmt.Sprintf("%s:%d", f.File, f.Line)
			},
		},
	)
}
