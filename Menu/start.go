package main

import (
	"fmt"
)

/// Constants for class types
//TODO: Relocate these probably into a class structure
const (
	UNKNOWN = int(iota)
	WARRIOR
	ROGUE
	ARCHER
	WIZARD
	MAGE
	MAX_CLASSES
)

type sSupportedClasses struct {
	Classes     [MAX_CLASSES]sClassTypes
	Initialized bool
}

func (s *sSupportedClasses) init() {
	s.Initialized = true
}

func (s *sSupportedClasses) isInit() (retval bool) {
	return s.Initialized
}

type sClassTypes struct {
	Type int
	Name string
}

func (s *sClassTypes) getClassName() (retval bool) {
	retval = false
	if s.Type > UNKNOWN && s.Type < MAX_CLASSES {
	}
	return
}

type sCharacter struct {
	Name  string
	Age   int
	Class sClassTypes
}

func getClassType(classType int) (class sClassTypes) {
	if classType > UNKNOWN && classType < MAX_CLASSES {
		class.Type = classType
		// TODO: Could clean this up to have a single assignment to s.Name
		switch class.Type {
		case WARRIOR:
			class.Name = "Warrior"
		case ROGUE:
			class.Name = "Rogue"
		case ARCHER:
			class.Name = "Archer"
		case WIZARD:
			class.Name = "Wizard"
		case MAGE:
			class.Name = "Mage"
		default:
			fmt.Println("Failure to select appropriate class!")
			class.Name = "Unknown"
		}
	}
	return
}

func (s *sCharacter) Print() {
	fmt.Println("------------------------")
	fmt.Println("Character:")
	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Age: %d\n", s.Age)
	fmt.Printf("Type: %s\n", s.Class.Name)
}

func createCharacter() *sCharacter {
	name := ""
	age, cType := 0, 0

	fmt.Println("------------------------")
	fmt.Println("    CREATE CHARACTER")
	fmt.Println("------------------------")

	fmt.Printf("Name: ")
	_, err := fmt.Scanf("%20s", &name)

	if err != nil {
		fmt.Println("Error retrieving name!  Cancelling...")
		return nil
	}

	fmt.Printf("Age: ")
	_, err = fmt.Scanf("%d", &age)

	if err != nil {
		fmt.Println("Error retrieving age!  Cancelling...")
		return nil
	}

	fmt.Printf("Type: ")
	_, err = fmt.Scanf("%d", &cType)

	if err != nil {
		fmt.Println("Error retrieving type!  Cancelling...")
		return nil
	}

	return &sCharacter{name, age, getClassType(cType)}
}

func printMenu() {
	fmt.Printf("Please make a selection from the following menu: \n")
	fmt.Println(" 0) Exit")
	fmt.Println(" 1) Load scenario")
	fmt.Println(" 2) Create character")
	fmt.Println(" 3) Create scenario")
	fmt.Printf("\tSelection: ")
}

func checkSelection(sel int) bool {
	retVal := false
	switch sel {
	case 0:
		retVal = true
	case 1:
		fmt.Println("Attempting to load file -- please stand by!")
	case 2:
		charStore := createCharacter()
		if charStore == nil {
			fmt.Println("Failed to create character!")
		} else {
			charStore.Print()
		}
	case 3:
		fmt.Println("CREATE SCENARIO:\n!!")
	default:
		fmt.Printf("Your selection (%d) is invalid -- Please try again!\n", sel)
	}

	return retVal
}

func main() {
	selection := 0
	fmt.Println("Hello!  Welcome to Sim-U!")

	for {
		printMenu()
		_, err := fmt.Scanf("%d", &selection)
		if err != nil {
			fmt.Printf("Invalid selection made...")
			break
		} else {
			if checkSelection(selection) != false {
				break
			}
		}
	}

	fmt.Println("Good Bye!")
}
