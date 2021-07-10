package artloghandler


import (
"fmt"
logger "github.com/apresfiux/artlogger"
"os"
)

// LogError log errors
func LogError(err error, printErr bool, Critical bool) {
	if err != nil {
		logger.LogError.Printf("%s", err)
		if printErr {
			fmt.Println("Error:", err)
		}
		if Critical {
			os.Exit(1)
		}
	}
}

// AppLog log application logs
func AppLog(msg string) {
	logger.AppLog.Printf(msg)
}

// QueueLog log queue events
func QueueLog(msg string) {
	logger.QueueLog.Printf(msg)
}

// AccessLog - log access requests
func AccessLog(msg string) {
	logger.AccessLog.Printf(msg)
}