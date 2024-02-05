package ciclos

import "fmt"

func Ciclo() {
	for i := 0; i < 10; i++ {
		if i == 9 {
			continue
		}
		fmt.Println(i)

	}
}
