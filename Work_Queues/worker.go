package main

import (
	"bytes"
	"log"
	"time"

	"github.com/walter-arzapalo/ejercicio_rabbit/connection"
	"github.com/walter-arzapalo/ejercicio_rabbit/helpers"
)

/**
 * Main function of the Consumer
 * Return:
 *			- Send the message to a rabbitMQ queue
 * 			- Print the sent message
 */
func main() {
	// Makes and send the conecction or send the error to the failOnError function
	conn, err := connection.Connection()
	helpers.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Creates a server channel to process the pack of AMQP messages
	// Send the error if fail to the failOnError function
	ch, err := conn.Channel()
	helpers.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

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
	 * Consuming messages from a queue on a RabbitMQ
	 * Parameters:
	 *			- prefetch count: represents the maximum number of messages that the server will deliver to a consumer before it receives an acknowledgement
	 *			- prefetch size: represents the maximum number of bytes that the server will deliver to a consumer before it receives an acknowledgement
	 *			- global: indicates if the QoS settings should apply to the entire connection or just to the current channel
	 */
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	helpers.FailOnError(err, "Failed to set QoS")

	/**
	 * Consuming messages from a queue on a RabbitMQ
	 * Parameters:
	 *			- queue: the name of the queue (consumer tag)
	 *			- consumer: the name of the consumer, if is empty, the server should generate a consumer tag
	 *			- auto-ack: indicates if the messages should be acknowledged automatically or not (for assign the message to another consumer if the consumer disconnect)
	 *			- exclusive: indicates if the queue should be exclusive to the connection that declares it
	 *			- no-local: indicates if the consumer should receive messages that were published by the same connection or not
	 *			- no-wait: indicates if the method should return immediately or wait for the queue to be created
	 *			- arguments: if you want to pass aditional arguments to the queue
	 */
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	helpers.FailOnError(err, "Failed to register a consumer")

	// Declares the forever channel
	var forever chan struct{}

	// Starts a goroutine in loop for all the messages
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte(".")) // Variable that count the number of occurrences of a byte "." in the message
			t := time.Duration(dotCount)                 // Variable that contains the quantity of dots
			time.Sleep(t * time.Second)                  // Duration for processing every message
			log.Printf("Done")                           // Prints Done when the time finish
			d.Ack(false)                                 // Marks the current message like acknowledged
		}
	}()

	// Prints a message when the consumer is up
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	// Wait for the forever channel to close
	<-forever
}
