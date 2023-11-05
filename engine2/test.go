package main
import (
	"fmt"
	"math"
	//baseagent "github.com/MattSScott/basePlatformSOMAS/BaseAgent"
)



// type ICounterAgent interface {
// 	baseagent.IAgent[ICounterAgent]
// 	DoCount()
// 	GetCount()
// }
type xy_map struct{
	bike_x float32
	bike_y float32
}

type SingleAgent struct{
	Pedaling float64
	Braking float64
	Turning float64
	Mass float32
}
type MegaBike struct{
	Agents []SingleAgent
	acceleration float64
	velocity float64
	orientation float64
}

func (mb *MegaBike) add_Agent(agent SingleAgent){
	mb.Agents = append(mb.Agents, agent)
}
func (mb *MegaBike)force_decomposation(agent SingleAgent)(f_x float64, f_y float64){
	f_x=agent.Pedaling*float64(math.Cos(float64(math.Pi*(agent.Turning+mb.orientation))))
	f_y=agent.Pedaling*float64(math.Sin(float64(math.Pi*(agent.Turning+mb.orientation))))
	return
}
func (mb *MegaBike)vector_addition(){
	if len(mb.Agents)==0{
		return
	}
	F_X := 0.0
	F_Y := 0.0
	Mass := 0.0
	Brake := 0.0
	for _,agent := range mb.Agents{
		f_x,f_y:=mb.force_decomposation(agent)
		F_X+=f_x
		F_Y+=f_y
		Brake+=agent.Braking
		Mass+=float64(agent.Mass)
	}
	F:=math.Sqrt(F_X*F_X+F_Y*F_Y)
	mb.orientation = math.Atan2(F_Y,F_X)
	mb.acceleration = (F-Brake)/Mass
}

func (mb *MegaBike) calculate_Velocity(dt float64){
	if mb.velocity+mb.acceleration*dt<0{
		mb.velocity=0
	}else{
		mb.velocity+=mb.acceleration*dt
	}
}
func (Map *xy_map) update_loc(mb *MegaBike){
	Map.bike_x+=float32(mb.velocity*float64(math.Cos(math.Pi*mb.orientation)))
	Map.bike_x+=float32(mb.velocity*float64(math.Sin(math.Pi*mb.orientation)))
}
var gan float32
var y float64
var x float64
func main(){
	gan=0.8
	y=0.8
	x=1
	fmt.Println("go language sucks",math.Atan2(y,x)/math.Pi*180)
}