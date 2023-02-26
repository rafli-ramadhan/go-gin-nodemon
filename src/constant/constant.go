package constant

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

const (
	DBServerMaster    = "master"
	DBServerSlave     = "slave"
	DBServerCMSMaster = "cms_master"
	DBServerCMSSlave  = "cms_slave"

	ContentTypeApplicationJson = "application/json"
	DOBLayout                  = `01/01/2023`
)

var (
	// error	
	ErrInvalidID            = errors.New("invalid id")
	ErrInvalidFormat        = errors.New("invalid format")
	ErrUserExist            = errors.New("user already exist")
	ErrUserNotRegistered    = errors.New("user not registered")

	// db connection
	_ = godotenv.Load()

	ServiceName = os.Getenv("SERVICE_NAME")

	DBPostgresConnMaxLifetime = os.Getenv("DB_POSTGRES_CONN_MAX_LIFETIME")
	DBPostgresConnMaxIdleTime = os.Getenv("DB_POSTGRES_CONN_MAX_IDLE_TIME")
	DBPostgresMaxOpenConns    = os.Getenv("DB_POSTGRES_MAX_OPEN_CONNS")
	DBPostgresMaxIdleConns    = os.Getenv("DB_POSTGRES_MAX_IDLE_CONNS")

	DBPostgresHostMaster = os.Getenv("DB_POSTGRES_HOST_MASTER")
	DBPostgresHostSlave  = os.Getenv("DB_POSTGRES_HOST_SLAVE")
	DBPostgresUsername   = os.Getenv("DB_POSTGRES_USERNAME")
	DBPostgresPassword   = os.Getenv("DB_POSTGRES_PASSWORD")
	DBPostgresPort       = os.Getenv("DB_POSTGRES_PORT")
	DBPostgresDatabase   = os.Getenv("DB_POSTGRES_DATABASE")

	DBPostgresHostCMSMaster = os.Getenv("DB_POSTGRES_HOST_CMS_MASTER")
	DBPostgresHostCMSSlave  = os.Getenv("DB_POSTGRES_HOST_CMS_SLAVE")
	DBPostgresUsernameCMS   = os.Getenv("DB_POSTGRES_USERNAME_CMS")
	DBPostgresPasswordCMS   = os.Getenv("DB_POSTGRES_PASSWORD_CMS")
	DBPostgresPortCMS       = os.Getenv("DB_POSTGRES_PORT_CMS")
	DBPostgresDatabaseCMS   = os.Getenv("DB_POSTGRES_DATABASE_CMS")
)
