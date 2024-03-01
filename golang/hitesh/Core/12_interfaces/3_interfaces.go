package main

import "fmt"

//****** TYPE ASSERTION *****
//from interface find out which struct it belongs to

type expense interface {
	cost() float64
}

type email struct {
	emailId      string
	costPerEmail int
}

func (e email) cost() float64 {
	return float64(e.costPerEmail)
}

type sms struct {
	phNum      int64
	costPerSms int
}

func (s sms) cost() float64 {
	return float64(s.costPerSms)
}

//now say we have a method for the interface expense

func getDetailsAsPerType(exp expense) {
	//say we wat to do one thing if its email and another if sms

	myEmail, ok := exp.(email)

	if ok {
		fmt.Println("its email")
		fmt.Println(myEmail.emailId)
	}

	mySms, ok := exp.(sms)

	if ok {
		fmt.Println("its sms")
		fmt.Println(mySms.phNum)
	}

	//Alternatively u can also do but remember .(type) can only be used inside switch
	switch e := exp.(type) {
	case email:
		fmt.Println("It's an email")
		fmt.Println("Email ID:", e.emailId)
		fmt.Println("Cost per email:", e.costPerEmail)
	case sms:
		fmt.Println("It's an SMS")
		fmt.Println("Phone number:", e.phNum)
		fmt.Println("Cost per SMS:", e.costPerSms)
	default:
		fmt.Println("Unknown type of expense")
	}
}

func main() {
	// Create instances of email and sms
	myEmail := email{emailId: "example@example.com", costPerEmail: 10}
	mySMS := sms{phNum: 1234567890, costPerSms: 5}

	// Call the getCostAsPerType function for each expense
	getDetailsAsPerType(myEmail)
	getDetailsAsPerType(mySMS)
}
