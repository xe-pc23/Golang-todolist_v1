package functodo

import "fmt"

func Maptodo() {
	mapTask := make(map[string]string)
	var choicefunction string
	var key, value string

	fmt.Println("====マップ型ToDoリスト====")

	for {
		fmt.Println("1:タスク表示,2:タスク追加,3:タスク編集,4:タスク削除,5:終了")
		fmt.Print("選択: ")
		fmt.Scan(&choicefunction)

		switch choicefunction {

		case "1": //タスク表示
			fmt.Println("----「ToDoリスト」----")
			fmt.Println("タスク名 -> タスク内容")
			fmt.Println("----------------------")

			for k, v := range mapTask {
				fmt.Printf(" %s -> %s \n", k, v)
			}
			fmt.Printf("\n")

		case "2": //タスク追加
			fmt.Print("タスク名:")
			fmt.Scan(&key)
			fmt.Print("タスク内容:")
			fmt.Scan(&value)
			mapTask[key] = value
			fmt.Println("タスクを追加しました")
			fmt.Printf("\n")

		case "3": //タスク編集
			fmt.Print("編集する名：")
			fmt.Scan(&key)
			if _, ok := mapTask[key]; ok {
				fmt.Print("新しいタスク内容:")
				fmt.Scan(&value)
				mapTask[key] = value
				fmt.Println("タスクを更新しました")
			} else {
				fmt.Println("そのタスクは存在しません")
			}
			fmt.Printf("\n")

		case "4": //タスクを削除
			fmt.Print("削除するタスク名：")
			fmt.Scan(&key)
			if _, ok := mapTask[key]; ok {
				delete(mapTask, key)
				fmt.Println("タスクを削除しました")
			} else {
				fmt.Println("そのタスクは存在しません")
			}
			fmt.Printf("\n")

		case "5": //終了
			fmt.Println("終了します")
			return

		default:
			fmt.Println("1-5の数値を入力してください")
			fmt.Printf("\n")
		}

	}
}
