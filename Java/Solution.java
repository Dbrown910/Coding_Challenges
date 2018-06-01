import java.io.*;
import java.util.*;

public class Solution {
	//instantiate variables;

	public static void main(String args[]) {
		entryTime("5111","923857614");
	}

	static int entryTime(String s, String keypad) {
        int totalTime = 0;
        int[][] keypadGrid = new int[3][3];
        System.out.println(s);
        for (int i = 0; i < keypadGrid.length; i++) {
        	String row = "";
        	if (i == 0) {
        		row = keypad.substring(0, 3);
        	} else if (i == 1) {
        		row = keypad.substring(3, 6);
        	} else if (i == 2) {
        		row = keypad.substring(6, 9);
        	}
        	System.out.println(row);
            for (int j = 0; j < keypadGrid[i].length; j++) {
                keypadGrid[i][j] = row.charAt(j);
            }
        }
        
        for (int i = 0; i < s.length() - 1; i++) {
        	if (i != 0 && (int)s.charAt(i) == (int)s.charAt(i-1)) {
        		continue;
        	}
            totalTime += findDist((int)s.charAt(i), (int)s.charAt(i+1), keypadGrid);
        }
        
        
        return totalTime;

    }

    static int findDist(int currNum, int numToFind, int[][] keypad) {
        int currX = 0;
        int currY = 0;
        int targetX = 0;
        int targetY = 0;
        int time = 0;
        
        boolean baseCoordsFound = false;
        boolean targetCoordsFound = false;
        
        for (int i = 0; i < keypad.length; i++) {
            for (int j = 0; j < keypad[i].length; j++) {
                if (!baseCoordsFound && keypad[i][j] == currNum) {
                    currX = i;
                    currY = j;
                    baseCoordsFound = true;
                } else if (!targetCoordsFound && keypad[i][j] == numToFind) {
                    targetX = i;
                    targetY = j;
                    targetCoordsFound = true;
                } else if (targetCoordsFound && baseCoordsFound) {
                    break;
                }
                
            }
            if (targetCoordsFound && baseCoordsFound) {
                break;
            }
        }
        System.out.println("base: " + currX + "," +currY);
        System.out.println("target: " + targetX + "," +targetY);
       
        
        if (Math.abs(currX - targetX) >= Math.abs(currY - targetY) ) {
            time = Math.abs(currX - targetX);
        } else {
            time = Math.abs(currY - targetY);
        }
         System.out.println(time);
        return time;
    }
}