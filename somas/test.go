package example

import (
	baseagent "github.com/MattSScott/basePlatformSOMAS/BaseAgent"
)

type ICounterAgent interface {
	baseagent.IAgent[ICounterAgent]
	DoCount()
	GetCount()
}

type BaseCounterAgent struct{
	
}
