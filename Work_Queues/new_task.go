package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go" // Implement the AMQP 0.9.1 protocol
	"github.com/walter-arzapalo/ejercicio_rabbit/connection"
	"github.com/walter-arzapalo/ejercicio_rabbit/helpers"
)

/**
 * Main function of the Producer
 * Return:
 *			- Send the message to a rabbitMQ queue
 * 			- Print the sent message
 */
func main() {
	// Makes and send the conecction or send the error to the failOnError function
	conn, err := connection.Connection()
	helpers.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close() // Wait to finish all the function to close the connection

	// Creates a server channel to process the pack of AMQP messages
	// Send the error if fail to the failOnError function
	ch, err := conn.Channel()
	helpers.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Scanning the input by the user
	var qtyMessage, qtyDots int
	fmt.Println("Please enter the number of messages you wanna send: ")
	fmt.Scanln(&qtyMessage)
	qtyMessage = validateNumber(qtyMessage)
	fmt.Scanln(qtyMessage)

	fmt.Println("Please enter the seconds you wanna wait for every sent message: ")
	fmt.Scanln(&qtyDots)
	qtyDots = validateNumber(qtyDots)
	fmt.Scanln(qtyDots)

	// For loop to concatenate the dots after the message
	dots := "."
	for d := 1; d < qtyDots; d++ {
		dots += "."
	}

	// For loop for sending many messages as the user inputted
	for i := 1; i <= qtyMessage; i++ {
		/**
		 * Declares a new queue on a RabbitMQ server or connect to existing
		 * Parameters:
		 * 			- name: the name of the queue
		 *			- durable: indicate if queue will survive a server restart
		 *			- delete when unused: indicates if the queue should be deleted when it is no longer in use
		 *			- exclusive: indicates if the queue should be exclusive to the connection that declares it
		 *			- no-wait: indicates if the method should return immediately or wait for the queue to be created
		 *			- arguments: if you want to pass aditional arguments to the queue
		 * Return:
		 *			- Queue struct which contains the stated properties
		 *			- Error if fail to the failOnError function
		 */
		q, err := ch.QueueDeclare(
			"walter", // name
			true,     // durable
			false,    // delete when unused
			false,    // exclusive
			false,    // no-wait
			nil,      // arguments
		)
		helpers.FailOnError(err, "Failed to declare a queue")

		/**
		 * Creates a new context to handle a deadline
		 * Parameters:
		 *			- context.Background(): creates a new background context that is used as the parent context for the new context with a timeout
		 *			- time.Second: the duration for the timeout
		 * Return:
		 *			- ctx: the new context with the timeout
		 *			- cancel: a function that can be used to cancel the context
		 */
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel() // Cancel the context when finish the function

		// Set the body of the message from bodyFrom function
		body := bodyFrom(os.Args, i, dots)

		/**
		 * Publish a message to a queue of rabbitMQ
		 * Parameters:
		 * 			- ctx: the context
		 *			- exchange: the name of the exchange to route the message to the queue
		 *			- routing key: the name of the queue used like key to route the message
		 *			- mandatory: declares if a message must be delivered to a queue, or an error will be returned
		 *			- inmediate: indicates if a message send immediate if exist a consumer or return to the sender
		 *			- amqp.Publishing: is a struct which contains the properties of the message
		 *									- DeliveryMode: indicates if the message will survive a server restart
		 *									- ContentType: sets the type of the content for the sent messages
		 *									- Body: the message as a byte array
		 * Return:
		 *			- Error which be nil if the message was published successfully
		 *			- Otherwise contains the error message and send to failOnError function
		 */
		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(body),
			})
		helpers.FailOnError(err, "Failed to publish a message")

		// Prints the sent message
		log.Printf(" [x] Sent %s\n", body)
	}
}

/**
 * Generate the message that will be send
 * Parameters:
 *      - args: is a slice of strings
 *      - i: the number of the messages
 *      - dots: the dots that will be added after every message
 * Return:
 *      - s: the message string
 */
func bodyFrom(args []string, i int, dots string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = fmt.Sprintf("Message N°%v%s", i, dots)
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

/**
 * Validate the inputs of the user (should be positive number)
 * Parameters:
 *      - num: the input
 * Return:
 *      - num: the number validate
 */
func validateNumber(num int) int {
	for num <= 0 {
		fmt.Println("Please enter a valid number: ")
		fmt.Scanln(&num)
	}
	return num
}
