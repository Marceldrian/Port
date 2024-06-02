package main

import "fmt"

const NMAX int = 155

type tabString [NMAX]string

type comment struct {
	username string
	cmd      string
}

type tabComment [NMAX]comment

type account struct {
	username       string
	password       string
	status         string
	friends        tabString
	nFriends       int
	statusComments tabComment
	nComments      int
	friendRequests tabString
	nRequests      int
}

type tabAccount [NMAX]account

var acc_database tabAccount
var nAccounts int
var active_id int = -1

func main() {
	var choice int = -1
	var logged_in bool = false
	for choice != 0 {
		menu(&choice, &logged_in)
	}
}
func in_view() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("================================================================")
	fmt.Println("1. View Profile")
	fmt.Println("2. New Status")
	fmt.Println("3. Add Friends")
	fmt.Println("4. View Friends")
	fmt.Println("5. View Requests")
	fmt.Println("6. Remove Friends")
	fmt.Println("7. Edit Profile")
	// fmt.Println("7. View Sorted Friends")
	fmt.Println("8. Explore")
	fmt.Println("9. Logout")
	fmt.Println("0. Exit")
	fmt.Println("================================================================")
	fmt.Print("Choose One: ")
}

func menu(choice *int, logged_in *bool) {
	if *logged_in {
		in_view()
		fmt.Scan(choice)
		switch *choice {
		case 1:
			view_profile()
		case 2:
			add_status()
		case 3:
			add_friends()
		case 4:
			view_friends()
		case 5:
			view_request()
		case 6:
			remove_friend()
		case 7:
			edit_profile()
		case 8:
			explore()
		case 9:
			*logged_in = false
		case 0:
			fmt.Println("Exiting...")
		default:
			fmt.Println("Invalid Choice")
		}
	} else {
		out_view()
		fmt.Scan(choice)
		switch *choice {
		case 1:
			login(logged_in)
		case 2:
			register(logged_in)
		case 0:
			fmt.Println("Exiting...")
		default:
			fmt.Println("Invalid Choice")
		}
	}
}

func out_view() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("================================================================")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("0. Exit")
	fmt.Println("================================================================")
	fmt.Print("Choose One: ")
}

func login(logged_in *bool) {
	var uname, pass string
	var idx int

	fmt.Print("\033[H\033[2J")

	fmt.Print("Username: ")
	fmt.Scan(&uname)
	fmt.Print("Password: ")
	fmt.Scan(&pass)
	// for i := 0; i < nAccounts; i++ {
	// 	if uname == acc_database[i].username && pass == acc_database[i].password {
	// 		fmt.Println("Login Success")
	// 		active_id = i
	// 		*logged_in = true
	// 	}
	// }
	idx = search_database(uname)
	if idx != -1 {
		if pass == acc_database[idx].password {
			fmt.Println("Login Success")
			active_id = idx
			*logged_in = true
		} else {
			fmt.Println("Login Failed")
		}
	} else {
		fmt.Println("Login Failed")
	}
	// if *logged_in == false {
	// 	fmt.Println("Login Failed")
	// }
	// fmt.Println("Login Failed")
}

func register(logged_in *bool) {
	var uname, pass string
	var unique bool = true
	var idx int

	fmt.Print("\033[H\033[2J")

	fmt.Print("Username: ")
	fmt.Scan(&uname)
	fmt.Print("Password: ")
	fmt.Scan(&pass)
	for i := 0; i < nAccounts; i++ {
		if uname == acc_database[i].username {
			fmt.Println("User already exist")
			fmt.Println("Register Failed")
			unique = false
		}
	}
	if unique {
		acc_database[nAccounts].username = uname
		acc_database[nAccounts].password = pass
		nAccounts++
		sort_database()
		idx = search_database(uname)
		active_id = idx
		fmt.Println("Register Success")
		*logged_in = true
	}
}

func view_profile() {
	var temp string

	fmt.Print("\033[H\033[2J")

	fmt.Println("Username:", acc_database[active_id].username)
	if acc_database[active_id].status == "" {
		fmt.Println("Status: ** NO STATUS **")
	} else {

		fmt.Println("Status:", acc_database[active_id].status)
	}
	fmt.Println("================================================================")
	printComments()
	fmt.Print("Type any to go back to home: ")
	inputString(&temp)
}

