
package main

import (
    "github.com/01-edu/z01"
)

func Raid1a(x, y int) {    // задаем ширину и длину 
   if y >= 1 {             //условие: если длина больше или ровна 1
        if x > 0 {         //условаие:  если ширина больше 0
            z01.PrintRune('o')   //принтится А
        }
        for i := 1; i <= x-2; i++ {  //цикл запускается если i <= x-2
            z01.PrintRune('-')
        }
        if x > 1 {   //если ширина x > 1
            z01.PrintRune('o')
        }
        if x > 0 {   //если ширина x > 0
            z01.PrintRune(10)
        }

    }

    if y > 1 { //если длина  > 1
        if x > 0 {  //если ширина x > 0
            for j := 1; j <= y-2; j++ {    //цикл запускается если i <= длина-2
                z01.PrintRune('|')
                for i := 1; i <= x-2; i++ {   //цикл запускается если i <= ширина-2
                    z01.PrintRune(' ')
                }
                if x > 1 {
                    z01.PrintRune('|')
                }
                if x > 0 {
                    z01.PrintRune(10)
                }
            }
        }
    }

    if y > 1 {
        if x > 0 {
            z01.PrintRune('o')       }
        for i := 1; i <= x-2; i++ {   //цикл запускается если i <= ширина-2
            z01.PrintRune('-')
        }
        if x > 1 {
            z01.PrintRune('o')       }
        if x > 0 {
            z01.PrintRune(10)
        }
    }
}

func main() {
    Raid1a(5,3)
    Raid1a(5,1)
    Raid1a(1,1)
    Raid1a(1,5)
}