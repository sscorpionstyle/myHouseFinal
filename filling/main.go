package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

//func main() {
//	runProject()
//}

func main() {
	urlExample := "postgres://house_new:123@localhost:5436/test_db1"
	_, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unableeeee to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Fprintf(os.Stderr, "ABLE to connect to database: %v\n", err)
		fmt.Println("ready")
	}

	runProject()
}
