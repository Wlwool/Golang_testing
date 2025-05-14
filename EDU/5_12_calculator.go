package main

import (
	"errors"
)


func Calculate(n1, n2 float64, operation string)(float64, error){
	switch operation {
	case "add":
		return n1 + n2, nil
	case "subtract":
        return n1 - n2, nil
    case "multiply":
        return n1 * n2, nil
    case "divide":
		if n2 == 0 {
			return 0, errors.New("division by zero")
		}
		return n1 / n2, nil

	default:
		return 0, errors.New("unknown operation")
	}

}


