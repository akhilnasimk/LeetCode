func countBattleships(board [][]byte) int {
    count:=0
    for row,_:=range board{
        for col,_:=range board[row]{
            if string(board[row][col])=="X"{
                count++
                Visit(&board ,row,col)
            }
        }
    }
    return count 
}

func Visit(grid *[][]byte , r,c int){
    if r<0 || r>=len((*grid)) || c<0 || c>=len((*grid)[r]){
        return 
    }
    if string((*grid)[r][c])!="X"{
        return 
    }
    (*grid)[r][c]='.'

    Visit(grid,r+1,c)
    Visit(grid,r-1,c)
    Visit(grid,r,c+1)
    Visit(grid,r,c-1)
}