const fs = require("fs");
const { parseAnInt, getCommonBits, getUncommonBits } = require("./helpers");

const numbers = fs
  .readFileSync("input.txt")
  .toString()
  .split("\n")
  .filter((a) => a)
  .map((a) => a.split("").map((a) => parseInt(a)));

function getRating(getBitsOfInterest, arr) {
  let commonBits = getBitsOfInterest(arr);

  let candidates = arr;
  let index = 0;
  while (candidates.length > 1) {
    candidates = candidates.filter((c) => c[index] === commonBits[index]);
    index += 1;
    commonBits = getBitsOfInterest(candidates);
  }
  const rating = parseAnInt(candidates[0]);
  return rating;
}

const oxygenRating = getRating((arr) => getCommonBits(arr, 1), numbers);
const co2Rating = getRating((arr) => getUncommonBits(arr, 0), numbers);
console.log(co2Rating * oxygenRating);
