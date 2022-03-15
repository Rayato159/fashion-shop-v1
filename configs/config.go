package configs

type Config struct {
	Database Database
}

type Database struct {
	Url string
}
