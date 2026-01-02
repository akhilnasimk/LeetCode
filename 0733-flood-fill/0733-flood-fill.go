func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    starting:=image[sr][sc]
    if starting==color{
        return image
    }
    Visit(&image,sr,sc,color,starting)
    return image 
}

func Visit(image *[][]int,sr int ,sc int ,color int ,strtcolor int){
    if sr<0 || sr>=len(*image) || sc<0 ||sc>=len((*image)[sr]){
        return  
    }

    if (*image)[sr][sc]!=strtcolor{
        return 
    }

    (*image)[sr][sc]=color 

    Visit(image,sr+1,sc,color,strtcolor)
    Visit(image,sr-1,sc,color,strtcolor)
    Visit(image,sr,sc+1,color,strtcolor)
    Visit(image,sr,sc-1,color,strtcolor)
}




