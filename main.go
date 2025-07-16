package main

import "tgbot/rabbitmq"

func main() {
	rabbitmq.Send()
	rabbitmq.Receive()
}
