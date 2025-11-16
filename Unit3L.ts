/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-15
 * @fileoverview This program chooses a sweater colour and prints a message based on the choice.
 */

// variables
let colour: string = "";

// input colour
colour = prompt("Please choose a sweater colour from the available choices: Blue, Black, Red, White.") || "";
colour = colour.trim();

// check colour choice
if (colour.toLowerCase() === "black" || colour.toLowerCase() === "white") {
  console.log("You have enough sweaters in this colour.");
} else if (colour.toLowerCase() === "red") {
  console.log("This colour will look good on you.");
  console.log("How about a pair of jeans to go with the sweater?");
} else if (colour.toLowerCase() === "blue") {
  console.log("This colour doesn’t go well with your eyes.");
} else {
  console.log("Your colour choice is invalid.");
}

console.log("\nDone.");
