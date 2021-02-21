package request

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/dotvezz/trading-badly/internal/messaging"
	"github.com/dotvezz/trading-badly/internal/stocks"
	"github.com/dotvezz/trading-badly/internal/stocks/chart"
	"time"
)

type endpoints struct {
	chart   messaging.Endpoint
	csv     messaging.Endpoint
	watch   messaging.Endpoint
	unWatch messaging.Endpoint
	help    messaging.Endpoint
}

func initEndpoints() endpoints {
	return endpoints{
		chart: func(args []string, ctx context.Context, request messaging.Request, response *messaging.Response) error {
			ticks, err := stocks.Get(
				args[0],
				time.Date(2021, 02, 16, 0, 30, 0, 0, time.Now().Location()),
				time.Date(2021, 02, 19, 24, 00, 0, 0, time.Now().Location()),
			)
			if err != nil {
				return err
			}
			response.Img, err = chart.FromTicks(ticks)
			return err
		},
		csv: func(args []string, ctx context.Context, request messaging.Request, response *messaging.Response) error {
			ticks, err := stocks.Get(
				args[0],
				time.Date(2021, 02, 16, 0, 30, 0, 0, time.Now().Location()),
				time.Date(2021, 02, 19, 24, 00, 0, 0, time.Now().Location()),
			)
			if err != nil {
				return err
			}

			buff := &bytes.Buffer{}
			file := csv.NewWriter(buff)

			for _, t := range ticks {
				file.Write([]string{t.Timestamp.Format(time.RFC1123), fmt.Sprint(t.Price), fmt.Sprint(t.Volume)})
			}

			response.TextFile = buff.Bytes()
			return nil
		},
	}
}
