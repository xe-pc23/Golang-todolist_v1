package functodo

import "fmt"

func Freetodo() {
	freeTask := make([]string, 0, 3) //0をなくして3にすると０が3個入って4個目にタスクが生成される
	var choicefunction string
	var input_task string
	var index int

	fmt.Println("====可変長配列ToDoリスト====")

	for {
		fmt.Printf("1:タスク表示,2:タスク追加,3:タスク編集,4:タスク削除,5:終了\n")
		fmt.Printf("選択:")

		fmt.Scan(&choicefunction)
		switch choicefunction {

		case "1": //タスク表示
			fmt.Printf("--「ToDoリスト」--\n")
			for i := 0; i < len(freeTask); i++ {
				fmt.Print(i, ":")
				fmt.Println(freeTask[i])
			}
			fmt.Printf("\n キャパシティは：%d ", cap(freeTask))
			fmt.Printf("\n")

		case "2": //タスク追加
			fmt.Print("タスクを入力: ")
			fmt.Scan(&input_task)
			freeTask = append(freeTask, input_task)
			fmt.Printf("タスクを配列%d番目に追加しました\n ", index)
			fmt.Printf("\n")

		case "3": //タスク編集
			fmt.Print("編集する位置: ")
			fmt.Scan(&index)

			if index >= 0 && index <= len(freeTask) {
				fmt.Print("新しいタスク: ")
				fmt.Scan(&input_task)
				freeTask[index] = input_task
				fmt.Println("タスクを更新しました")
				fmt.Printf("\n") //０～スライスの最大要素数の中から選択して選んだところを更新
			} else {
				fmt.Println("無効な位置です")
			}

		case "4": //タスク削除
			fmt.Print("削除する位置: ")
			fmt.Scan(&index)

			if index >= 0 && index <= len(freeTask) {
				if freeTask[index] == "" {
					fmt.Println("何も入ってないよ")
					fmt.Printf("\n")
				} else {
					freeTask[index] = ""
					fmt.Println("タスクを削除しました")
					fmt.Printf("\n")
				}
			} else {
				fmt.Println("無効な位置です")
				fmt.Printf("\n")
			}

		case "5": //終了
			fmt.Println("終了します")
			return

		default:
			fmt.Println("1-5の数値を入力してください")
			fmt.Printf("\n")
		}
	}
}
