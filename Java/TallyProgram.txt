import java.io.*;
import java.util.*;

public class TallyProgram {

	static StringEvaluator stringEvaluator = null;

	public static void main(String args[]) {

		if (args.length == 0) {
			System.out.println("Enter an input string");
			System.out.println("Uppercase == -1");
			System.out.println("Lowercase == +1");
			return;
		}

		stringEvaluator = new StringEvaluator(args[0]);

		stringEvaluator.evaluate();
		stringEvaluator.printResults();
	}
}
