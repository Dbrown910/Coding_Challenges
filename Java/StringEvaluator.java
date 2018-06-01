import java.io.*;
import java.util.*;

public class StringEvaluator {

	HashMap<Character, Integer> tallyValues = null;
	private String _inputString = "";

	public StringEvaluator(String str) {
		_inputString = str;
	}

	public void evaluate() {
		tallyValues = new HashMap<Character,Integer>();

		for (char s : _inputString.toCharArray()) {
			if (Character.isUpperCase(s)) {
				// convert to lowercase for mapping
				s = Character.toLowerCase(s);
				deductPoint(s);
			} else {
				addPoint(s);
			}
		}
	}

	public void printResults() {
		Iterator entries = tallyValues.entrySet().iterator();
		while (entries.hasNext()) {
			Map.Entry pair = (Map.Entry)entries.next();
			System.out.print(pair.getKey() + ":" + pair.getValue() + " ");
			entries.remove();
		}
	}

	private void deductPoint(char key) {
		if (tallyValues.containsKey(key)) {
			Integer decValue = tallyValues.get(key) - 1;
			tallyValues.replace(key, decValue); 
		} else {
			tallyValues.put(key, -1);
		}

	}

	private void addPoint(char key) {
		if (tallyValues.containsKey(key)) {
			Integer incValue = tallyValues.get(key)  + 1;
			tallyValues.replace(key, incValue); 
		} else {
			tallyValues.put(key, 1);
		}
	}

	
}