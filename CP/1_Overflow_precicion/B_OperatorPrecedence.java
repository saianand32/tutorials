import java.util.*;


public class B_OperatorPrecedence {
    public static void main(String[] args) {

        System.out.println(7/2*3); // prints 9 
        System.out.println(3*7/2); // prints 10
        //the aboove happens coz again its all calculated in int

        System.out.println(7/2.0*3); // prints 10.5
        System.out.println(3*7/2.0); // prints 10.5
        
        String sb = new StringBuilder("sai").reverse().toString();
        System.out.println(sb);

    }
}
