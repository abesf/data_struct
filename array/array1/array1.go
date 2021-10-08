package main

import "fmt"

func main()  {
	//数组是连续的
	arr1:=[8]int8{1,2,3,4,5,6,7,8}
	for i:=range arr1{
		fmt.Printf("一维数组addr =%v\n",&arr1[i])
	}
	//二维数组
	arr2:=[2][8]int8{
		{1,2,3,4,5,6,7,8},
		{1,2,3,4,5,6,7,8},
	}
	for i:=range arr2{
		for i2:=range arr2[i]{
			fmt.Printf("二位数组addr =%v\n",&arr2[i][i2])
		}

	}
}