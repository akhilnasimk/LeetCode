type MinHeap struct{
    heap []Item 
}

type Item struct {
    number int 
    freq int 
}
func (R *MinHeap )Insert(itemstruct Item ,max int ){
    R.heap=append(R.heap,itemstruct)
    lastindex:=len(R.heap)-1
    R.Heapifyup(lastindex ,itemstruct.freq)
    if len(R.heap)>max{
        R.Exctract()
    }
}

func (R *MinHeap)Exctract()int{
    if len(R.heap)==0{
        return -1 
    }
    res:=R.heap[0]
    lastind:=len(R.heap)-1
    R.heap[0]=R.heap[lastind]
    R.heap=R.heap[:lastind]
    R.HeapifyDown(0)
    return res.number 
}


func (R *MinHeap) HeapifyDown(ind int) {
    lastindex := len(R.heap) - 1

    for {
        l := LeftChild(ind)
        r := RightChild(ind)

        if l > lastindex {
            return // no children
        }

        childtoswap := l

        if r <= lastindex && R.heap[r].freq < R.heap[l].freq {
            childtoswap = r
        }

        if R.heap[ind].freq > R.heap[childtoswap].freq {
            R.Swap(ind, childtoswap)
            ind = childtoswap
        } else {
            return
        }
    }
}



func (R *MinHeap)Heapifyup(ind int ,freq int ){
    for R.heap[Parent(ind)].freq > freq &&  ind>0 {
        R.Swap(Parent(ind),ind)
        ind=Parent(ind)
    }
}

func(R *MinHeap) Swap(ind1 ,ind2 int){
    R.heap[ind1],R.heap[ind2]=R.heap[ind2],R.heap[ind1]
}

func Parent(c_index int )int{
    return (c_index-1)/2 
}

func LeftChild(p_ind int )int {
    return p_ind*2 +1 
}
func RightChild(p_ind int )int {
    return p_ind*2 +2
}

func topKFrequent(nums []int, k int) []int {
    res:=[]int{}
    map1:=make(map[int]int)
    for _,val:=range nums{
        map1[val]++
    }
    NewMinHeap:=MinHeap{}
    for key,val:=range map1{
        NewMinHeap.Insert(Item{number:key,freq:val},k)
    }
    
    for _,val:=range NewMinHeap.heap{
        res=append(res,val.number)
    }

    return res
}
