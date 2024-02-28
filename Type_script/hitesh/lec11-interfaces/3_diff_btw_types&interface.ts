// 1. Difference between types and interfaces

// One major difference between type aliases vs interfaces
// are that interfaces are open and type aliases are closed.
// This means you can extend an interface by declaring it
// a second time.

//this is called "REOPENING OF INTERFACE"
interface Kitten {
    purrs: boolean;
}

interface Kitten {
    colour: string;
}

let kitt1: Kitten = { purrs: false, colour: "black" }

// In the other case a type cannot be changed outside of
// its declaration.

type Puppy = { //err
    color: string;
};

type Puppy = { //err
    toys: number;
};

// Depending on your goals, this difference could be a
// positive or a negative. However for publicly exposed
// types, it's a better call to make them an interface.


export { }