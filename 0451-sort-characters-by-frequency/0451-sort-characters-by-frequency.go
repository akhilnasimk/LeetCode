type data struct {
    str string 
    count int 
}

type baseheap []data 

func (r baseheap)Len()int{
    return len(r)
}

func (r baseheap)Less(i,j int)bool{
    return r[i].count>r[j].count
}

func (r baseheap )Swap(i,j int){
    r[i],r[j]=r[j],r[i]
}

func (r *baseheap)Push(val interface{}){
    *r=append(*r,val.(data))
}

func (r *baseheap)Pop()(interface{}){
    old:=*r 
    n:=len(old)
    val:=old[n-1]
    *r=old[:n-1]
    return val 
}

func frequencySort(s string) string {
    map1:=make(map[string]int)
    for _,val:=range s{
        map1[string(val)]++
    }
    heap1:=&baseheap{}
    heap.Init(heap1)

    for key,val:=range map1{
        val:=data{str:key,count:val}
        heap.Push(heap1,val)
    }

    n:=len(*heap1)
    res:=""
    for i:=0;i<n;i++{
        val:=heap.Pop(heap1).(data)
        for j:=0;j<val.count;j++{
            res+=val.str
        }
    }

    return res
}