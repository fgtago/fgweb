package appsmodel

type Configuration struct {
	Port     int `yaml:"port"`
	Database struct {
		Name     string `yaml:"name"`
		Server   string `yaml:"server"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	SecureCookie bool `yaml:"securecookie"`
	Cookie       struct {
		Secure   string `yaml:"secure"`
		LifeTime int    `yaml:"lifetime"`
	} `yaml:"cookie"`
	Template struct {
		Cached bool   `yaml:"cached"`
		Dir    string `yaml:"dir"`
	} `yaml:"template"`
	Application struct {
		PageDir    string `yaml:"pagedir"`
		ContentDir string `yaml:"contentdir"`
	} `yaml:"application"`
	Logging struct {
		Enabled bool `yaml:"enabled"`
	} `yaml:"logging"`
}
