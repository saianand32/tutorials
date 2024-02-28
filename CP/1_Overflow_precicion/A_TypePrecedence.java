import java.util.*;

// Topics - Type precedence, Operator precedence and overflow 

public class A_TypePrecedence {
    public static void main(String[] args) {
        
        // 1. While calculating datatypes precedence is 
        // Double <- Float <- long <- int <- char

        System.out.println(7/2); // though should be 3.5 it will print 7 coz both are int so int is printed
        System.out.println(7/2.0); // prints 3.5 as even one is float which has higher precedence so loat val is printed
        System.out.println('c'+1); // prints 100 as int has higher precedence that char so answer is in int

        float x  = 7/2; // this also prints 3.0 coz still both numerator & denom are int so calculated in int and then stored in float
    }
}