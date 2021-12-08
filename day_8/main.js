const fs = require("fs");

const easyToParseNumbersFound = fs
  .readFileSync("input.txt")
  .toString()
  .split("\n")
  .flatMap((a) => {
    return a.split("|")[1].trim().split(" ");
  })
  .filter((a) => [2, 3, 4, 7].includes(a.length));

console.log("Part 1", easyToParseNumbersFound.length);