func add_status() {
	var status string
	var new tabComment
	fmt.Print("\033[H\033[2J")
	fmt.Print("Status: ")
	inputString(&status)
	acc_database[active_id].status = status
	acc_database[active_id].statusComments = new
	acc_database[active_id].nComments = 0
	// fmt.Println(acc_database[active_id].status)
}

func sort_database() {
	var pass int = nAccounts - 1
	var j int
	var temp account
	print_database()
	for i := 0; i < pass; i++ {
		j = i
		temp = acc_database[i+1]
		for j >= 0 && acc_database[j].username > temp.username {
			acc_database[j+1] = acc_database[j]
			j--
		}
		acc_database[j+1] = temp
	}
	fmt.Println("Sorted: ")
	print_database()
}

func print_database() {
	for i := 0; i < nAccounts; i++ {
		fmt.Print(acc_database[i].username, " ")
		fmt.Print(acc_database[i].status, " ")
		fmt.Print(acc_database[i].password, " ")
		// fmt.Print(acc_database[i].friends, " ")
		fmt.Println()
	}
}

func search_database(uname string) int {
	var idx = -1
	var left, right, mid int
	left = 0
	right = nAccounts - 1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if acc_database[mid].username == uname {
			idx = mid
		} else if acc_database[mid].username > uname {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return idx
}

func add_friends() {
	var uname string
	var temp string
	fmt.Print("\033[H\033[2J")
	if acc_database[active_id].nFriends == NMAX {
		fmt.Println("You have reached the maximum number of friends")
	} else {
		fmt.Print("Search: ")
		fmt.Scan(&uname)
		var idx int = search_database(uname)
		if uname == acc_database[active_id].username {
			fmt.Println("Itu lu sendiri bodoo")
		} else if idx == -1 {
			fmt.Println("Username Doesn't Exist")
		} else if search_string(acc_database[active_id].friends, uname, acc_database[active_id].nFriends) != -1 {
			fmt.Println(uname, "is already your friend")
		} else {
			acc_database[idx].friendRequests[acc_database[idx].nRequests] = acc_database[active_id].username
			acc_database[idx].nRequests++
			fmt.Println(acc_database[idx].friendRequests)
			// sort_strings(&acc_database[idx].friendRequests, acc_database[idx].nRequests)
		}
		fmt.Print("Press Enter to continue...")
		inputString(&temp)

	}
}

func view_friends() {

	var id int
	fmt.Print("\033[H\033[2J")
	fmt.Println("My Friends:")
	fmt.Println("================================================================")

	for i := 0; i < acc_database[active_id].nFriends; i++ {
		id = search_database(acc_database[active_id].friends[i])
		if id == -1 {
			fmt.Println("The account is not in the database, probably deleted")
		} else {
			fmt.Print(i+1, ". ", acc_database[id].username, " Status: ")
			if acc_database[id].status == "" {
				fmt.Print("**NO STATUS**\n")
			} else {
				fmt.Print(acc_database[id].status, "\n")
			}
		}
	}
	fmt.Println("================================================================")
	fmt.Println("Type 1 to comment a story")
	fmt.Print("Type any: ")
	var choice int
	fmt.Scan(&choice)
	if choice == 1 {
		add_comment()
	}
}

func sort_strings(A *tabString, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if (*A)[j] < (*A)[minIdx] {
				minIdx = j
			}
		}
		// Swap elements
		(*A)[i], (*A)[minIdx] = (*A)[minIdx], (*A)[i]
	}
	fmt.Println("Sorted: ")
	print_arr(*A, n)
}

func print_arr(A tabString, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(A[i])
		fmt.Println()
	}
}

func search_string(A tabString, uname string, n int) int {
	for i := 0; i < n; i++ {
		if A[i] == uname {
			return i
		}
	}
	return -1
}

func view_request() {
	var n int = acc_database[active_id].nRequests
	for i := 0; i < n; i++ {
		acc_idx := search_database(acc_database[active_id].friendRequests[i])
		if acc_idx != -1 {
			fmt.Print(i+1, ". ", acc_database[acc_idx].username, " Is Requesting To Be Your Friend\n")
		}
	}
	var temp int
	fmt.Print("Press 1 to accept someone as your friend ")
	fmt.Scan(&temp)
	if temp == 1 {
		acc_friends()
	}
}

