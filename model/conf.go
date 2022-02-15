package model

//App app.conf struct
type App struct {
	App struct {
		APIPort, Host             string
		SSL, DuplicateURL         bool
		IssueRateLimit, MurmurBit int16
		LogLevel                  string
	}
	MySQL struct {
		HostPort, User, Password, Database string
	}
	Redis struct {
		HostPort, Password string
	}
}