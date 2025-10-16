package main

import "fmt"

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
		limittodo()
	case 2:
		fmt.Printf("自由 ToDoリストですね!タスク上限は無制限の可変長です。\n")
		freetodo()
	} //選択完了時のメッセージ

}

func limittodo() {
	var limitTask = [5]string{}
	var choicefunction string
	var input_task string
	var index int

	fmt.Println("====固定長配列ToDoリスト====")

	for {
		fmt.Printf("1:タスク表示,2:タスク追加,3:タスク編集,4:タスク削除,5:終了\n")
		fmt.Printf("選択:")

		fmt.Scan(&choicefunction)
		switch choicefunction {

		case "1": //タスク表示
			fmt.Printf("--「ToDoリスト」--\n")
			for i := 0; i < 5; i++ {
				fmt.Print(i)
				fmt.Println(limitTask[i])
			}
			fmt.Printf("\n")

		case "2": //タスク追加
			canAdd := false //追加できるか確認

			for i := 0; i < len(limitTask); i++ {
				if limitTask[i] == "" {
					fmt.Print("タスクを入力: ")
					fmt.Scan(&input_task)
					limitTask[i] = input_task
					fmt.Printf("タスクを配列%d番目に追加しました\n ", index)
					fmt.Printf("\n")
					canAdd = true
					break
				} //空きがあって追加できたらtrueを返して終了、空きがなかったら別処理
			}
			if !canAdd {
				fmt.Println("タスク上限に達しました(上限5個)")
				fmt.Printf("\n")
			}

		case "3": //タスク編集
			fmt.Print("編集する位置(0-4): ")
			fmt.Scan(&index)

			if index >= 0 && index < 5 {
				fmt.Print("新しいタスク: ")
				fmt.Scan(&input_task)
				limitTask[index] = input_task
				fmt.Println("タスクを更新しました")
				fmt.Printf("\n") //0-4を選択して選んだところを更新
			} else {
				fmt.Println("無効な位置です")
			}

		case "4": //タスク削除
			fmt.Print("削除する位置: ")
			fmt.Scan(&index)

			if index >= 0 && index < 5 {
				if limitTask[index] == "" {
					fmt.Println("何も入ってないよ")
					fmt.Printf("\n")
				} else {
					limitTask[index] = ""
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

func freetodo() {
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
				fmt.Print(i)
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
