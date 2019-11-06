package testmod

import (
	"errors"
	"fmt"
)

func Hi(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi,%s!", name), nil
	case "pt":
		return fmt.Sprintf("0i,%s!", name), nil
	case "es":
		return fmt.Sprintf("iHola,%s!", name), nil
	case "fr":
		return fmt.Sprintf("Bonjour,%s", name), nil
	default:
		return "", errors.New("unknown language")
	}
	//fmt.Sprint("Hi,%s\n",name)
	//return fmt.Sprintf("Hi,%s!\n",name)
}
