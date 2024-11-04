import java.util.*;

public class B_BinaryStrings {
    public static void main(String[] args) {

        String str = "100100";


        // 1. 1st method using parseInt 
        int x = Integer.parseInt(str, 2);
        System.out.println(x);


        //2. logical method
        // 101 = 2^2*1 + 2^1*0 + 2^0*1 
        int num = 0;
        for(int i = 0; i<str.length(); i++){
            if(str.charAt(i) == '1'){
                num+= Math.pow(2, str.length()-1-i);
            }
        }

        System.out.println(num);


        // Do this question its nice 
        // https://www.hackerearth.com/problem/algorithm/city-tour/
    }
}
