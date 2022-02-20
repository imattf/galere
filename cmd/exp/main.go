package main

import (
	"fmt"
	"html/template"
	"os"
)

// Regular code approach
type User struct {
	Name string
	Bio  string
	// Bio  template.HTML     //to keep as HTML and not have template change to string
	Age  int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	println("Experimental main.go")
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	// Regular code approach
	// user := User{
	// 	Name: "John Smith",
	// }

	// Testing/Expirmental approach
	// Naming and instantiating at same time to avoid declaring
	//  a type outside the main() as above
	// user := struct {
	// 	Name string
	// 	Age  int
	// }{
	// 	Name: "John Smith",
	// 	Age:  101,
	// }

	user := User{
		Name: "Bilbo Baggins",
		Bio:  `<script>alert("Haha, you have been h4x0r3d!");</script>`,
		Age:  112,
		Meta: UserMeta{
			Visits: 4,
		},
	}

	fmt.Println(user.Meta.Visits)

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
