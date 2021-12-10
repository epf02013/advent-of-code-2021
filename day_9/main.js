const fs = require("fs");

const key = ({ row, col }) => `${row},${col}`;
const getNeighbouringPoints = (matrix, { row, col }) => {
  const rowMin = Math.max(row - 1, 0);
  const rowMax = Math.min(matrix.length - 1, row + 1);
  const colMin = Math.max(col - 1, 0);
  const colMax = Math.min(matrix[0].length - 1, col + 1);

  const points = [
    { row: rowMin, col },
    { row: rowMax, col },
    { row, col: colMin },
    { row, col: colMax },
  ];
  const uniquePoints = Object.values(
    points.reduce((acc, curr) => ({ ...acc, [key(curr)]: curr }), {})
  ).filter((a) => a.row != row || a.col != col);
  return uniquePoints;
};
const getValue = (matrix) => (p) => matrix[p.row][p.col];
const isLowPoint = (matrix, point) => {
  const neighbouringValues = getNeighbouringPoints(matrix, point).map(
    getValue(matrix)
  );
  return !neighbouringValues.some((v) => v <= matrix[point.row][point.col]);
};

function getLowPointInfo(data) {
  const lowPointData = data
    .flatMap((rowData, row) =>
      rowData.map((number, col) => [
        number,
        isLowPoint(data, { row, col }),
        { row, col },
      ])
    )
    .filter((a) => a[1]);
  return lowPointData;
}
const findBasin = (matrix, point) => {
  const visited = {};
  const queue = [point];
  visited[key(point)] = point;
  while (queue.length > 0) {
    const curr = queue.pop();
    const neighbours = getNeighbouringPoints(matrix, curr);
    const higherUnvisitedNeighbours = neighbours.filter((n) => {
      return !visited[key(n)] && getValue(matrix)(n) != 9;
    });
    higherUnvisitedNeighbours.forEach((p) => {
      visited[key(p)] = p;
    });
    queue.push(...higherUnvisitedNeighbours);
  }
  return Object.values(visited);
};

const part1 = (data) => {
  const sum = (acc, curr) => acc + curr;
  const lowPointInfo = getLowPointInfo(data);

  let riskValues = lowPointInfo.map((a) => a[0] + 1);

  console.log("Part 1:", riskValues.reduce(sum));
};

const part2 = (data) => {
  const lowPoints = getLowPointInfo(data).map((a) => a[2]);

  const allBasins = lowPoints
    .map((p) => findBasin(data, p))
    .sort((a, b) => b.length - a.length);
  let scorePart2 = allBasins
    .slice(0, 3)
    .reduce((mult, b) => mult * b.length, 1);
  console.log("score part 2 ", scorePart2);
};

const main = () => {
  const data = fs
    .readFileSync("input.txt")
    .toString()
    .split("\n")
    .map((a) => a.split("").map((b) => parseInt(b)));
  part1(data);
  part2(data);
};

main();
