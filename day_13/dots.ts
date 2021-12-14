import * as fs from "fs";

type Dot = {
  x: number;
  y: number;
};
type Fold = {
  x: number;
  y: number;
};
const key = (dot: Dot) => `${dot.x},${dot.y}`;

function getNewLocation(
  foldLocation: number,
  pointLocation: number,
  max: number
) {
  if (foldLocation != 0) {
    if (pointLocation < foldLocation) {
      if (foldLocation < max / 2) {
        return pointLocation + max - foldLocation;
      } else {
        return pointLocation;
      }
    } else {
      return 2 * foldLocation - pointLocation;
    }
  } else {
    return pointLocation;
  }
}

const fold = (
  dots: Dot[],
  fold: Fold,
  height: number,
  width: number
): { width: number; transformedDots: any; height: number } => {
  const transformedDots: { [k: string]: Dot } = {};
  dots.forEach((d) => {
    const newDot = { x: 0, y: 0 };
    newDot.x = getNewLocation(fold.x, d.x, width);
    newDot.y = getNewLocation(fold.y, d.y, height);
    transformedDots[key(newDot)] = newDot;
  });
  return {
    transformedDots: Object.values(transformedDots),
    height:
      height >= 2 * fold.y
        ? height - fold.y - 1
        : height - (height - fold.y) - 1,
    width:
      width >= 2 * fold.x ? width - fold.x - 1 : width - (width - fold.x) - 1,
  };
};

const printDots = (dots: Dot[], height: number, width: number) => {
  const dotHash: { [k: string]: Dot } = dots.reduce(
    (acc, d) => ({ ...acc, [key(d)]: d }),
    {}
  );
  const rows = [];
  for (let y = 0; y <= height + 3; y++) {
    const row = [];
    for (let x = 0; x <= width + 3; x++) {
      let dot = { x, y };
      if (dotHash[key(dot)]) {
        row.push("#");
      } else {
        row.push("*");
      }
    }
    rows.push(row);
  }
  console.log(rows.map((r) => r.join("")).join("\n"));
};
const part1 = (fileName: string) => {
  const [dotsString, foldString] = fs
    .readFileSync(fileName)
    .toString()
    .split("\n\n");

  const dots = dotsString.split("\n").map((a) => {
    const [x, y] = a.split(",");
    return { x: parseInt(x), y: parseInt(y) };
  });
  const folds = foldString.split("\n").map((f) => {
    if (f.includes("y")) return { x: 0, y: parseInt(f.split("=")[1]) };
    return { y: 0, x: parseInt(f.split("=")[1]) };
  });
  // console.log(dots);
  // console.log(folds);
  const xMax = dots.reduce((max, d) => (d.x > max ? d.x : max), 0);
  const yMax = dots.reduce((max, d) => (d.y > max ? d.y : max), 0);
  // console.log(xMax);
  // console.log(yMax);
  let { transformedDots: folded } = fold(dots, folds[0], yMax, xMax);
  console.log(fileName, folded.length);
};

const part2 = (fileName: string) => {
  const [dotsString, foldString] = fs
    .readFileSync(fileName)
    .toString()
    .split("\n\n");

  const dots = dotsString.split("\n").map((a) => {
    const [x, y] = a.split(",");
    return { x: parseInt(x), y: parseInt(y) };
  });
  const folds = foldString.split("\n").map((f) => {
    if (f.includes("y")) return { x: 0, y: parseInt(f.split("=")[1]) };
    return { y: 0, x: parseInt(f.split("=")[1]) };
  });
  let initialWidth = dots.reduce((max, d) => (d.x > max ? d.x : max), 0);
  let initialHeight = dots.reduce((max, d) => (d.y > max ? d.y : max), 0);
  const finalDots = folds.reduce(
    (acc, f) => {
      const {
        transformedDots: dots1,
        height,
        width,
      } = fold(acc.dots, f, acc.height, acc.width);
      return { dots: dots1, height, width };
    },
    { dots, width: initialWidth, height: initialHeight }
  );
  console.log(fileName, "Part 2");
  printDots(finalDots.dots, finalDots.height, finalDots.width);
};

part1("test-input.txt");
part1("input.txt");

part2("test-input.txt");
part2("input.txt");
