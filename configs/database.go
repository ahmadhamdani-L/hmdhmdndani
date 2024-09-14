package configs

import (
	"strconv"
	"strings"
	"sync"
	"time"
)

type DbConfig struct {
	host         string
	port         string
	name         string
	user         string
	pass         string
	sslmode      string
	tz           string
	maxOpenConns int
	maxIdleConns int
	connLifetime time.Duration
}

var (
	db     *DbConfig
	dbOnce sync.Once
)

func (dbc *DbConfig) Host() string {
	return dbc.host
}

func (dbc *DbConfig) Port() string {
	return dbc.port
}

func (dbc *DbConfig) Name() string {
	return dbc.name
}

func (dbc *DbConfig) User() string {
	return dbc.user
}

func (dbc *DbConfig) Password() string {
	return dbc.pass
}

func (dbc *DbConfig) SslMode() string {
	return dbc.sslmode
}

func (dbc *DbConfig) Timezone() string {
	return dbc.tz
}

func (dbc *DbConfig) MaxOpenConnections() int {
	return dbc.maxOpenConns
}

func (dbc *DbConfig) MaxIdleConnections() int {
	return dbc.maxIdleConns
}

func (dbc *DbConfig) ConnectionLifetime() time.Duration {
	return dbc.connLifetime
}

func DB() *DbConfig {
	dbOnce.Do(func() {
		listSSLMode := map[string]bool{
			"disable":     true,
			"allow":       true,
			"prefer":      true,
			"require":     true,
			"verify-ca":   true,
			"verify-full": true,
		}
		givenSSLMode := strings.ToLower(strings.TrimSpace("disable"))
		if _, ok := listSSLMode[givenSSLMode]; !ok {
			givenSSLMode = "disable"
		}

		mOpenConns, err := strconv.Atoi("20")
		if err != nil {
			mOpenConns = 0
		}

		mIdleConns, err := strconv.Atoi("18")
		if err != nil {
			mIdleConns = 18
		}

		ltConn, err := time.ParseDuration("0s")
		if err != nil {
			ltConn = 0 * time.Second
		}

		db = &DbConfig{
			host:         "postgres.railway.internal",
			port:         "5432",
			name:         "railway",
			user:         "postgres",
			pass:         "RPpXCbONvVntIDzVFdEHRlHjWlSAvHRm", // Hardcoded password
			sslmode:      givenSSLMode,
			tz:           "UTC",
			maxOpenConns: mOpenConns,
			maxIdleConns: mIdleConns,
			connLifetime: ltConn,
		}
	})

	return db
}
