// This program displays the day of a week number .... TRYING TO IMPLEMENT SWITCH CASE in GO

package main
import "fmt"
func main() {	

    i := 2
    fmt.Print("Day ", i, " is ")
    switch i {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    	
     case 4:
        fmt.Println("Thrusday")
    	
     case 5:
        fmt.Println("Friday")
    	
     case 6:
        fmt.Println("Saturday")
    	
    case 7:
        fmt.Println("Sunday")
    	
    }

}
