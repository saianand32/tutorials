
// to handle edge cases

function provideId(id: string | null){
    if(id == null){ //have to put this check to narrow down the type
        console.log("provide id")
        return
    }
    return id.toLowerCase()
}


