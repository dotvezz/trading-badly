package messaging

import (
	"context"
	"image"
)

type Endpoint func(args []string, ctx context.Context, request Request, response *Response) error

type Request struct {
	Author  string
	Channel string
}

type Response struct {
	Img  image.Image
	Body string
}
