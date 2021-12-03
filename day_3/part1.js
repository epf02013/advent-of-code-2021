const fs = require("fs");
const { parseAnInt, getCommonBits, getUncommonBits } = require("./helpers");

const numbers = fs
  .readFileSync("input.txt")
  .toString()
  .split("\n")
  .filter((a) => a)
  .map((a) => a.split("").map((a) => parseInt(a)));

const commonBits = getCommonBits(numbers, 1);
const unCommonBits = getUncommonBits(numbers, 1);

const gammaRate = parseAnInt(commonBits);
let epsilonRate = parseAnInt(unCommonBits);

console.log(epsilonRate * gammaRate);
