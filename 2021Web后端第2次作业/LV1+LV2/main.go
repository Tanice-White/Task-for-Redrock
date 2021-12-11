package main

import "fmt"

type BiLiBiLi struct{
	Name string //名字
	VIP bool    //是否为会员
	Icon string //头像
	Signature string //签名
	Focus int   //关注人数
	Video map[string]int //视频播放量 & 点赞 收藏 投币 & 一键三连
}
func (B *BiLiBiLi)MyName(name string){
	B.Name = name
}
func (B *BiLiBiLi)MyVIP(ToF bool){
	B.VIP = ToF
}
func (B *BiLiBiLi)MyIcon(icon string){
	B.Icon = icon
}
func (B *BiLiBiLi)MySignature(signature string){
	B.Signature = signature
}
func (B *BiLiBiLi)MyFocus(focus int){
	B.Focus = focus
}
func (B *BiLiBiLi)MyVideo(video map[string]int){
	B.Video = video
}
func MyPrint (B BiLiBiLi){
	fmt.Printf("%+v",B)
}
func main() {
	var B = &BiLiBiLi{}
	var video = map[string]int{
		"视频":1020e4,
		"点赞":2000e4,
		"投币":320e4,
		"收藏":520e4,
		"一键三连":400e4,
	}
	B.MyName("逗逼的雀巢")
	B.MyVIP(true)
	B.MyIcon("雀巢")
	B.MySignature("搞笑区知名up猪")
	B.MyFocus(408e4)
	B.MyVideo(video)
	MyPrint(*B)
}