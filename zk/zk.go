package zk

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"time"
)

var zkClient *zk.Conn

// Connect ...
func Connect(uri string) error {
	c, _, err := zk.Connect([]string{uri}, time.Second*30, zk.WithLogInfo(false))
	if err != nil {
		fmt.Println("Error when connect to Zookeeper", uri, err)
		return err
	}

	fmt.Printf("⚡️[zookeeper]: connected to %s \n", uri)

	// Set client
	zkClient = c

	return nil
}
