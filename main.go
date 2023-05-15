package main

import (
	"fmt"

	"github.com/joho/godotenv"

	"github.com/shin-hama/tascker/redmine"
)

func loadEnv() {
	err := godotenv.Load(".env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
		panic(err)
	}
}

func main() {
	loadEnv()

	redmine.Feed()
}
