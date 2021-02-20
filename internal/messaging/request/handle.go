package request

import (
	"context"
	"github.com/dotvezz/trading-badly/internal/messaging"
	"image"
)

func Handle(req messaging.Request, resp *messaging.Response, args ...string) {
	hs := initEndpoints()
	var err error

	ctx := context.Background()
	switch args[0] {
	case "chart":
		err = hs.chart(args, ctx, req, resp)
	case "watch":
	case "unwatch":
	case "help":
		resp.Body = "hmm"
	case "img":
		resp.Img = image.NewRGBA(image.Rectangle{
			Max: image.Point{
				X: 100,
				Y: 600,
			},
		})
		resp.Body = "Here's your image"
	default:
		resp.Body = "Sorry, I didn't understand. Try `help` to see what I can do!"
	}

	if err != nil {
		resp.Body = "There was a problem, sorry!"
	}
}
