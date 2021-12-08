const fs = require("fs");
const { shared, ck } = require("./helpers");

const easyToParseNumbersFound = fs
  .readFileSync("input.txt")
  .toString()
  .split("\n")
  .flatMap((a) => {
    return a.split("|")[1].trim().split(" ");
  })
  .filter((a) => [2, 3, 4, 7].includes(a.length));

console.log("Part 1", easyToParseNumbersFound.length);

const parseHardNumber = (easyNumbers, number) => {
  if (Object.values(easyNumbers).find((a) => a === number)) return {};
  if (number.length === 6) {
    if (shared(easyNumbers.four, number).length === 4) {
      return { [ck(number)]: 9 };
    }
    if (shared(easyNumbers.one, number).length === 2) {
      return { [ck(number)]: 0 };
    }
    return { [ck(number)]: 6 };
  }
  if (shared(easyNumbers.one, number).length === 2) return { [ck(number)]: 3 };
  if (shared(easyNumbers.four, number).length === 3) return { [ck(number)]: 5 };
  return { [ck(number)]: 2 };
};
const parseEasyNumbers = (line) => {
  const res = {};
  line
    .split("|")[0]
    .trim()
    .split(" ")
    .forEach((a) => {
      switch (a.length) {
        case 2:
          res["one"] = a;
          break;
        case 4:
          res["four"] = a;
          break;
        case 3:
          res["seven"] = a;
          break;
        case 7:
          res["eight"] = a;
          break;
        default:
          break;
      }
    });
  return res;
};

function parseNumber(line) {
  const easyNumbers = parseEasyNumbers(line);
  const testDigits = line.split("|")[0].trim().split(" ");
  const res = testDigits.reduce(
    (acc, curr) => {
      return {
        ...acc,
        ...parseHardNumber(easyNumbers, curr),
      };
    },
    {
      [ck(easyNumbers.one)]: 1,
      [ck(easyNumbers.four)]: 4,
      [ck(easyNumbers.seven)]: 7,
      [ck(easyNumbers.eight)]: 8,
    }
  );

  const finalDigits = line.split("|")[1].trim().split(" ");
  const number = finalDigits
    .map((n) => res[ck(n)])
    .reduce(
      (sum, curr, i, l) => sum + curr * Math.pow(10, l.length - 1 - i),
      0
    );
  return number;
}

const lines = fs.readFileSync("input.txt").toString().split("\n");
const sum = lines.reduce((sum, curr) => sum + parseNumber(curr), 0);
console.log("part 2 ", sum);
