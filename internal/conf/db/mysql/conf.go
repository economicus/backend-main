package mysql

import (
	"fmt"
	"log"
	"os"
)

type conf struct {
	dial string
	user string
	pwd  string
	host string
	port string
	name string
}

func newConf() conf {
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.SetPrefix("[WARNING] ")

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	if host == "" {
		log.Println("MISSING DATABASE ENV: empty host\nChange to default host mysql")
		host = "localhost"
	}
	if port == "" {
		log.Println("MISSING DATABASE ENV: empty port\nChange to default port 3306")
		port = "3306"
	}
	if user == "" {
		log.Println("MISSING DATABASE ENV: empty user\nChange to default user root")
		user = "root"
	}
	if pwd == "" {
		log.Println("MISSING DATABASE ENV: empty dial\nChange to default password pwd")
		pwd = "4406"
	}
	if name == "" {
		log.Println("MISSING DATABASE ENV: empty dial\nChange to default name economicus")
		name = "economicus"
	}
	return conf{
		dial: "mysql",
		user: user,
		pwd:  pwd,
		host: host,
		port: port,
		name: name,
	}
}

func (c conf) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.user, c.pwd, c.host, c.port, c.name)
}

func (c conf) Info() {
	fmt.Println("========== DB ==========")
	fmt.Println("Dial: ", c.dial)
	fmt.Println("User: ", c.user)
	fmt.Println("Password: ", c.pwd)
	fmt.Println("Host: ", c.host)
	fmt.Println("Port: ", c.port)
	fmt.Println("Name: ", c.name)
}
