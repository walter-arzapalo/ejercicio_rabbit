package helpers

import "log"

/**
 * Function to handle errors on fail
 * Parameters:
 * 			- err: the error
 *			-	msg: the string personalized message
 * Return:
 * 			- Print the personalized message and error
 */
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
