package functions

import "fmt"

type Server struct {
	config Config
}

type Config struct {
	listenAddr string
	id         int
	name       string
}

func NewServer(config Config) *Server {
	return &Server{config}
}

func NewConfig() Config {
	return Config{listenAddr: ":8000", id: 0, name: "default"}
}

func (c Config) WithListenAddr(addr string) Config {
	c.listenAddr = addr
	return c
}
func (c Config) WithID(id int) Config {
	c.id = id
	return c
}
func (c Config) WithName(name string) Config {
	c.name = name
	return c
}

func Example() {
	config := NewConfig().
		WithID(1).
		WithListenAddr(":5000").
		WithName("http")

	server := NewServer(config)

	fmt.Println(server)
}
