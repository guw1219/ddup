package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	letter, numvber := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}

	go func() {
		i := 1
		for true {
			select {
			case <-numvber:
				fmt.Println(i)
				i++
				fmt.Println(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for  {
			select {
			case <-letter:
				if i >= strings.Count(str,"")-1{
					wg.Done()
					return
				}
				fmt.Println(str[i:i+1])
				i++
				fmt.Println(str[i:i+1])
				i++
				numvber <- true
				break
			default:
				break
			}
		}
	}(&wg)
    numvber <- true
    wg.Wait()

}
