import java.io.*;
import java.util.*;

public class Solution {

    public static void main(String[] args) {
        /* Enter your code here. Read input from STDIN. Print output to STDOUT. Your class should be named Solution. */
        Scanner scan = new Scanner(System.in);
        
        String str = scan.nextLine();
        int m = scan.nextInt();
        
        scan.close();
        
        int length = str.length();
        int minRepeatLength = length;
        for (int i=0; i<length; i++) {
            int repeatLength=i+1;
            if (length%repeatLength == 0) {
                int numRepeat = length/repeatLength;
                if (new String(new char[numRepeat]).replace("\0", str.substring(0, repeatLength)).equals(str)) {
                    minRepeatLength = repeatLength;
                    break;
                }
            }
        }
        
        System.out.println(m/minRepeatLength);
        
    }
}