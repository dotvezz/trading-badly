package request

import "github.com/dotvezz/trading-badly/internal/messaging"

type endpoints struct {
	chart   messaging.Endpoint
	watch   messaging.Endpoint
	unWatch messaging.Endpoint
	help    messaging.Endpoint
}

func initEndpoints() endpoints {
	return endpoints{}
}
