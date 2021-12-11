const fs = require("fs");
const { getNeighboursForMatrix, key } = require("./helper");
const { produce } = require("immer");

const MAX_ENERGY_BEFORE_FLASH = 9;
const increment = (pointData) => {
  return { ...pointData, value: pointData.value + 1 };
};
const incrementAllByOne = (matrix) => matrix.map((row) => row.map(increment));
const getPointsThatWillFlash = (matrix) => {
  return matrix.reduce((acc, row) => {
    const flashPointsInRow = row.filter(
      (i) => i.value > MAX_ENERGY_BEFORE_FLASH
    );
    return [...acc, ...flashPointsInRow];
  }, []);
};

const incrementPoints = (matrix, points) =>
  produce(matrix, (draft) => {
    points.forEach((n) => {
      draft[n.row][n.col] = increment(draft[n.row][n.col]);
    });
  });

const markPointsAsFlashed = (workingMatrix, points) =>
  produce(workingMatrix, (draft) => {
    points.forEach((p) => {
      const curr = draft[p.row][p.col];
      draft[p.row][p.col] = { ...curr, flashed: true };
    });
  });

const propagateFlash = (matrix, flashPoint) => {
  if (matrix[flashPoint.row][flashPoint.col].flashed)
    return { matrix, numberOfPointsFlashed: 0 };
  let workingMatrix = matrix;
  const visited = { [key(flashPoint)]: flashPoint };
  const queue = [flashPoint];
  while (queue.length > 0) {
    const currentPoint = queue.pop();
    const neighbours = getNeighboursForMatrix(workingMatrix)(currentPoint);
    workingMatrix = incrementPoints(workingMatrix, neighbours);
    const neighboursToVisit = neighbours.filter(
      (n) =>
        !visited[key(n)] &&
        !workingMatrix[n.row][n.col].flashed &&
        workingMatrix[n.row][n.col].value > MAX_ENERGY_BEFORE_FLASH
    );
    queue.push(...neighboursToVisit);
    neighboursToVisit.forEach((n) => (visited[key(n)] = n));
  }
  const pointsThatFlashed = Object.values(visited);
  const numberOfPointsFlashed = pointsThatFlashed.length;
  workingMatrix = markPointsAsFlashed(workingMatrix, pointsThatFlashed);
  return {
    matrix: workingMatrix,
    numberOfPointsFlashed: numberOfPointsFlashed,
  };
};

const resetFlash = (matrix) =>
  matrix.map((row) => row.map((p) => ({ ...p, flashed: false })));

const resetEnergyLevelsOfFlashedPoints = (matrix) =>
  matrix.map((row) =>
    row.map((p) => ({ ...p, value: p.value < 10 ? p.value : 0 }))
  );

const propagateFlashAndCountNumberOfFlashes = (acc, p) => {
  const { matrix, numberOfPointsFlashed } = propagateFlash(acc.matrix, p);
  return {
    matrix,
    numberOfPointsFlashed: numberOfPointsFlashed + acc.numberOfPointsFlashed,
  };
};

const takeStep = (initialMatrix) => {
  let matrixWithFlashedReset = resetFlash(initialMatrix);
  let matrixWithPowerIncreased = incrementAllByOne(matrixWithFlashedReset);
  const pointsThatWillFlash = getPointsThatWillFlash(matrixWithPowerIncreased);
  const { matrix: matrixAfterAllFlashes, numberOfPointsFlashed } =
    pointsThatWillFlash.reduce(propagateFlashAndCountNumberOfFlashes, {
      matrix: matrixWithPowerIncreased,
      numberOfPointsFlashed: 0,
    });
  return {
    matrix: resetEnergyLevelsOfFlashedPoints(matrixAfterAllFlashes),
    numberOfPointsFlashed,
  };
};

function part1(initialMatrix) {
  let count = 0;
  let currMatrix = initialMatrix;
  let pointsFlashed = 0;
  while (count < 100) {
    const { matrix, numberOfPointsFlashed } = takeStep(currMatrix);
    currMatrix = matrix;
    count += 1;
    pointsFlashed += numberOfPointsFlashed;
  }
  console.log("Part 1", pointsFlashed);
}

const main = () => {
  const initialMatrix = fs
    .readFileSync("input.txt")
    .toString()
    .split("\n")
    .map((row, rowI) =>
      row.split("").map((b, colI) => ({
        value: parseInt(b),
        row: rowI,
        col: colI,
        flashed: false,
      }))
    );
  part1(initialMatrix);
};

main();
