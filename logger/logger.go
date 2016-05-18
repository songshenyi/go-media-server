package logger

import (
	log "github.com/cihub/seelog"
)

func LaunchLog() error{
	launchLogConf :=`<seelog type="sync" minlevel="debug">
	<outputs  formatid="launchformat">
		<console />
		<rollingfile type="date" filename="logs/launch.log" datepattern="20060102" maxrolls="30"/>
	</outputs>
	<formats>
		<format id="launchformat" format="%Date(2006-01-02 15:04:05)%t%LEV%t[%File:%Line] [%FuncShort]%t%Msg%n"/>
	</formats>
	</seelog>`
	logger, err:= log.LoggerFromConfigAsString(launchLogConf)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	log.ReplaceLogger(logger)
	return nil
}

func InitLog(logConfigFile string) error{
	logger, err:= log.LoggerFromConfigAsFile(logConfigFile)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	log.ReplaceLogger(logger)
	return nil
}
