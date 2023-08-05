package db

import (
	"fmt"
	"testing"
)

func TestConnectDB(t *testing.T) {
	got, err := connectDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(got)
}

func Test_connectMongoDB(t *testing.T) {
	connectMongoDB()
}
