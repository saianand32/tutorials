
public class C_Overflow {
    public static void main(String[] args) {
        
        // 1. int range exact(-2^31 to 2^31-1) or approx (-10^9 to 10^9)
        // 2. long range exact(  -2^63 to 2^63-1)  or approx (-10^18 to 10^18)

        System.out.println(Float.MAX_VALUE);

        // now something from type precedence again
        int a = 1000000;
        int b = 1000000;
        long c1 = a*b;

        System.out.println(c1); //-727379968 overflow as although lhs is long but int*int is calculated in int and then converted too long
        
        // so wee need to add a long in between

        long c2 = 1l*a*b;
        System.out.println(c2); //1000000000000

        c2 = a*1l*b; //1000000000000
        c2 = a*b*1l; //-727379968 no use to put 1l after int*int


        // Look into max size of arr possible in java

    }
}