func acc_friends() {
	var acc_uname string
	var acc_idx int
	fmt.Print("Who do you want to accept as a friend: ")
	fmt.Scan(&acc_uname)
	acc_idx = search_database(acc_uname)

	if acc_idx == -1 {
		fmt.Println("Username Doesn't Exist")
	} else if search_string(acc_database[active_id].friendRequests, acc_uname, acc_database[active_id].nRequests) == -1 {
		fmt.Println(acc_uname, "is not requesting to be your friend")
	} else {
		acc_database[active_id].friends[acc_database[active_id].nFriends] = acc_uname
		acc_database[active_id].nFriends++
		acc_database[acc_idx].friends[acc_database[acc_idx].nFriends] = acc_database[active_id].username
		acc_database[acc_idx].nFriends++
		// fmt.Println(acc_database[acc_idx].friends)
		// fmt.Println(acc_database[active_id].friends)

		// acc_database[acc_idx].friendRequests[acc_database[acc_idx].nRequests] = ""
		for i := search_string(acc_database[active_id].friendRequests, acc_uname, acc_database[active_id].nRequests); i < acc_database[active_id].nRequests; i++ {
			acc_database[active_id].friendRequests[i] = acc_database[active_id].friendRequests[i+1]
		}
		acc_database[active_id].nRequests--
		// sort_strings(&acc_database[acc_idx].friendRequests, acc_database[acc_idx].nRequests)

		// sort_strings(&acc_database[active_id].friends, acc_database[active_id].nFriends)

		// add active_id's username as acc_idx's friend
		// acc_database[acc_idx].friends[acc_database[acc_idx].nFriends] = acc_database[active_id].username
		// acc_database[acc_idx].nFriends++

	}
	sort_strings(&acc_database[acc_idx].friends, acc_database[acc_idx].nFriends)
	sort_strings(&acc_database[active_id].friends, acc_database[active_id].nFriends)
	fmt.Println("Press Enter to continue...")
	var temp string
	inputString(&temp)

}

func remove_friend() {
	var uname string
	var A *tabString = &acc_database[active_id].friends
	var n *int = &acc_database[active_id].nFriends

	fmt.Print("\033[H\033[2J")
	fmt.Print("Who do you want to remove as a friend: ")
	fmt.Scan(&uname)
	var idx = search_string(*A, uname, *n)

	if idx == -1 {
		fmt.Println("Username Is Not Your Friend or Doesn't Exist")
	} else {
		*n = *n - 1
		for i := idx; i < *n; i++ {
			A[i] = A[i+1]
		}
	}
}

func inputString(s *string) {
	var c byte
	var temp string
	for {
		fmt.Scanf("%c", &c)
		temp += string(c)
		if c == 13 {
			break
		}
	}
	*s = temp[1 : len(temp)-1]

}

func add_comment() {
	var uname string
	var idx int
	var cmd string
	var friend_idx int
	var friends tabString = acc_database[active_id].friends
	var n int = acc_database[active_id].nFriends
	fmt.Print("Whose status do you want to comment: ")
	fmt.Scan(&uname)
	friend_idx = search_string(friends, uname, n)
	idx = search_database(uname)
	if idx == -1 {
		fmt.Println("Username Doesn't Exist")
	} else if idx == active_id {
		fmt.Println("Itu lu sendiri bodoo")
	} else if friend_idx == -1 {
		fmt.Println("Username Is Not Your Friend")
	} else {
		fmt.Print("Comment: ")
		inputString(&cmd)
		acc_database[idx].statusComments[acc_database[idx].nComments].cmd = cmd
		acc_database[idx].statusComments[acc_database[idx].nComments].username = acc_database[active_id].username
		acc_database[idx].nComments++

	}

}

func printComments() {
	var A tabComment = acc_database[active_id].statusComments
	var n int = acc_database[active_id].nComments

	for i := 0; i < n; i++ {
		fmt.Print(A[i].username, " comments ")
		fmt.Println(A[i].cmd)
	}
}

