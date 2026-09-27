package hamming

import (
    "errors")



func Distance(a, b string) (int, error) {

err := errors.New("Non equal strings")
var distance int 

		if a == "" && b == ""{
            distance, err = 0, nil 
        }
    
		if len(a) == len(b){
        for index := range a {
            if a[index] != b[index] {
                distance++
                err = nil
            	} else if a[index] == b[index]{
                	err = nil 
                	continue
                }
        	}
        }  
    return distance, err 
}

