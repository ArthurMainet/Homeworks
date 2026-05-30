package config

type Config struct {
	Email    string
	Password string
	Address  string
}

func LoadConfig() *Config {
	return &Config{
		Email:    "testmail@gmail.com",
		Password: "123",
		Address:  "smtp.gmail.com:587",
	}
}
