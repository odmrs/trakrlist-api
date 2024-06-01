package config

var logger *Logger

func Init() error {
	err := InitializePostgres()
	if err != nil {
		logger.Errorf("initializePostgres error: %v", err)
		return err
	}

	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
