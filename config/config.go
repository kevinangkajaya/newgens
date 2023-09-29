package config

// Config store configuration value
type Config struct {
	Database struct {
		// Host is server host name (domain name) or ip where database is running
		Host string
		// Port is the database host port number
		Port uint16
		// UserName is database credential user name used to access db
		UserName string
		// Password is database credential password used to access db
		Password string
		// Name is the name of database
		Name string
	}
}
