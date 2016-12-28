import java.io.*;
import java.util.*;
import java.math.*;

public class Solution {

    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        long x = scan.nextLong();
        long y = scan.nextLong(); 
        long a = scan.nextLong(); 
        long b = scan.nextLong(); 
        scan.close();
        
        BigDecimal bX = BigDecimal.valueOf(x);
        BigDecimal bY = BigDecimal.valueOf(y);
        BigDecimal bA = BigDecimal.valueOf(a);
        BigDecimal bB = BigDecimal.valueOf(b);

        BigDecimal bK = bX.multiply(bA).add(bY.multiply(bB)).divide(bA.pow(2).add(bB.pow(2)), 10, RoundingMode.HALF_UP);
        BigDecimal bKB = bK.multiply(bB);
        BigDecimal bN = sqrt(bX.subtract(bK.multiply(bA)).pow(2).add(bY.subtract(bKB).pow(2))).divide(sqrt(bA.pow(2).add(bB.pow(2))), 10, RoundingMode.HALF_UP);

        if (bKB.compareTo(bY) > 0) {
            bN = bN.multiply(BigDecimal.valueOf(-1));
        }
        
        System.out.println(bK.toPlainString());
        System.out.println(bN.toPlainString());        
    }
    
    public static BigDecimal sqrt(BigDecimal number)
    {
        BigDecimal result = BigDecimal.ZERO;
        BigDecimal flipA = result;
        BigDecimal flipB = result;
        BigDecimal guess = BigDecimal.ONE;
        boolean first = true;
        while( result.compareTo(guess) != 0 )
        {
            if(!first)
                guess = result;
            else
                first=false;

            result = number.divide(guess, 10, RoundingMode.HALF_UP).add(guess).divide(new BigDecimal(2), 10, RoundingMode.HALF_UP);
            // handle flip flops
            if(result.equals(flipB))
                return flipA;

            flipB = flipA;
            flipA = result;
        }
        return result;
    }
}