package main

import "reflect"

type Vehicle interface {
    CalculateTravelTime(distance float64) float64
}

type Car struct {
    Speed    float64
    Type     string
    FuelType string
}

type Motorcycle struct {
    Speed float64
    Type  string
}

func NewCar(speed float64, model string) Car {
    return Car{
        Speed:   speed,
        Type:    model,
        FuelType: "",
    }
}

func NewMotorcycle(speed float64, model string) Motorcycle {
    return Motorcycle{
        Speed: speed,
        Type:  model,
    }
}

// Оставляем методы со значением-приемником
func (c Car) CalculateTravelTime(distance float64) float64 {
    return distance / c.Speed
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
    return distance / m.Speed
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
    m := make(map[string]float64)
    for _, v := range vehicles {
        // Получаем полное имя типа (например, "main.Car")
        typeName := reflect.TypeOf(v).String()
        m[typeName] = v.CalculateTravelTime(distance)
    }
    return m
}