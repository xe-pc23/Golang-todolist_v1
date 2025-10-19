package main

import (
	"fmt"

	"local.package/main/functodo"
)

func main() {
	var choice int //使うToDoリストの選択（1 or 2)
	fmt.Printf("todolist_v1へようこそ\n")

	for {
		fmt.Println("制限付きToDoリスト(1)か自由ToDoリスト(2)か選んでください。")
		fmt.Scan(&choice)

		if choice == 1 || choice == 2 {
			break
		} else {
			fmt.Println("正しい値を入力してください")
		}
	} //配列かスライスを選択させる

	switch choice {
	case 1:
		fmt.Printf("制限付きToDoリストですね!タスク上限は5個までの固定長です。\n")
		functodo.Limittodo()
	case 2:
		fmt.Printf("自由 ToDoリストですね!タスク上限は無制限の可変長です。\n")
		functodo.Freetodo()
	} //選択完了時のメッセージ

}
