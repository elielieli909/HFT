package ftx_ws

import (
	"database/sql"
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func SubscribeOB(conn *sql.DB) {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// start receiving updates
	done := make(chan struct{})
	updates := make(chan OBData)
	go func() {
		defer close(done)
		defer close(updates)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			var obu OBUpdate
			json.Unmarshal(message, &obu)
			updates <- obu.Data
			// log.Printf("recv: %s", message)
			// log.Println(obu)

		}
	}()

	// Start storing updates in a binary file

	go func() {
		for {
			data, ok := <-updates
			if !ok {
				break
			}
			log.Printf("Dumping %d new updates", len(data.Bids)+len(data.Asks))
			dump(conn, data)
		}
	}()

	// Subscribe to OB
	request := make(map[string]string)
	request["op"] = "subscribe"
	request["channel"] = "orderbook"
	request["market"] = "BTC-PERP"

	err = c.WriteJSON(request)
	if err != nil {
		log.Println("write:", err)
		return
	}

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
