package stocks

import (
	"time"

	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
)

type Tick struct {
	Timestamp time.Time
	Volume    int
	Price     float64
}

func Get(ticker string, from, to time.Time) ([]Tick, error) {
	itr := chart.Get(&chart.Params{
		Symbol:     ticker,
		Start:      datetime.New(&from),
		End:        datetime.New(&to),
		Interval:   datetime.FifteenMins,
		IncludeExt: false,
	})
	if itr.Err() != nil {
		return nil, itr.Err()
	}

	ticks := make([]Tick, itr.Count())
	for i := 0; itr.Next(); i++ {
		bar := itr.Bar()
		price, _ := bar.Open.Float64()
		ticks[i] = Tick{
			Timestamp: time.Unix(int64(bar.Timestamp), 0),
			Volume:    bar.Volume,
			Price:     price,
		}
	}

	return ticks, nil
}
