package postgres

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_errors "github.com/sinameshkini/auth-service/pkg/errors"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	TimeZone string
	Debug    bool
}

func (c *Config) IsValid() bool {
	return c != nil &&
		c.Host != "" &&
		c.Port != "" &&
		c.Username != "" &&
		c.Password != "" &&
		c.Database != "" &&
		c.TimeZone != ""
}

func Connect(config *Config) (db *gorm.DB, err error) {
	if !config.IsValid() {
		return nil, _errors.ErrInvalidDatabaseConfig
	}

	var newLogger logger.Interface
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s",
		config.Host,
		config.Username,
		config.Password,
		config.Database,
		config.Port,
		config.TimeZone,
	)

	if config.Debug {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // Disable color
			},
		)
	}

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
}
