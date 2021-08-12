package ftx_ws

import "flag"

var addr = flag.String("addr", "ftx.com", "ftx ws address")

type OBData struct {
	Time     float32     `json:"time"`
	Checksum int         `json:"checksum"`
	Bids     [][]float32 `json:"bids"`
	Asks     [][]float32 `json:"asks"`
	Action   string      `json:"action"`
}

type OBUpdate struct {
	Channel string `json:"channel"`
	Market  string `json:"market"`
	Type    string `json:"type"`
	Data    OBData `json:"data"`
}

type TradeData struct {
	Id          int     `json:"id"`
	Price       float32 `json:"price"`
	Size        float32 `json:"size"`
	TakerSide   string  `json:"side"`
	Liquidation bool    `json:"liquidation"`
	Time        string  `json:"time"`
}

type TradeUpdate struct {
	Channel string      `json:"channel"`
	Market  string      `json:"market"`
	Type    string      `json:"type"`
	Data    []TradeData `json:"data"`
}
