import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Solution {

    public static void main(String[] args) {
        /* Enter your code here. Read input from STDIN. Print output to STDOUT. Your class should be named Solution. */
        Scanner in = new Scanner(System.in);
        int q = Integer.parseInt(in.nextLine());
        String a[] = new String[q];
        String b[] = new String[q];
        for(int i=0; i < q; i++){
            a[i] = in.nextLine();
            b[i] = in.nextLine();
        }
        
        for (int i=0; i < q; i++) {
            if (MatchExists(a[i], b[i], 0, 0, a[i].length(), b[i].length())) {
                System.out.println("YES");
            }
            else {
                System.out.println("NO");
            }
        }
    }
    
    public static boolean MatchExists(String a, String b, int aIndex, int bIndex, int aLength, int bLength) {
        if ((bLength - bIndex) > (aLength - aIndex)) {
            return false;
        }
        if (bIndex == bLength) {
            for (int i=aIndex; i<aLength; i++) {
                if (Character.isUpperCase(a.charAt(i))) {
                    String temp = "";
                    for (char ch: a.toCharArray()) {
                        if (Character.isUpperCase(ch)) temp += ch;
                    }
                    return temp.equals(b);
                }
            }
            return true;
        }
        if (aIndex == aLength) {
            return false;  
        }
        if (Character.toLowerCase(a.charAt(aIndex))==Character.toLowerCase(b.charAt(bIndex))) {
            return MatchExists(a, b, aIndex+1, bIndex+1, aLength, bLength);
        }
        else if (Character.isUpperCase(a.charAt(aIndex))) {
            String temp = "";
            for (char ch: a.toCharArray()) {
                if (Character.isUpperCase(ch)) temp += ch;
            }
            return temp.equals(b);
        }
        else {
            return MatchExists(a, b, aIndex+1, bIndex, aLength, bLength);
        }
    }
}