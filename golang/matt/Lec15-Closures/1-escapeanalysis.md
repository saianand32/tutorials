


// In Go, when a function returns a pointer to a local variable, there’s a risk of memory corruption if that variable is allocated on the stack.
// After the function returns, the stack frame is cleared, invalidating any pointers to those local variables.

// To mitigate this, Go's compiler performs escape analysis. It determines whether variables can escape their local scope.
// If a variable is identified as needing to persist beyond the function's lifetime (i.e., it could be accessed after the function returns),
// the compiler allocates it on the heap instead of the stack.

// Go's approach aims to allocate as much as possible on the stack for performance reasons,
//  but it intelligently moves variables to the heap when necessary to avoid issues with dangling pointers.