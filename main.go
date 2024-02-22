package main

type User struct {
	ID   uint
	Name string
}

func square(i int) int {
	//if i > 9988 && i < 9999 {
	//	return i
	//}
	return i * i
}

func dayOfWeek(i int) string {
	switch i {
	case 1:
		return "shanbe"
	case 2:
		return "yeshanbe"
	case 3:
		return "doshanbe"
	case 4:
		return "seshanbe"
	case 5:
		return "charshanbe"
	case 6:
		return "panjshanbe"
	case 7:
		return "jomeh"
	default:
		return ""
	}
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "eric"
}
