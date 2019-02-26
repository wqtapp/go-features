package inheritance

import "testing"

//因为PartGame只重写了Start方法,当调用game.Start()时相当于调用game.BaseGame.Start()
//此时在Start()中调用self.End()相当于调用game.BaseGame.End()
func TestPartGame_End(t *testing.T) {
	game := PartGame{}
	game.Start()
	/*
		output:
		BaseGame start
		BaseGame end
	 */
}

//因为FullGame重写了Start方法和End方法,当调用game.Start()时相当于调用game.Start()
//此时在Start()中调用self.End()相当于调用game.End()
//这种特性导致go实现继承的同时，想提供final方法或者公共的无需重写的方法成为一种奢侈行为
//Golang并没有传统意义上的多态。
func TestFullGame_End(t *testing.T) {
	game := FullGame{}

	//var baseGame BaseGame
	//baseGame = game 无法将FullGame当做BaseGame使用

	game.Start()

	/*
		output:
		FullGame Start
		FullGame End
	 */
}


