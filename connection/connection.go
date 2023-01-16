package connection

import (
	"fmt"
	"io/ioutil"
	"log"

	amqp "github.com/rabbitmq/amqp091-go" // Implement the AMQP 0.9.1 protocol
	"gopkg.in/yaml.v2"
)

// Struct for the rabbit config
type Config struct {
	RabbitMQ struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Vhost    string `yaml:"vhost"`
	} `yaml:"rabbitmq"`
}

/**
 * Function to read the yml file
 * Parameters:
 *			- file: string with the file path of the yml file
 * Return:
 *			- &config: the address of the config variable
 *			- nil: a nil error to indicate that the function completed successfully
 */
func ReadConfig(file string) (*Config, error) {
	var config Config
	// Read the file or send an error
	source, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	// Parse the content of the file or send an error
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

/**
 * Makes the conecction
 * Return:
 *			- *amqp.Connection: represents a connection to a RabbitMQ server
 *			- error: the error of the connection
 */
func Connection() (*amqp.Connection, error) {
	// Read the file rabbit.yml
	config, err := ReadConfig("rabbit.yml")
	if err != nil {
		log.Fatal(err)
	}

	// String with the connection
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		config.RabbitMQ.Username,
		config.RabbitMQ.Password,
		config.RabbitMQ.Host,
		config.RabbitMQ.Port,
		config.RabbitMQ.Vhost)

	// Makes the connection
	conn, err := amqp.Dial(url)
	return conn, err
}
