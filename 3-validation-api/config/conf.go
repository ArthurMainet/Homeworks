package config

type Config struct {
	Email    string
	Password string
	Adress   string
}

func LoadConfig() *Config {
	return &Config{
		Email:    "testmail@gmail.com",
		Password: "123",
		Adress:   "smtp.gmail.com:587",
	}
}
