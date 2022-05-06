package main

import "fmt"

func print_board(board [9]string) {
	fmt.Println(board[0:3])
	fmt.Println(board[3:6])
	fmt.Println(board[6:])
}

func get_win_indexs() [8][3]int { 
	top_row := [3]int{0, 1, 2}
	middle_row := [3]int{3, 4, 5}
	bottom_row :=[3]int{6, 7, 8}
	left_column := [3] int{0, 3, 6}
	middle_column := [3]int{1, 4, 7}
	right_column := [3]int{2, 5, 8}
	first_diagonals := [3]int{0, 4, 8}
	second_diagonal := [3]int{2, 4, 6}
	all_indexes := [8][3]int{
		top_row,
		middle_row,
		bottom_row,
		left_column,
		middle_column,
		right_column,
		first_diagonals,
		second_diagonal}
	return all_indexes
}

func update_board_with_user_input(board [9]string, user_input int, user_symbol string) [9]string { 
	board[user_input-1] = user_symbol //dong know why cant just return this Line?
	return board
}

func three_in_row(board [9]string, line[3]int) bool {
	if (board[line[0]] == "_"){
		return false
	}
	var same_first_second = board[line[0]] == board[line[1]]
	var same_second_third = board[line[1]] == board[line[2]]
	return same_first_second && same_second_third
}
	
func check_winner(board [9]string) bool {
	win_indexs := get_win_indexs()
	for _, line := range win_indexs{
		if three_in_row(board, line) {
			return true
		}
	}
	return false
}

func check_user_selection_valid(board [9]string, selection int) bool {
	return true
}

func main() {
	board := [9]string{"_", "_", "_", "_", "_", "_", "_", "_", "_"}
	positions := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	var game_over = false
	var user_input int
	player_tokens := [2]string{"X", "O"}
	i := 0

	for i<9 && !game_over {
		var user_selection_invalid = false
		for !user_selection_invalid {
			
			fmt.Println("please select your move using the following numbers:")
			print_board (positions)
			fmt.Println("board:")
			print_board(board)
			fmt.Scanln(&user_input)		

			current_token := player_tokens[i%2]
			user_selection_invalid = check_user_selection_valid(board, 1)
			if user_selection_invalid {
				board = update_board_with_user_input(board, user_input, current_token)
			}
		}
		game_over = check_winner(board)
		i += 1
	}

	fmt.Println("Game Over")
	print_board(board)

	if i<9{
		fmt.Println("Player ", player_tokens[(1+i)%2], " wins")
		return
	}
	if check_winner(board){
		fmt.Println("Player ", player_tokens[1+i%2], " wins")
	} else {
		fmt.Println("Draw")
	}
}