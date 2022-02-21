package main

import (
	_ "context"
	_ "google.golang.org/grpc"
	_ "log"
	_ "net"
)

func main() {
	cards := newDeck()
	cards.print()
}
