package default_implementation

import (
	"fmt"
)
//定义游戏接口
type IGame interface {
	Start()
	End()
}
//定义游戏A
type AGame struct {

}

func (self *AGame) Start() {
	fmt.Println("AGame start")
}

func (self *AGame) End() {
	fmt.Println("AGame end")
}
//定义游戏B
type BAame struct {

}

func (self *BAame) Start() {
	fmt.Println("BGame start")
}

func (self *BAame) End() {
	fmt.Println("BGame end")
}

//定义一个指导类，提供对游戏类的默认实现
type GameDirector struct {
	game IGame
}

func (self *GameDirector) Play(){
	self.game.Start()
	self.game.End()
}

func NewGameDirector(game IGame) GameDirector{
	return GameDirector{game}
}
