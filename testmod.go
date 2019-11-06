package testmod

import "fmt"

func  Hi(name string) string {
	fmt.Sprint("Hi,%s\n",name)
	return fmt.Sprintf("Hi,%s!\n",name)
}
