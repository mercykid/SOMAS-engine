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
	Pedaling float32
	Braking float32
	Turning float32
	Mass float32
}
type MegaBike struct{
	Agents []SingleAgent
	acceleration float32
	velocity float32
	orientation float32
}
type Motion interface{
	calculate_Orientation()
	calculate_ResultantAcceleration()
	calculate_Velocity()
	update_loc()
}
func (mb *MegaBike) add_Agent(agent SingleAgent){
	mb.Agents = append(mb.Agents, agent)
}

func (mb *MegaBike)calculate_Orientation(){
	if len(mb.Agents)==0{
		return
	}
	Total_turning:=0.0
	for _,agent := range mb.Agents{
		Total_turning+=float64(agent.Turning)
	}
	Average_turning:=Total_turning/float64(len(mb.Agents))
	mb.orientation+=float32(Average_turning)
}
func (mb *MegaBike) calculate_ResultantAcceleration(){
	force_map:=4.0
	if len(mb.Agents)==0{
		return
	}
	Total_pedal:=0.0
	Total_brake:=0.0
	Total_mass:=0.0
	for _,agent := range mb.Agents{
		Total_mass+=float64(agent.Mass)
		if agent.Pedaling !=0{
			Total_pedal+=float64(agent.Pedaling)
		}else{
			Total_brake+=float64(agent.Braking)
		}
	}
	F:=force_map*(float64(Total_pedal)-float64(Total_brake))
	mb.acceleration= float32(F/float64(Total_mass))
}
func (mb *MegaBike) calculate_Velocity(dt float32){
	if mb.velocity+mb.acceleration*dt<0{
		mb.velocity=0
	}else{
		mb.velocity+=mb.acceleration*dt
	}
}
func (Map *xy_map) update_loc(mb *MegaBike){
	Map.bike_x+=mb.velocity*float32(math.Cos(float64(math.Pi*mb.orientation)))
	Map.bike_x+=mb.velocity*float32(math.Sin(float64(math.Pi*mb.orientation)))
}


var gan float32
func main(){
	gan=0.8
	
	fmt.Println("go language sucks",math.Acos(0))
}