package request

import (
	"context"
	"github.com/dotvezz/trading-badly/internal/messaging"
)

func Handle(req messaging.Request, resp *messaging.Response, args ...string) (err error) {
	hs := initEndpoints()

	ctx := context.Background()
	switch args[0] {
	case "chart":
		err = hs.chart(args[1:], ctx, req, resp)
	case "csv":
		err = hs.csv(args[1:], ctx, req, resp)
	case "watch":
	case "unwatch":
	case "help":
		resp.Body = "hmm"
	default:
		resp.Body = "Sorry, I didn't understand. Try `help` to see what I can do!"
	}

	if err != nil {
		resp.Body = "There was a problem, sorry!"
	}

	return err
}
