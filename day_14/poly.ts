const fs = require("fs");

const getNextString = (start: string[], rules: { [k: string]: string }) => {
  const res = [];
  const items = start;
  res.push(items[0]);
  for (let i = 0; i < items.length; i++) {
    const key = items[i] + items[i + 1];
    if (rules[key]) {
      res.push(rules[key]);
    }
    if (items[i + 1]) {
      res.push(items[i + 1]);
    }
  }
  return res;
};
const getOccurrences = (list: string[]) => {
  const counts: any = {};
  list.forEach((c) => {
    if (counts[c]) {
      counts[c] = counts[c] + 1;
    } else {
      counts[c] = 1;
    }
  });
  return counts;
};
const part1 = (fileName: string, iters) => {
  const [startPoly, rulesStrings] = fs
    .readFileSync(fileName)
    .toString()
    .split("\n\n");

  const ruleMap: { [k: string]: string } = rulesStrings
    .split("\n")
    .map((r: any) => r.split(" -> "))
    .reduce((acc: any, [pair, sub]: any) => ({ ...acc, [pair]: sub }), {});
  let next = startPoly.split("");
  for (let i = 0; i < iters; i++) {
    next = getNextString(next, ruleMap);
  }
  let o = getOccurrences(next);
  const max = next.reduce((acc, curr) => (acc > o[curr] ? acc : o[curr]));
  const min = next.reduce((acc, curr) => (acc < o[curr] ? acc : o[curr]));
  console.log("Part 1 solution", max - min);
};

function getAllPairs(
  allLetters: string[],
  startPoly: any
): { [p: string]: number } {
  const pairs = {};
  allLetters.forEach((a) => {
    allLetters.forEach((b) => {
      pairs[a + b] = 0;
    });
  });
  for (let i = 0; i < startPoly.length - 1; i++) {
    pairs[startPoly.substr(i, 2)] += 1;
  }
  return pairs;
}

function getStartCountsOfSingles(allLetters: string[], startPoly) {
  const singles = allLetters.reduce((acc, l) => ({ ...acc, [l]: 0 }), {});
  startPoly.split("").forEach((c) => {
    singles[c] += 1;
  });
  return singles;
}

const part2 = (fileName: string, iters) => {
  let inputString = fs.readFileSync(fileName).toString();
  const [startPoly, rulesStrings] = inputString.split("\n\n");

  const allLetters = Object.keys(
    inputString
      .split("")
      .filter((a) => a !== a.toLowerCase())
      .reduce((a, c) => ({ ...a, [c]: c }), {})
  );
  const ruleMap: { [k: string]: string[] } = rulesStrings
    .split("\n")
    .map((r: any) => r.split(" -> "))
    .reduce(
      (acc: any, [pair, sub]: any) => ({
        ...acc,
        [pair]: [`${pair.charAt(0)}${sub}`, `${sub}${pair.charAt(1)}`],
      }),
      {}
    );

  const pairs = getAllPairs(allLetters, startPoly);
  const singles = getStartCountsOfSingles(allLetters, startPoly);
  for (let i = 0; i < iters; i++) {
    const dup = { ...pairs };
    Object.entries(dup).forEach(([p, v]) => {
      if (v !== 0) {
        const newPairs = ruleMap[p];
        const newChar = newPairs[1].charAt(0);
        singles[newChar] += v;
        newPairs.forEach((np) => {
          pairs[np] = pairs[np] + v;
        });
        pairs[p] = Math.max(pairs[p] - v, 0);
      }
    });
  }
  const max = Object.values(singles).reduce(
    (acc, curr) => (acc > curr ? acc : curr),
    0
  );
  const min = Object.values(singles)
    .filter((a) => a > 0)
    .reduce((acc, curr) => (acc < curr ? acc : curr), 10000000000000);
  console.log("Part 2 solution", max - min);
};

part1("input.txt", 10);

part2("input.txt", 40);
