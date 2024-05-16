# My Golang If-elseif-else Practice Repository


Explore a collection of easy to mid-level if-else programs shared in this repository. Dive into practical examples and exercises designed to enhance your understanding of conditional logic in programming. From simple decision-making tasks to more complex scenarios, each program offers insights into efficient condition handling in various contexts. Whether you're a beginner honing your skills or an experienced developer looking to refresh your knowledge, these programs provide valuable learning opportunities. Join the journey of mastering if-else statements in programming and unlock your potential to tackle real-world challenges with clarity and confidence.

## How to Use

**1. Install Golang:** Before using this repository, ensure that Golang is installed on your computer. You can download and install Golang from the official website: [golang.org](https://golang.org/).

**2. Clone the Repository:** Clone this repository to your local machine using Git. Open a terminal and run the following command:
```bash
git clone https://github.com/imrezaulkrm/Golang-if-elseif-else.git
```
**3. Navigate to the Folder:** Once the repository is cloned, navigate into the repository folder using the terminal:
```bash
cd Golang-if-elseif-else
```
**4. Run the Code:** Inside the repository folder, open the terminal and execute the Go files using the `go run` command followed by the filename. For example:
```bash
go run filename.go
```
Replace `<filename>` with the name of the Go file you want to run. This command will compile and execute the Go program.
## Index of Go-lang If-elseif-else Programs


1. [Electricity Bill Calculator](#electricity-bill-calculator)
2. [Employee Gross Salary Calculator](#employee-gross-salary-calculator)
3. [Money Note Counter](#money-note-counter)
4. [Student Result Calculator](#student-result-calculator)
5. [Profit or Loss Calculator](#profit-or-loss-calculator)
6. [Triangle Type Checker](#triangle-type-checker)
7. [Triangle Validity Checker](#triangle-validity-checker)
8. [Triangle Angle Validator](#triangle-angle-validator)
9. [Check Character Type](#check-character-type)
10. [Check Week Day](#check-week-day)
11. [Check Leap Year](#check-leap-year)

### [Electricity Bill Calculator](calculate-ebill.go)

**Description:** This Golang program calculates the electricity bill based on the total units consumed. It follows the tariff rates as follows:

-   For the first 50 units, the rate is 0.50 Taka per unit.
-   For the next 100 units, the rate is 0.75 Taka per unit.
-   For the next 100 units, the rate is 1.20 Taka per unit.
-   For units above 250, the rate is 1.50 Taka per unit.

Additionally, an additional surcharge of 20% is added to the total bill.
***Example***  
Input
```bash
Enter total units consumed: 150
```
Output
```bash
Electric Bill is 120 Taka.
```
click the title [Electricity Bill Calculator](calculate-ebill.go)

### [Employee Gross Salary Calculator](calculate-gross-salay.go)

**Description:** This Golang program calculates the gross salary of an employee based on the provided basic salary. It applies different allowances (HRA and DA) based on the basic salary range:

- For a basic salary less than or equal to 10000 Taka, the HRA is 20% and DA is 80%.
- For a basic salary between 10001 and 20000 Taka, the HRA is 25% and DA is 90%.
- For a basic salary greater than or equal to 20001 Taka, the HRA is 30% and DA is 95%.

The gross salary is calculated by adding the basic salary, HRA, and DA.

***Example:***  
Input
```bash
Input basic salary of an employee: 15000
```
Output
```bash
If Your Salary is 15000 then Your Gross Salary : 26250.00
```

click the title [Employee Gross Salary Calculator](calculate-gross-salay.go) to view the program file.

### [Money Note Counter](check-money-note.go)

**Description:** This Golang program calculates the total number of notes required to represent a given amount of money. It prompts the user to input the amount of money, and then it calculates the number of each denomination of notes required (i.e., 1000 Taka, 500 Taka, 200 Taka, 100 Taka, 50 Taka, 20 Taka, 10 Taka, 5 Taka, 2 Taka, and 1 Taka).

***Example:***  
Input
```bash
Count total number of notes in given amount
Enter amount of money: 2487
```
Output
```bash
Total number of notes: 6
1000: 2
500: 0
200: 2
100: 0
50: 1
20: 1
10: 1
5: 1
2: 1
1: 0
```

Click the title [Money Note Counter](check-money-note.go) to view the program file.

### [Student Result Calculator](calculate-result.go)

**Description:** This Golang program calculates the percentage and grade of a student based on their subject marks. It prompts the user to input marks for subjects like Bangla, English, Math, Physics, Chemistry, and Biology. After calculating the percentage, it determines the grade according to the following criteria:

- If the percentage is equal to or greater than 90%, the grade is A+.
- If the percentage is equal to or greater than 80%, the grade is A.
- If the percentage is equal to or greater than 70%, the grade is B.
- If the percentage is equal to or greater than 60%, the grade is C.
- If the percentage is equal to or greater than 50%, the grade is D.
- If the percentage is equal to or greater than 40%, the grade is E.
- If the percentage is less than 40%, the grade is F.

***Example:***  
Input
```bash
Enter student marks and find percentage and grade
Input marks of Bangla subjects: 80
Input marks of English subjects: 85
Input marks of Math subjects: 75
Input marks of Physics subjects: 90
Input marks of Chemistry subjects: 85
Input marks of Biology subjects: 80
```
Output
```bash
You Are Pass and Your Total Result Percentage is 82.50%
Your Grade: A
```

Click the title [Student Result Calculator](calculate-result.go) to view the program file.

### [Profit or Loss Calculator](calculate-p-l.go)

**Description:** This Golang program calculates the profit or loss based on the given cost price (c) and selling price (s). It prompts the user to input the cost price and selling price, then determines whether there is a profit or loss. If the selling price is greater than the cost price, it calculates the profit; otherwise, it calculates the loss.

***Example:***
Input
```bash
Calculate profit or loss
Input cost price: 150
Input selling price: 200
```
Output
```bash
Profit: 50 TAKA
```

Click the title [Profit or Loss Calculator](calculate-p-l.go) to view the program file.
### [Triangle Type Checker](check-triangle-e-s-i.go)

**Description:** This Golang program checks whether a triangle is equilateral, scalene, or isosceles based on the lengths of its sides. It prompts the user to input the lengths of the three sides of the triangle. After validating the input, it determines the type of triangle:

- If all three sides are equal, it's an equilateral triangle.
- If no two sides are equal, it's a scalene triangle.
- If two sides are equal, it's an isosceles triangle.

***Example:***
Input
```bash
Check whether triangle is equilateral, scalene or isosceles
Input first sides: 5
Input second sides: 5
Input third sides: 5
```
Output
```bash
The triangle is Equilateral
```

Click the title [Triangle Type Checker](check-triangle-e-s-i.go) to view the program file.

### [Triangle Validity Checker](check-triangle-size.go)

**Description:** This Golang program checks whether a triangle is valid or not based on the lengths of its sides. It prompts the user to input the lengths of the three sides of the triangle. After validating the input, it determines whether the triangle is valid:

- A triangle is valid if the sum of the lengths of any two sides is greater than the length of the third side for all combinations of sides.

***Example:***  
Input
```bash
Check whether triangle is valid or not if sides are given
Input first sides: 5
Input second sides: 4
Input third sides: 6
```
Output
```bash
The triangle is valid
```

Click the title [Triangle Validity Checker](check-triangle-size.go) to view the program file.

### [Triangle Angle Validator](check-triangle-angle.go)

**Description:** This Golang program checks whether a triangle is valid or not based on the sum of its angles. It prompts the user to input the measures of the three angles of the triangle. After validating the input, it determines whether the triangle is valid:

- A triangle is valid if the sum of its angles is equal to 180 degrees.

***Example:***  
Input
```bash
Check whether triangle is valid or not if angles are given
Input first angle: 60
Input second angle: 60
Input third angle: 60
```
Output
```bash
The triangle is valid
```

Click the title [Triangle Angle Validator](check-triangle-angle.go) to view the program file.

### [Check Character Type](check-character.go)

**Description:** This Golang program determines whether a given character is an alphabet, digit, or special character. It prompts the user to input a character, and then it checks the character against the ASCII values to categorize it accordingly.

***Example:***  
Input
```bash
Check whether a character is alphabet, digit or special character
Enter Character: A
```
Output
```bash
A is a Alphabet
```

Click the title [Check Character Type](check-character.go) to view the program file.

### [Check Week Day](check-week-day.go)

**Description:** This Golang program takes a week day number as input and prints the corresponding day of the week. It uses a switch statement to match the input number with the respective day of the week.

***Example:***
Input
```bash
Enter week number and print day of week
Enter week day number: 3
```
Output
```bash
Monday
```

Click the title [Check Week Day](check-week-day.go) to view the program file.

### [Check Leap Year](leepyear.go)

**Description:** This Golang program checks whether a given year is a leap year or not. It prompts the user to enter a year and then determines if it's divisible by 4. If it is, the program declares it as a leap year; otherwise, it considers it a normal year.

**Example:**  
Input
```bash
Enter Check Leap Year: 2024
```
Output
```bash
2024 is a Leap Year
```

Click the title [Check Leap Year](leepyear.go) to view the program file.

## Conclusion

Mastering conditional logic is essential for any programmer, and this repository serves as a valuable resource for honing those skills using Golang. By exploring a diverse range of if-else programs, from basic decision-making tasks to more complex scenarios, developers can enhance their understanding of condition handling in programming.

Whether you're a beginner looking to strengthen your foundational knowledge or an experienced developer seeking to refresh your skills, these programs offer practical examples and exercises to help you tackle real-world challenges with clarity and confidence. Through hands-on practice and experimentation, you can unlock your potential to write efficient and effective conditional statements in Golang.

Join the journey of mastering if-else statements and empower yourself to become a more versatile and proficient programmer.

## License

This repository and its contents are licensed under the [MIT License](https://opensource.org/licenses/MIT). You are free to use, modify, and distribute the code for both personal and commercial purposes. However, please note that while the code is provided as-is, without any warranty, attribution to the original author (if applicable) is appreciated. For more details, refer to the [LICENSE](LICENSE) file included in this repository.