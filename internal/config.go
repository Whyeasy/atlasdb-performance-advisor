package internal

//Config struct for holding config for exporter and Gitlab
type Config struct {
	LogFormat  string
	LogLevel   string
	ProjectID  string
	GroupID    string
	PublicKey  string
	PrivateKey string
	Since      int
}
