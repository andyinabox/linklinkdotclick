package mailservice

type Config struct {
	SmtpAddr string
}

type Service struct {
	conf *Config
}

func New(conf *Config) *Service {
	return &Service{conf}
}
