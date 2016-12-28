import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Solution {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        long n = in.nextLong();
        long m = in.nextLong();
        int k = in.nextInt();
        long r[] = new long[k];
        long c1[] = new long[k];
        long c2[] = new long[k];
        HashMap<Long, List<Length>> lengths = new HashMap<Long, List<Length>>();
        for(int c_i=0; c_i < k; c_i++){
            r[c_i] = in.nextLong();
            c1[c_i] = in.nextLong();
            c2[c_i] = in.nextLong();
            if (!lengths.containsKey(r[c_i])) {
                lengths.put(r[c_i], new ArrayList<Length>());
            }
        }
        
        for (int i=0; i<k; i++) {
            List<Length> lengthList = lengths.get(r[i]);
            int index=0;
            for (Length length: lengthList) {
                if (length.c1 > c1[i]) {
                    break;
                }
                index++;
            }
            Length l = new Solution().new Length();
            l.c1 = c1[i];
            l.c2 = c2[i];
            lengthList.add(index, l);
        }
        
        Collection<List<Length>> values = lengths.values();        
        long occupiedCells = 0;
        for (List<Length> list: values) {
            long lastEndIndex = 0;
            for (Length ints: list) {
                if (ints.c2 >  lastEndIndex) {
                    occupiedCells += ints.c2 - (Math.max(lastEndIndex+1, ints.c1)) + 1;
                    lastEndIndex = ints.c2;
                }
            }
        }
        
        long area = (m*n) - occupiedCells;
        System.out.println(area);
    }
    
    public class Length {
        long c1;
        long c2;
    }
}