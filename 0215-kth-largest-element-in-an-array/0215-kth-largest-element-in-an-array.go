type MaxHeap struct {
    heap []int 
}

func Parent(ind int)int{
    return (ind-1 )/2
}

func (R *MaxHeap)Swap(ind1,ind2 int){
    R.heap[ind1],R.heap[ind2]=R.heap[ind2],R.heap[ind1]
}

func(R *MaxHeap) Insert(val int){
    R.heap=append(R.heap,val)
    lastindex:=len(R.heap)-1
    R.Heapifyup(lastindex)
}

func (R *MaxHeap)Extract()int{
    if len(R.heap)==0{
        fmt.Println("Empty heap")
        return -1 
    }
    res:=R.heap[0]
    lastindex:=len(R.heap)-1
    R.heap[0]=R.heap[lastindex]
    R.heap=R.heap[:lastindex]

    R.HeapifyDown(0)

    return res
}

func (R *MaxHeap) Heapifyup(ind int) {
    for ind > 0 && R.heap[ind] > R.heap[Parent(ind)] {
        R.Swap(ind, Parent(ind))
        ind = Parent(ind)
    }
}

func(R *MaxHeap)HeapifyDown(ind int){
    last:=len(R.heap)-1
    l,r:=Left(ind),Right(ind)
    childtoswap:=0
    for l<=last{
        if l==last{
            childtoswap=l
        }else if R.heap[l]>R.heap[r]{
            childtoswap=l
        }else{
            childtoswap=r
        }

        if R.heap[ind]<R.heap[childtoswap]{
            R.Swap(ind,childtoswap)
            ind=childtoswap
            l,r=Left(ind),Right(ind)
        }else{
            return 
        }
    }
}

func Left(ind int )int {
    return 2*ind +1 
}
func Right(ind int )int {
    return 2*ind +2 
}

func findKthLargest(nums []int, k int) int {
    NMaxHeap:=MaxHeap{}
    for _,val:=range nums{
        NMaxHeap.Insert(val)
    }

    i:=0
    res:=0
    for i<k{
        res=NMaxHeap.Extract()
        // fmt.Println(NMaxHeap.heap)
        i++
    } 

    return res
    
    
}
