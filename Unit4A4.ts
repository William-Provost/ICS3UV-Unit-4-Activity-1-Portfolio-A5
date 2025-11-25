/**
 * @author William Provost
 * @version 1.0.1
 * @date 2025-11-25
 * @fileoverview This program prompts the user for a starting and ending ASCII value (between 32 and 126)
 * and displays each number with its corresponding character.
 */

// variables
let startValue3: number = 0;
let endValue3: number = 0;
let validEnd: boolean = false;

// get starting value
startValue3 = parseInt(prompt("Please enter a number larger than 32, and less than 126: ") || "32");

// get ending value
while (!validEnd) {
  endValue3 = parseInt(prompt(`Please enter a number larger than ${startValue3} and less than 126: `) || "33");
  if (endValue3 > startValue3 && endValue3 < 127) {
    validEnd = true;
  } else {
    console.log("Invalid ending value. It must be larger than start value and less than 126.");
  }
}

// display ASCII values and characters
console.log("\nASCII values and their characters:");
for (let asciiNumber = startValue3; asciiNumber <= endValue3; asciiNumber++) {
  console.log(`${asciiNumber} = ${String.fromCharCode(asciiNumber)}`);
}

console.log("\nDone.");
