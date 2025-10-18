package personaldata

import "fmt"

type Personal struct {
    Name   string
    Weight float64
    Height float64
}

func (p Personal) Print() {
    //fmt.Println("Имя:", p.Name)
    //fmt.Println("Вес:", p.Weight)
    //fmt.Println("Рост:", p.Height)

    fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n", p.Name, p.Weight, p.Height)

}
