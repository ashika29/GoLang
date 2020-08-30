package main
//one goroutine is the main goroutine that comes by default
import (
	"fmt"
	"runtime"
	"sync"
)
var wgIns sync.WaitGroup

func main(){
	counter := 0 //shared variable
	wgIns.Add(10) //the other 10 goroutines are supposed to come from here
	for i:=0;i<10;i++{
		go func() (){	//goroutines are made
			for j:=0;j<10;j++{
				counter += 1	//shared variable execution
				//100 should be the counter value but it may be equal to 100 or lesser due to race condition
			}
			wgIns.Done()
		}()
	}
	fmt.Println("The number of goroutines before wait = ",runtime.NumGoroutine())
	//this value should actually be 11
	wgIns.Wait()
	fmt.Println("Counter value = ",counter) //value should be 100
	fmt.Println("The number of goroutines after wait = ",runtime.NumGoroutine())
	//this value is supposed to be 1
	//but lets see if these values stay consistently same every time we run the code
}