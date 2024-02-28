
public class A_Uppercase_charTonum {
    public static void main(String[] args) {
        
        // lower case to upper
        char lower = 'c';
        char upper = (char)('A' + (lower-'a'));
        System.out.println(upper);


        // char to num
        char ch = '9';
        int num = ch-'0';
        System.out.println(num);
    }
}
