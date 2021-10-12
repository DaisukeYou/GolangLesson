package main
// 「GO111MODULE=off」とターミナルに記述すると実行できる
import (
	"fmt"

	"./animals"
)

func main() {
	fmt.Println(AppName()) /*関数AppNameの呼び出しを追加*/

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}