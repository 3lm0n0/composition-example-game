package main

import "fmt"


type Position struct {
	x float64
	y float64
}

type Player struct {
	Name string
	*Position
}

type Enemy struct {
	Name string
	*Position
}

func (p *Position) Move(x, y float64) {
	p.x += x
	p.y += y
}

func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}

func (p *Position) MoveFaster(s int) {
	p.x = p.x * float64(s)
	p.y = p.y * float64(s)
}

func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
		Position: &Position{},
	}
}

func NewEnemy(name string) *Enemy {
	return &Enemy{
		Name: name,
		Position: &Position{},
	}
}

func main() {
	player1 := NewPlayer("player1")
	fmt.Printf("%s position is %v \n",player1.Name, player1.Position)
	
	enemy :=  NewEnemy("Orc")
	fmt.Printf("%s position is %v \n",enemy.Name, enemy.Position)
	
	player1.Move(2,4)
	fmt.Printf("%s position is %v \n",player1.Name, player1.Position)

	player1.MoveFaster(int(2))
	fmt.Printf("%s position is %v \n",player1.Name, player1.Position)

	enemy.Teleport(4,8)
	fmt.Printf("%s position is %v \n",enemy.Name, enemy.Position)

	player1.Teleport(0,0)
	fmt.Printf("%s position is %v \n",player1.Name, player1.Position)

	player1.Move(-2,-4)
	fmt.Printf("%s position is %v \n",player1.Name, player1.Position)

	player1.MoveFaster(int(2))
	fmt.Printf("%s position is %v \n",player1.Name, player1.Position)

	player1.Teleport(0,0)
	fmt.Printf("%s position is %v \n",player1.Name, player1.Position)
}