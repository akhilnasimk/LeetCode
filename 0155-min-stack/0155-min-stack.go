type MinStack struct {
    Main []int
    min []int
}


func Constructor() MinStack {
    return MinStack{
        Main :[]int{},
        min :[]int{},
    }
}



func (this *MinStack) Push(val int)  {
    this.Main=append(this.Main,val)
    if len(this.min)!=0{
        if this.min[len(this.min)-1]>=val{
            this.min=append(this.min,val)
        }
    }else{
        this.min=append(this.min,val)
    }
}


func (this *MinStack) Pop()  {
    if len(this.Main)==0 || len(this.min)==0{
        return 
    }
    itemToRemove:=this.Main[len(this.Main)-1]
    this.Main=this.Main[:len(this.Main)-1]
    if this.min[len(this.min)-1]==itemToRemove{
        this.min=this.min[:len(this.min)-1]
    }
}


func (this *MinStack) Top() int {
    return this.Main[len(this.Main)-1]
}


func (this *MinStack) GetMin() int {    
    if len(this.min)==0{
        return -1 
    }
    return this.min[len(this.min)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */