const fs = require("fs");

const numbers = fs
  .readFileSync("input.txt")
  .toString()
  .split("\n")
  .filter((a) => a)
  .map((a) => a.split("").map((a) => parseInt(a)));

const sumBits = numbers.reduce((acc, curr) => {
  return acc.map((x, i) => x + curr[i]);
}, Array(numbers[0].length).fill(0));

const commonBits = sumBits.map((a) => (a / numbers.length >= 0.5 ? 1 : 0));
const unCommonBits = commonBits.map((a) => (a ? 0 : 1));

const parseAnInt = (bitArray) => parseInt(bitArray.join(""), 2);
const gammaRate = parseAnInt(commonBits);
let epsilonRate = parseAnInt(unCommonBits);

console.log(epsilonRate * gammaRate);
