import java.util.*;


//                Pattern 6
                // 12345
                // 1234
                // 123
                // 12
                // 1


public class P6{
    public static void main(String[] args) {

        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();

        for(int i = 0; i<n; i++){
            for(int j = 1; j<=n-i; j++){
                System.out.print(j);
            }
            System.out.println();
        }

    }
}