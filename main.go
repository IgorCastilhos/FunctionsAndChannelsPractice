package main

// reference types (pointers, slices, maps, functions, channels)
// interface type
// The variable (a *Animal) in the parentheses is the "Receiver". The function is called Says.

var keyPressChan chan rune

func main() {
	//dog := Animal{
	//	Name:         "Dog",
	//	Sound:        "Woof",
	//	NumberOfLegs: 4,
	//}
	//dog.Says()
	//cat := Animal{
	//	Name:         "Cat",
	//	Sound:        "Meow",
	//	NumberOfLegs: 4,
	//}
	//cat.Says()
	//fmt.Println()
	//cat.HowManyLegs()
	//fmt.Println()
	//dog.HowManyLegs()
	//keyPressChan = make(chan rune)
	//go listenForKeyPress()
	//fmt.Println("Press any key, or q to quit")
	//_ = keyboard.Open()
	//defer func() {
	//	keyboard.Close()
	//}()
	//for {
	//	char, _, _ := keyboard.GetSingleKey()
	//	if char == 'q' || char == 'Q' {
	//		break
	//	}
	//	keyPressChan <- char
	//
	//}

	dog := Dog{
		Name:         "Dog",
		Sound:        "Woof",
		NumberOfLegs: 4,
	}
	Riddle(&dog)

	var cat Cat
	cat.Name = "Cat"
	cat.NumberOfLegs = 4
	cat.Sound = "meow"
	cat.HasTail = true

	Riddle(&cat)
}

func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}

	return total
}
