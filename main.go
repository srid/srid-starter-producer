package main

import (
	"bufio"
	"encoding/json"
	"github.com/bmizerany/lpx"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func makeGinAccounts() gin.Accounts {
	return gin.Accounts{
		config.Username: config.Password,
	}
}

func main() {
	router := gin.Default()
	routerAuthed := router.Group("/", gin.BasicAuth(makeGinAccounts()))
	routerAuthed.POST("/logs", logsReceived)
	log.Printf("Running app at %v\b", config.Port)
	router.Run(":" + config.Port)
}

func logsReceived(c *gin.Context) {
	defer c.Request.Body.Close()
	err := handleLogplexRequest(c.Request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{})
	}
}

// 70 <174>1 2012-07-22T00:06:26+00:00 host erlang console - Hi from erlang
func handleLogplexRequest(r *http.Request) error {
	// TODO: handle frame-id to handle duplicates
	var records []Record
	lp := lpx.NewReader(bufio.NewReader(r.Body))

	for lp.Next() {
		header := lp.Header()
		data := lp.Bytes()
		json, err := encodeLogplexRecord(header, data)
		if err != nil {
			// TODO: don't abort, but report and continue with other records
			return err
		}
		// Application name (header.Name) should be unique.
		record := Record{partitionKey: string(header.Name), data: json}
		records = append(records, record)
	}
	return putRecords(records)
}

type LogplexLogLine struct {
	Header *lpx.Header
	Data   []byte
}

func encodeLogplexRecord(header *lpx.Header, data []byte) ([]byte, error) {
	info := LogplexLogLine{header, data}
	return json.Marshal(info)
}
