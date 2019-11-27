package godao

import "fmt"

// DBConfig database config interface
type DBConfig interface {
	GetDriver() string
	GetDSN() string
	GetShowSQL() bool
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
}

var (
	_ DBConfig = PostgreSQLConfig{}
)

// PostgreSQLConfig postgresql database config
type PostgreSQLConfig struct {
	Host              string `yaml:"host" json:"host"`
	Port              int    `yaml:"port" json:"port"`
	User              string `yaml:"user" json:"user"`
	Password          string `yaml:"password" json:"password"`
	Database          string `yaml:"database" json:"database"`
	SSLMode           string `yaml:"sslmode" json:"sslmode"`
	ShowSQL           bool   `yaml:"showsql" json:"showsql"`
	MaxIdleConnection int    `yaml:"maxidleconnection" json:"maxidleconnection"`
	MaxOpenConnection int    `yaml:"maxopenconnection" json:"maxopenconnection"`
	ApplicationName   string `yaml:"applicationname" json:"applicationname"`
}

func (c PostgreSQLConfig) GetDriver() string {
	return "postgres"
}

func (c PostgreSQLConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=%s password=%s fallback_application_name=%s",
		c.Host, c.Port, c.User, c.Database, c.SSLMode, c.Password, c.ApplicationName,
	)
}

func (c PostgreSQLConfig) GetShowSQL() bool {
	return c.ShowSQL
}

func (c PostgreSQLConfig) GetMaxIdleConnection() int {
	return c.MaxIdleConnection
}

func (c PostgreSQLConfig) GetMaxOpenConnection() int {
	return c.MaxOpenConnection
}

// MemorySqliteConfig in memory sqlite3 config
type MemorySqliteConfig struct {
}

func (c MemorySqliteConfig) GetDriver() string {
	return "sqlite3"
}

func (c MemorySqliteConfig) GetDSN() string {
	return ":memory:"
}

func (c MemorySqliteConfig) GetShowSQL() bool {
	return true
}

func (c MemorySqliteConfig) GetMaxIdleConnection() int {
	return 1
}

func (c MemorySqliteConfig) GetMaxOpenConnection() int {
	return 1
}

type Config struct {
	Dialect     string
	URL         string
	ShowSQL     bool
	MaxIdleConn int
	MaxOpenConn int
}
