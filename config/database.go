package config

type DBConfig struct {
	host     string
	port     string
	database string
	username string
	password string
	charset  string
}

var DBCfg DBConfig

func init() {
	DBCfg = DBConfig{
		host:     "localhost",
		port:     "3306",
		database: "yb_stdata",
		username: "root",
		password: "hzkjzyjsxy",
		charset:  "utf8mb4",
	}
}

func (dbc DBConfig) GetHost() string {
	return dbc.host
}

func (dbc DBConfig) GetPort() string {
	return dbc.port
}

func (dbc DBConfig) GetDatabase() string {
	return dbc.database
}

func (dbc DBConfig) GetUsername() string {
	return dbc.username
}

func (dbc DBConfig) GetPassword() string {
	return dbc.password
}

func (dbc DBConfig) GetCharset() string {
	return dbc.charset
}
