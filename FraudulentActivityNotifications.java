import java.io.*;
import java.util.*;

public class Solution {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int d = in.nextInt();
        int[] arr = new int[n];
        for(int i=0; i < n; i++){
            arr[i] = in.nextInt();
        }
        in.close();
        Map<Integer, Integer> map = new HashMap<Integer, Integer>();
        List<Integer> list = new ArrayList<Integer>();
        for (int i=0; i<=200; i++) {
        	map.put(i, 0);
        }

        int count=0;
        for (int i=0; i<n; i++) {

//            for (int num : list) {
//                System.out.print(num + " ");
//            }
//            System.out.println();
            
            if (i >= d) {
                int medianDouble = -1;
                if (d%2 == 0){
                	int firstIndex=d/2;
                	int secondIndex=(d/2)+1;
                	int parsedIndex=0;
                	for (int num: list) {
                		parsedIndex += map.get(num);
//                        System.out.println(num + " " + map.get(num));
                		if ((parsedIndex >= firstIndex) && (medianDouble == -1)) medianDouble = num;
                		if (parsedIndex >= secondIndex) {
                			medianDouble += num;
                			break;
                		}
                	}
                }
                else {
                	int firstIndex=(d+1)/2;
                	int parsedIndex=0;
                	for (int num: list) {
                		parsedIndex += map.get(num);
//                        System.out.println(num + " " + map.get(num));
                		if (parsedIndex >= firstIndex) {
                			medianDouble = num*2;
                			break;
                		}
                	}
                }

//                System.out.println(medianDouble);
                if (arr[i] >= medianDouble) count++;
                map.put(arr[i-d], map.get(arr[i-d])-1);
                if (map.get(arr[i-d]) == 0) {
                	list.remove((Integer)arr[i-d]);
                }
            }

            map.put(arr[i], map.get(arr[i]) + 1);
            if (!list.contains(arr[i])) {
            	if (list.size() == 0) list.add(arr[i]);
            	else list.add(getIndex(list, arr[i], 0, list.size()-1), arr[i]);
            }
        }

        System.out.println(count);
    }
    
    private static int getIndex(List<Integer> list, int element, int startIndex, int endIndex) {
        //System.out.println(startIndex + " " + endIndex + " " + list.size());
        if (startIndex >= endIndex) return (list.get(endIndex) >= element) ? endIndex : endIndex + 1;
        int median = startIndex + ((endIndex-startIndex)/2);
        if (list.get(median) < element) return getIndex(list, element, median+1, endIndex);
        else return getIndex(list, element, startIndex, median);
    }
}