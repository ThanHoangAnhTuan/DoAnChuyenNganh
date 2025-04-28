package setting

type Config struct {
	Server   ServerSetting `mapstructure:"server"`
	Logger   LoggerSetting `mapstructure:"logger"`
	Mysql    MysqlSetting  `mapstructure:"mysql"`
	Redis    RedisSetting  `mapstructure:"redis"`
	Kafka    KafkaSetting  `mapstructure:"kafka"`
	Email    EmailSetting  `mapstructure:"email"`
	JWT      JWTSetting    `mapstructure:"jwt"`
	Frontend Frontend      `mapstructure:"frontend"`
}

type Frontend struct {
	Url string `mapstructure:"url"`
}

type ServerSetting struct {
	Port       int    `mapstructure:"port"`
	Mode       string `mapstructure:"mode"`
	DriverName string `mapstructure:"driverName"`
}

type LoggerSetting struct {
	LogLevel   string `mapstructure:"logLevel"`
	Filename   string `mapstructure:"fileName"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
	Compress   bool   `mapstructure:"compress"`
}

type MysqlSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DatabaseName    string `mapstructure:"databaseName"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
	PoolSize int    `mapstructure:"poolSize"`
}

type EmailSetting struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type KafkaSetting struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JWTSetting struct {
	Token_Hour_Lifespan string `mapstructure:"TOKEN_HOUR_LIFESPAN"`
	JWT_Expiration      string `mapstructure:"JWT_EXPIRATION"`
	Api_secret          string `mapstructure:"API_SECRET"`
}
