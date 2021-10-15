package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)


var contactList []listNum
var diariesList []diaries



type diaries struct {

	diar string
}


type listNum struct {
	 name string
	family string
	 phoneNumber string
	 postCode string
}
func mainMenu() int {
	scanner := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("please select")

		fmt.Println("1 = contact list")
		fmt.Println("2 = add contact")
		fmt.Println("3 = see diaries")
		fmt.Println("4 = add new diaries")

		text, _ := scanner.ReadString('\n')
		fixed := strings.ReplaceAll(text, "\n", "")

		if fixed == "1" {
			showContactlist()

		} else if fixed == "2" {
			output := addnewContact()
			if output != nil {
				fmt.Println(output)
			}
		}else if fixed == "3" {

			seeDiaries()


		} else if fixed == "4"{

			addnewDiaries()

		} else {

			fmt.Println("wrong number")

		}
	}
}

func seeDiaries() {

	fmt.Println(diariesList)

}

func addnewDiaries() {
	scan := bufio.NewReader(os.Stdin)


	fmt.Println("please write your diaries : " )


	text, _ := scan.ReadString('\n')
	diary := strings.ReplaceAll(text, "\n", "")

	var name = diaries{

		diar: diary,
	}


	diariesList = append(diariesList,name)



	}







func addnewContact() error {
	scan := bufio.NewReader(os.Stdin)

	fmt.Print("enter your name :")

	text, _ := scan.ReadString('\n')
	name := strings.ReplaceAll(text, "\n", "")
	if len(name) > 15 || len(name) < 3 {
		a := true
		for a == true {
			fmt.Println("the name value must be under 15 character and more than 3")
			fmt.Println(" 1 : try again ")
			fmt.Println(" 2 : exit ")
			t, _ := scan.ReadString('\n')
			insertedNumber := strings.ReplaceAll(t, "\n", "")
			if insertedNumber == "1" {
				fmt.Print("enter your name again :")
				text, _ := scan.ReadString('\n')
				name = strings.ReplaceAll(text, "\n", "")
				if len(name) > 15 || len(name) < 3 {
					continue
				} else {
					a = false
				}
			} else if insertedNumber == "2" {
				a = false
				return errors.New("exit")
			}
		}
	}
	fmt.Print("enter ur family :")
	text2, _ := scan.ReadString('\n')
	family := strings.ReplaceAll(text2, "\n", "")

	if len(family) > 15 || len(family) < 3 {
		b := true
		for b == true {
			fmt.Println("your family must between 3 and 15 characters")
			fmt.Println("1 : retry")
			fmt.Println("2 : exit")
			t1, _ := scan.ReadString('\n')
			insertedNums := strings.ReplaceAll(t1, "\n", "")
			if insertedNums == "1" {
				fmt.Print("enter ur family again :")
				t2, _ := scan.ReadString('\n')
				family = strings.ReplaceAll(t2, "\n", "")

				if len(family) > 15 || len(family) < 3 {
					continue

				} else {
					b = false
				}

			} else if insertedNums == "2" {
				b = false
				return errors.New("exit")
			}
		}
	}

	fmt.Print("enter ur phoneNumber : ")
	text3, _ := scan.ReadString('\n')
	phone := strings.ReplaceAll(text3, "\n", "")
	if phone[:2] != "09" {
		c := true
		for c == true {
			fmt.Println("you must be start with 09")
			fmt.Println("1 : retry")
			fmt.Println("2 : exit")

			t2, _ := scan.ReadString('\n')
			insertedNumz := strings.ReplaceAll(t2, "\n", "")

			if insertedNumz == "1" {
				fmt.Println("please enter your phoneNumber again")

				t3, _ := scan.ReadString('\n')
				phone = strings.ReplaceAll(t3, "\n", "")

				if phone[:2] != "09"{
					continue
				}else {
					c = false
				}

			}else if insertedNumz == "2" {
				c = false
				return errors.New("exit")

			}


		}

		} else if len(phone) != 11 {

			d := true
			for d == true{
			fmt.Println("your phoneNumber must be 11")
				fmt.Println("1 : retry")
				fmt.Println("2 : exit")

				t2, _ := scan.ReadString('\n')
				insertedNumzz := strings.ReplaceAll(t2, "\n", "")

				if insertedNumzz == "1" {

					fmt.Println("please enter your phoneNumber again")


					t3, _ := scan.ReadString('\n')
					phone = strings.ReplaceAll(t3, "\n", "")

					if len(phone) != 11{
						continue
					}else {
						d = false
					}


				}else if insertedNumzz == "2" {
					d = false
					return errors.New("exit")

				}




			}


	}







	fmt.Print("enter ur postCode : ")
	text4 ,_ := scan.ReadString('\n')
	postcode := strings.ReplaceAll(text4,"\n","")

	var contact = listNum{
		name:        name,
		family:      family,
		phoneNumber: phone,
		postCode:    postcode,


	}
	contactList = append(contactList,contact)
	return nil
}

func showContactlist() {


	fmt.Println(contactList)

}

func main()  {
	//a := true
	//for a == true {
	//	fmt.Println("retry?")
	//	s := fmt.Sprintln("exit?")
	//	if s == "exit?\n" {
	//		a = false
	//	}
	//}



	addDefaultContacts()
	addDiaries()
	mainMenu()

}
func addDefaultContacts() {
	var first = listNum{

		name:        "hamid",
		family:      "jahani",
		phoneNumber: "12346546",
		postCode:    "4987987",
	}
	var second = listNum{
		name:        "alireza",
		family:      "shabani",
		phoneNumber: "68987481484",
		postCode:    "54846",
	}
	contactList = append(contactList, first)
	contactList = append(contactList, second)


}

func addDiaries()  {

	var first = diaries{

		diar: "hello, this my diary and it will be fill soon....",

		}
		diariesList  = append(diariesList,first)

}
































