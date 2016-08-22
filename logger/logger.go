package logger

import (
	"github.com/cihub/seelog"
)

func InitLaunchLog() error {
	launchLogConf := `<seelog type="sync" minlevel="debug">
	<outputs  formatid="launchformat">
		<console />
		<rollingfile type="date" filename="logs/launch.log" datepattern="20060102" maxrolls="30"/>
	</outputs>
	<formats>
		<format id="launchformat" format="%Date(2006-01-02 15:04:05)%t%LEV%t[%File:%Line] [%FuncShort]%t%Msg%n"/>
	</formats>
	</seelog>`
	logger, err := seelog.LoggerFromConfigAsString(launchLogConf)

	if err != nil {
		seelog.Error(err.Error())
		return err
	}

	logger.SetAdditionalStackDepth(1)
	seelog.ReplaceLogger(logger)
	return nil
}

func InitLog(logConfigFile string) error {
	logger, err := seelog.LoggerFromConfigAsFile(logConfigFile)
	if err != nil {
		seelog.Error(err.Error())
		return err
	}

	logger.SetAdditionalStackDepth(1)
	seelog.ReplaceLogger(logger)
	return nil
}

func Tracef(format string, params ...interface{}) {
	seelog.Tracef(format, params...)
}

func Debugf(format string, params ...interface{}) {
	seelog.Debugf(format, params...)
}

func Infof(format string, params ...interface{}) {
	seelog.Infof(format, params...)
}

func Warnf(format string, params ...interface{}) error {
	return seelog.Warnf(format, params...)
}

func Errorf(format string, params ...interface{}) error {
	return seelog.Errorf(format, params...)
}

func Criticalf(format string, params ...interface{}) error {
	return seelog.Criticalf(format, params...)
}

func Trace(v ...interface{}) {
	seelog.Trace(v)
}

func Debug(v ...interface{}) {
	seelog.Debug(v)
}

func Info(v ...interface{}) {
	seelog.Info(v)
}

func Warn(v ...interface{}) error {
	return seelog.Warn(v)
}

func Error(v ...interface{}) error {
	return seelog.Error(v)
}

func Critical(v ...interface{}) error {
	return seelog.Critical(v)
}
