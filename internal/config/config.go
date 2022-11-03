package config

var Conf Config

type Config struct {
	Secret secret
}

type secret struct {
	URL      string `env:"URL" envDefault:"https://example.com"`
	UserID   string `env:"UserID" envDefault:"1111111"`
	Password string `env:"Password" envDefault:"your_nice_password"`
	Token    string `env:"Token" envDefault:"your_nice_slack_token"`
	Channel  string `env:"Channel" envDefault:"your_nice_slack_channel"`
}
