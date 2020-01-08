package main

import "fmt"

type sparseArr struct {
	i,k,v []int
}

const (
	x=11
	y
)

func main(){
	var a [x][y]int
	for i:=0;i<x;i++{
		for j:=0;j<y;j++{
			a[i][j]=0
		}
	}
	a[1][2]=3
	a[3][4]=5
	for i:=0;i<x;i++{
		for j:=0;j<y;j++{
			fmt.Print(a[i][j],"\t")
		}
		fmt.Println()
	}
	fmt.Println("=========================")

	sp:=sparseArr{
		i: make([]int,0),
		k: make([]int,0),
		v: make([]int,0),
	}
	for i:=0;i<x;i++{
		for j:=0;j<y;j++{
			if a[i][j]!=0{
				sp.i=append(sp.i,i)
				sp.k=append(sp.k,j)
				sp.v=append(sp.v,a[i][j])
			}
		}
	}
	fmt.Println(sp)
	fmt.Println("=================================")
	fmt.Println("还原")
	var b [x][y]int
	for i:=0;i<len(sp.i);i++{
		t1:=sp.i[i]
		t2:=sp.k[i]
		t3:=sp.v[i]
		b[t1][t2]=t3
	}
	for i:=0;i<x;i++{
		for j:=0;j<y;j++{
			if b[i][j]==0{
				b[i][j]=0
			}
		}
	}
	for i:=0;i<x;i++{
		for j:=0;j<y;j++{
			fmt.Print(b[i][j],"\t")
		}
		fmt.Println()
	}
}
