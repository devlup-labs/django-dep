package config

import (
	"sync"
)

type config struct {
	logLevel   string
	debugMode  bool
	port       int
	token      string
	scriptPath string
}

var instance *config
var once sync.Once

func getConfig() *config {
	once.Do(func() {
		instance = newConfig()
	})
	return instance
}

func newConfig() *config {
	v := Viper()
	port := v.GetInt("port")
	if port == 0 {
		port = 3000 // Fallback to default port
	}

	sp := v.GetString("scriptPath")
	if sp == "" {
		sp = "./script.sh"
	}

	return &config{
		logLevel:   v.GetString("log.level"),
		debugMode:  v.GetBool("debug"),
		port:       port,
		token:      v.GetString("token"),
		scriptPath: v.GetString("scriptPath"),
	}
}

// Update creates a new instance of the configuration and reads all values again
func Update() {
	instance = newConfig()
}

// LogLevel returns the log level for logger from the environment
func LogLevel() string {
	return getConfig().logLevel
}

// DebugModeEnabled returns the debug mode bool from the environment
func DebugModeEnabled() bool {
	return getConfig().debugMode
}

// Port returns the application port to be used from the environment
func Port() int {
	return getConfig().port
}

// Token returns the validation token to authorize the CI
func Token() string {
	return getConfig().token
}

// ScriptPath returns the path to the script which will be executed
func ScriptPath() string {
	return getConfig().scriptPath
}
