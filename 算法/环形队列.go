package main

import "fmt"

const n =  5
//head  end  free 3个属性没有用到
type Circle struct {
	head bool
	end bool
	free bool
	num int
	next *Circle
}
//创建环形链表
func CircleCreate() *Circle {
	c1:=&Circle{
		free:true,
		num:0,
		head:true,
		end:false,
	}
	temp:=make([]Circle,n)
	for i:=1;i<n;i++ {
		temp[i].num=0
		temp[i].free=true
		temp[i].head=false
		temp[i].end=false
		if i==1{
			c1.next=&temp[i]
			temp[i].next=&temp[i+1]
		}else if i==n-1{
			temp[i].head=false
			temp[i].end=true
			temp[n-1].next=c1
		}else{
			temp[i].next=&temp[i+1]
		}
	}

	return c1
}

func (c *Circle)IsEmpty()bool{
	temp:=c
	for i:=1;i<=n-1;i++{
		if temp.num==0{
			temp=temp.next
		}else{
			return false
		}
	}
	return true
}

func(c *Circle)IsFull()bool{
	count:=0
	temp:=c
	for i:=1;i<=n;i++{
		if temp.num!=0{
			count++
		}
		temp=temp.next
	}
	return count==n
}
//添加元素
func (c *Circle)Add(num int)string{
	temp:=c
	if c.IsFull(){
		return "list is full"
	}
	for i:=0;i<n;i++{
		if temp.num==0{
			temp.num=num
			return "add success"
		}else{
			temp=temp.next
		}
	}
	return ""
}

func (c *Circle)Del()string{
	temp:=c
	if c.IsEmpty(){
		return "list if null"
	}
	for i:=1;i<=n;i++{
		if temp.num!=0{
			temp.num=0
			return "del success"
		}
		temp=temp.next
	}
	return ""
}

func (c *Circle)List(){
	temp:=c
	t:=make([]int,n)
	if c.IsEmpty(){
		return
	}
	for i:=0;i<n;i++{
		if temp.num!=0{
			t[i]=temp.num
		}
		temp=temp.next
	}
	fmt.Println(t)
}

func (c *Circle)GetByNum(num int)int{
	temp:=c
	num=num%n
	for i:=0;i<=num;i++{
		if num==0{
			return temp.num
		}
		temp=temp.next
		if i==num-1{
			return temp.num
		}
	}
	return 0
}

func (c *Circle)GetHead()int{
	temp:=c
	if c.IsEmpty(){
		return 0
	}
	for i:=0;i<n;i++{
		if temp.num!=0{
			return temp.num
		}
		temp=temp.next
	}
	return 0
}

func (c *Circle)GetEnd()int{
	if c.IsEmpty(){
		return 0
	}
	if c.IsFull(){
		temp:=c
		for i:=0;i<n;i++{
			temp=temp.next
			if i==n-2{
				return temp.num
			}
		}
	}
	for i:=0;i<n;i++{
		temp:=c
		if temp.num!=0&&temp.next.num==0{
			return temp.num
		}
		temp=temp.next
	}
	return 0
}

func main(){
	c:=CircleCreate()
	fmt.Println(c.Add(1))
	fmt.Println(c.Add(2))
	fmt.Println(c.Add(3))
	fmt.Println(c.Add(4))
	fmt.Println(c.Add(5))
	fmt.Println(c.Add(6))
	c.List()
	c.Del()
	c.Del()
	fmt.Println(c.Add(2))
	fmt.Println(c.Add(2))
	c.List()
	fmt.Println(c.GetByNum(0))
	fmt.Println(c.GetByNum(1))
	fmt.Println(c.GetByNum(2))
	fmt.Println(c.GetByNum(3))
	fmt.Println(c.GetByNum(4))
	c.Del()
	c.Del()
	fmt.Println(c.Add(11))
	fmt.Println(c.Add(222))
	fmt.Println(c.GetByNum(0))
	fmt.Println(c.GetByNum(1))
	fmt.Println(c.GetByNum(2))
	fmt.Println(c.GetByNum(3))
	fmt.Println(c.GetByNum(4))
	fmt.Println(c.GetHead())
	fmt.Println(c.GetEnd())
}
