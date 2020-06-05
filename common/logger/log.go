package logger

import (
	error2 "github.com/ciazhar/golang-cqrs/common/error"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func InitLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.CallerSkipFrameCount = 3
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		log.Logger = log.Output(&lumberjack.Logger{
			Filename: "logs/app.log",
			MaxSize:  100,
			Compress: true,
			MaxAge:   28,
		})
	}
}

func Infof() *zerolog.Event {
	return log.Info().Caller()
}

func Errorf() *zerolog.Event {
	return log.Error().Caller()
}

func Info(msg string, v ...interface{}) {
	log.Info().Caller().Msgf(msg, v)
}

func Warn(msg string, err error) {
	errorString := msg + err.Error()
	log.Warn().Caller().Msg(errorString)
}

func Error(msg string, err error) {
	errorString := msg + err.Error()
	log.Error().Caller().Msg(errorString)
}

func ErrorS(msg error) error2.Error {
	log.Error().Caller().Msg(msg.Error())
	return error2.New(msg)
}

func WarnS(msg error) error2.Error {
	log.Warn().Caller().Msg(msg.Error())
	return error2.New(msg)
}

func WithError(err error) error {
	log.Error().Caller().Msg(err.Error())
	return err
}
