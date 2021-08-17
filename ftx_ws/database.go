package ftx_ws

import (
	"database/sql"
)

func dump(conn *sql.DB, data OBData) {
	// Need to split up the update by price
	for _, bid := range data.Bids {
		var action bool
		if data.Action == "update" {
			action = true
		} else {
			action = false
		}
		conn.Exec("INSERT INTO updates VALUES ($1, $2, $3, $4, $5)", data.Time, true, bid[0], bid[1], action)
	}
	for _, ask := range data.Asks {
		var action bool
		if data.Action == "update" {
			action = true
		} else {
			action = false
		}
		conn.Exec("INSERT INTO updates VALUES ($1, $2, $3, $4, $5)", data.Time, false, ask[0], ask[1], action)
	}
}

// func dump(data OBData) {
// 	f, err := os.OpenFile("./data/datadump.bin", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
// 	if err != nil {
// 		log.Fatal("Couldn't open file")
// 	}
// 	defer f.Close()

// 	// Need to split up the update by price
// 	for _, bid := range data.Bids {
// 		var action bool
// 		if data.Action == "update" {
// 			action = true
// 		} else {
// 			action = false
// 		}
// 		d := DeconstructedOBData{data.Time, true, bid[0], bid[1], action}
// 		err = binary.Write(f, binary.LittleEndian, d)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	for _, ask := range data.Asks {
// 		var action bool
// 		if data.Action == "update" {
// 			action = true
// 		} else {
// 			action = false
// 		}
// 		d := DeconstructedOBData{data.Time, true, ask[0], ask[1], action}
// 		err = binary.Write(f, binary.LittleEndian, d)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }
