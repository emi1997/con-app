package client

import (
	// "log"

	"context"
	// "encoding/json"
	// "io/ioutil"
	// "os"

	//"github.com/emi1997/con-app/scanner"

	//"net/http"
	// "fmt"
	// "log"
	"time"

	"github.com/olivere/elastic"
	//"github.com/spf13/cobra"
)

//Client is being called by the main.go every time the app is being started
var Client, err = elastic.NewClient(
	elastic.SetURL("http://localhost:9200"),
	elastic.SetHealthcheckInterval(10*time.Second),
)

var ctx = context.Background()

var bulkRequest = Client.Bulk()

//TestEntry represents a document in the index
type TestEntry struct {
	Fach  string `json:"Fach"`
	Stoff string `json:"Thema"`
	Wann  string `json:"Wann"`
}

//TestUpdate represents a document in the index
type TestUpdate struct {
	Fach  string `json:"Fach"`
	Stoff string `json:"Thema"`
	Wann  string `json:"Wann"`
}
