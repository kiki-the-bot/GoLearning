package raindrops

import ("fmt"
       "strings")

func Convert(number int) string {

builder := strings.Builder{}
    
    if number%7 != 0 && number%5 != 0 && number%3 != 0 {
		return fmt.Sprintf("%d", number)
        }
    
        if number%3 == 0 {
            builder.WriteString ("Pling")
        } 
    	if number%5 == 0 {
    		builder.WriteString ("Plang")        
        }
    	if number%7 == 0 {
            builder.WriteString ("Plong")
        }
    
	return builder.String()

}
