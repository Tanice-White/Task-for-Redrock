package main

import "fmt"

type EachHouse struct{    //每个房间有的性质
	HouseNumber int       //房间号
	CanUp int             //表示是否可以上楼 0 表示不能，1 表示能
	doorNumber int        //房间内的号码
}
func Select(n,c,origin,b int,Home map[int][]EachHouse)int{//找到第n个有楼梯的房间,b表示每层房间数,c表示层数,返回房间号,origin表示初始房间号
	var H []EachHouse                               //储存可以上楼的房间号
	var numberArr []int                             //储存可以上楼的房间号(已排序)
	var sum = 0                                     //可上楼的房间的个数
	for _,v:= range Home[c] {                       //所有的房号
		if v.CanUp==1{
			H = append(H,v)
			sum++
		}
	}
	for i:= origin;;i++{
		if i==b+1{
			i = 0
		}
		for _,v := range H{
			if v.HouseNumber==i{
				numberArr = append(numberArr,v.HouseNumber)
			}
		}
		if i==origin-1{
			break
		}
	}
	x:= n%sum                    //找到需要的房间号
	if x==0{
		return numberArr[sum-1]  //整除则返回最后一个
	}
	return numberArr[x-1]
}
func main() {
	var a,b int
	_,_ = fmt.Scan(&a,&b)  //a 表示层数，b 表示每层的房间数
	var Home =make(map[int][]EachHouse,a)
	var house EachHouse
	for z:=0;z<a;z++{      //表示初始化的层数
		for i:=0;i<b;i++{    //初始化这一层的所有房间
			_,_ = fmt.Scan(&house.CanUp,&house.doorNumber)
			house.HouseNumber = i
			Home[z] = append(Home[z],house)
		}
	}
	var origin int         //获取最开始进入的房间号
	_,_ = fmt.Scanln(&origin)
	var password = 0       //表示最后的密码
	for c:=0;c<a-1;c++{      //c表示层数
		//每一层的操作
		for _,v := range Home[c]{
			if v.HouseNumber==origin{  //找到进入的房间号
				password += v.doorNumber
				if v.CanUp==0{         //此房间不能上楼
					origin = Select(v.doorNumber,c,origin,b,Home)//找到第n个有楼梯的房间,是他上楼到的的房间
					password += origin
					break

				}else{                 //此房间能上楼时
					origin = Select(v.doorNumber-1,c,origin,b,Home)
					password += origin
					break
				}
			}
		}
	}
	fmt.Println(password)
}
