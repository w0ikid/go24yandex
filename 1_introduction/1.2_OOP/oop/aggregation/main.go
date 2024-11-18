package main

type Engine struct {
	Engine string
}

type Car struct {
	Model string
	Engine *Engine
}

// func main(){
// 	engine := &Engine{
// 		Engine: "astanaMotor",
// 	}
// 	car := &Car{
// 		Model:  "bmw",
// 		Engine: engine,
// 	}
// }