import java.util.Scanner;
import java.util.ArrayList;
import java.util.List;

public class Main {

    public static void main(String[] args) {

        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter a number: ");
        int number = scanner.nextInt();
        scanner.nextLine();

        if (number > 7) {
            System.out.println("Hello");
        }

        System.out.print("Enter a name: ");
        String name = scanner.nextLine();
        if (name.equals("John")) {
            System.out.println("Hello, John");
        } else {
            System.out.println("Theres no such name");
        }
        System.out.print("Enter numbers separated by spaces: ");
        String line = scanner.nextLine();
        List<Integer> nums = new ArrayList<>();
        if (!line.isEmpty()) {
            for (String part : line.split("\\s+")) {
                nums.add(Integer.parseInt(part));
            }
        }
        System.out.println("Multiples of 3: ");
        for (int v : nums) {
            if (v % 3 == 0) {
                System.out.print(v + " ");
            }
        }
        System.out.println();
        scanner.close();
    }
}
/*
    Answer: this string is not correct because the closing bracket ] appears while
    a round parentheses ( is still open, violating the proper nesting order.
    In go - probably as in java I would fix it via map as i did in another packet called main.go - you need to store (), [], {} in map
    and in a loop to go through the string checking each symbol whether it is has close bracket and if it does
    you need to pop this couple and move on
    But if the question is how to fix the string - replace the last two characters ]] with )]
*/
