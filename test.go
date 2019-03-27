package main

import "fmt"

func fact(n int)int {
	if n <= 1{
		return 1
	}
	return n * fact(n-1)
}

func redc(n int)int{
	if n <=1{
		return 1
	}
	return redc(n-1)
}

func main() {

	fmt.Println(redc(5))



	// var i= 1;
 //   for   {
   	
 //   	fmt.Printf("This loop will run forever.\n");
 //   	i +=1;
 //    if i>6 {
 //    	fmt.Printf("i is greater than 6.\n");
 //        break
 //    }
 //   }
}