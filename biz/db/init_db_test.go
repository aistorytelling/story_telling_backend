package db

import (
	"fmt"
	"testing"
)

func TestConnectDB(t *testing.T) {
	got, err := ConnectDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(got.)
}