func edit_profile() {
	var newUsername, newPassword string

	fmt.Print("\033[H\033[2J")

	fmt.Println("Edit Profile")
	fmt.Println("Leave the field empty if you don't want to change it.")

	// fmt.Scan(&newUsername)
	for {
		fmt.Print("New Username: ")
		fmt.Scan(&newUsername)
		if newUsername == acc_database[active_id].username {
			fmt.Println("Cannot use the same username twice")
		} else if newUsername == "" {
			break
		} else if search_database(newUsername) != -1 {
			fmt.Println("Username Already Exists")
		} else {
			acc_database[active_id].username = newUsername
			break
		}
	}

	// if search_database(newUsername) != -1 {
	// 	fmt.Println("Username Already Exists")

	// } else if newUsername != "" {
	// 	acc_database[active_id].username = newUsername
	// }
	for {
		fmt.Print("New Password: ")
		fmt.Scan(&newPassword)
		if newPassword == acc_database[active_id].password {
			fmt.Println("Cannot use the same password twice")
		} else if newPassword == "" {
			break
		} else {
			acc_database[active_id].password = newPassword
			break
		}
	}
	// fmt.Scan(&newPassword)
	// if newPassword == acc_database[active_id].password {
	// 	fmt.Println("Cannot use the same password twice")
	// } else if newPassword != "" {
	// 	acc_database[active_id].password = newPassword
	// }

	// fmt.Print("New Status: ")
	// fmt.Scan(&newStatus)
	// if newStatus != " " {
	// 	acc_database[active_id].status = newStatus
	// }

	fmt.Println("Profile Updated Successfully")
}

func view_sorted_friends() {
	var criteria int
	fmt.Println("Sort Friends By:")
	fmt.Println("1. Username")
	fmt.Println("2. Status")
	fmt.Print("Choose criteria: ")
	fmt.Scan(&criteria)

	switch criteria {
	case 1:
		sort_friends_by_username()
	case 2:
		sort_friends_by_status()
	default:
		fmt.Println("Invalid Choice")
		return
	}

	var id int
	fmt.Println("My Friends (Sorted):")
	fmt.Println("================================================================")
	for i := 0; i < acc_database[active_id].nFriends; i++ {
		id = search_database(acc_database[active_id].friends[i])
		if id == -1 {
			fmt.Println("The account is not in the database, probably deleted")
		} else {
			fmt.Print(i+1, ". ", acc_database[id].username, " Status: ")
			if acc_database[id].status == "" {
				fmt.Print("**NO STATUS**\n")
			} else {
				fmt.Print(acc_database[id].status, "\n")
			}
		}
	}
	fmt.Println("================================================================")
}

func sort_friends_by_username() {
	friends := &acc_database[active_id].friends
	n := acc_database[active_id].nFriends
	sort_strings(friends, n)
}

func sort_friends_by_status() {
	friends := &acc_database[active_id].friends
	n := acc_database[active_id].nFriends
	var pass int = n - 1
	var j int
	var temp string
	print_arr(*friends, n)
	for i := 0; i < pass; i++ {
		j = i
		temp = (*friends)[i+1]
		for j >= 0 && getStatus((*friends)[j]) > getStatus(temp) {
			(*friends)[j+1] = (*friends)[j]
			j--
		}
		(*friends)[j+1] = temp
	}
	fmt.Println("Sorted by status: ")
	print_arr(*friends, n)
}

func getStatus(username string) string {
	idx := search_database(username)
	if idx != -1 {
		return acc_database[idx].status
	}
	return ""
}

func explore() {
	var uname string
	var temp string

	fmt.Print("\033[H\033[2J")

	fmt.Print("Explore: ")
	fmt.Scan(&uname)
	idx := search_database(uname)
	if idx == -1 {
		fmt.Println("Username Doesn't Exist")
	} else if idx == active_id {
		fmt.Println("Itu lu sendiri bodoo")
	} else {
		fmt.Println(acc_database[idx].username, " Status: ")
		if acc_database[idx].status == "" {
			fmt.Print("**NO STATUS**\n")
		} else {
			fmt.Print(acc_database[idx].status, "\n")
		}
	}
	fmt.Print("Press enter to continue...")
	inputString(&temp)
}
