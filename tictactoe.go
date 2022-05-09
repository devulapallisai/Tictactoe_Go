package main
import "fmt";

func main() {
	
	// Initiating a 2D array for XO board

	xoboard := [3][3]string{}

	player:="x"

	for{

		fmt.Println("player",player)
	
		// Taking input from user for column value 

		var row int
	
		fmt.Println("Please enter a number for row wise 0 or 1 or 2")
	
		fmt.Scanf("%d\n",&row)
	
		if(row>2 || row<0){
			fmt.Println("Invalid entry!!!. Please enter numbers between 0 and 2")
			continue
		}

		// accepting column value where user's  choice 'x' or 'o' to be placed

		var col int
	
		fmt.Println("Please enter a number for column wise 0 or 1 or 2")
	
		fmt.Scanf("%d\n",&col)
	
		if(col>2 || col<0){
			fmt.Println("Invalid entry!!!. Please enter numbers between 0 and 2")
			continue
		}

		if(xoboard[row][col]==""){
			xoboard[row][col]=player
		}else{
			fmt.Println("Invalid entry:", row,col,xoboard[row][col])
			continue
		}

		// Print the XOBoard after placing 

		fmt.Println(xoboard[0])
		fmt.Println(xoboard[1])
		fmt.Println(xoboard[2])

		// did some one win
		for i := 0; i < 3 ; i++{
			// check row || columns
			if(xoboard[i][0]==xoboard[i][1]&&xoboard[i][1]==xoboard[i][2]&&xoboard[i][2]==player || xoboard[0][i]==xoboard[1][i]&&xoboard[1][i]==xoboard[2][i]&&xoboard[2][i]==player ){
				 fmt.Println("game ended , winner is player:", player)
				 return
			 }
		}
 
 
		if(xoboard[0][0] == xoboard[1][1]&&xoboard[1][1] == xoboard[2][2] && xoboard[2][2] == player || xoboard[0][2] == xoboard[1][1] && xoboard[1][1] == xoboard[2][0] && xoboard[2][0] == player) {
		 fmt.Println("game ended , winner is player:", player)
		 // 
		 return
		 }

		if(xoboard[0][0]!="" && xoboard[1][1] !=""&& xoboard[2][2]  !="" && xoboard[0][2] !="" && xoboard[2][0]  !=""&& xoboard[2][1] !=""&& xoboard[0][1] !=""&&xoboard[1][2] !=""&&xoboard[1][0] !=""){
			fmt.Println("Game draw , none is winner :(")
			return
		}

		if(player=="x"){
			player="o"
		}else{
			player="x"
		}
	}

}