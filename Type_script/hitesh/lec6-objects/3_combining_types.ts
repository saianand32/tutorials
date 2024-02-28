
// can combine previous types to make a new type


type cardNumber = {
    card_number : string
}

type cardDate = {
    card_date: string
}


// 1. The beloe is allowed but bad practice 

type cardDetails = cardNumber & cardDate & {
    cvv: number
}

// 2. always create a type and then combine (use this always)

type cvvNumber = {
    cvv: number
}

type cardDetails2 = cardNumber & cardDate & cvvNumber





export {}