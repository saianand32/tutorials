
public class D_PrecissionErr_double_float {
    public static void main(String[] args) {
        
        //Double and float have precision errors
        // though for namesake we have float 10^38 and double 10^308
        // but still we must keep in range of 10^18 as these give precision errors

        double c = 1e18+2;
        System.out.printf("%.0f", c); //100000000000000000000 but should be 100000000000000000002 there fore less precision

        c = 1e3+2;
        System.out.println(c); //1002 therefore with increasing size precision goes off
    }
}
