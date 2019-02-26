package inheritance

import "fmt"

//定义接口
type IGame interface {
	Start()
	End()
}
//定义基础游戏类
type BaseGame struct {

}

func (self *BaseGame) Start() {
	fmt.Println("BaseGame start")
	self.End()
}

func (self *BaseGame) End() {
	fmt.Println("BaseGame end")
}

//定义实际游戏类
type PartGame struct {
	BaseGame
}
//此处只覆写End方法，不覆写该方法的调用者
func (self *PartGame) End(){
	fmt.Println("PartGame end")
}
//定义实际游戏类，覆写被调用方法和调用方法
type FullGame struct {
	BaseGame
}
//此处覆写了调用方法，但是其实跟不覆盖基本没有区别，导致代码冗余
func (self *FullGame) Start(){
	fmt.Println("FullGame Start")
	self.End()
}

func (self *FullGame) End(){
	fmt.Println("FullGame End")
}
