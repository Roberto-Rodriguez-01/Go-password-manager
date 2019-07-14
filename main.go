package main

import (
	"fmt"
	"log"
	"os"
)

func file_creator() {
	var (
		Newfile *os.File
		err     error
	)
	Newfile, err = os.Create("storage.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(Newfile)
	Newfile.Close()
}

func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }

    return (err != nil)
}

func file_appeneder(Site string, Email string, Username string, Password string, Other string) {

    var file, err = os.OpenFile("storage.txt", os.O_RDWR, 0644)
    if isError(err) {
        return
    }
    defer file.Close()

    file.WriteString(Site)
    file.WriteString(" _-_ ")
    file.WriteString(Email)
    file.WriteString(" _-_ ")
    file.WriteString(Username)
    file.WriteString(" _-_ ")
    file.WriteString(Password)
    file.WriteString(" _-_ ")
    file.WriteString(Other)
    file.WriteString("\n")


    err = file.Sync()
    if isError(err) {
        return
    }

    fmt.Println("File Updated Successfully.")
    return
}


func add() {
	var (
		Site     string
		Email    string
		Username string
		Password string
		Other    string
		save     [5]string
	)

	fmt.Println("------------------- ADD ------------------")
	fmt.Println("# If not relavent leave blank \n")

	fmt.Print("Site Name or application: ")
	fmt.Scanln(&Site)
	print("\033[H\033[2J")
	fmt.Print("Email: ")
	fmt.Scanln(&Email)
	print("\033[H\033[2J")
	fmt.Print("Username: ")
	fmt.Scanln(&Username)
	print("\033[H\033[2J")
	fmt.Print("Password: ")
	fmt.Scanln(&Password)
	print("\033[H\033[2J")
	fmt.Print("Additional information: ")
	fmt.Scanln(&Other)
	print("\033[H\033[2J")
	fmt.Println("------------------------------------------")

	save[0] = Site
	save[1] = Email
	save[2] = Username
	save[3] = Password
	save[4] = Other

	fmt.Printf(`
  Site or application: "%v"
  Email: "%v" 
  Username: "%v" 
  Password: "%v" 
  Additional: "%v"
  `, save[0], save[1], save[2], save[3], save[4])

	file_creator()
	file_appeneder(Site, Email, Username, Password, Other)
	return

}

func lookup() {

}

func modder() {

}

func create() {

}

func main() {

	var choice int

	fmt.Print("------------ Password-Manager ------------")
	for true {
		fmt.Print(`
    What would you like to do ?
    ---------------------------
    1. Add a Password
    2. Lookup a Password
    3. Mod a Password stored
    4. Generate a Password
    5. Exit
    ---------------------------
    : `)

		fmt.Scanln(&choice)

		print("\033[H\033[2J")

		switch choice {
		case 1:
			add()
			continue
		case 2:
			lookup()
			continue
		case 3:
			modder()
			continue
		case 4:
			create()
			continue
		case 5:
			break
		default:
			fmt.Println("Not a Choice")
			continue
		}
		break
	}
}
