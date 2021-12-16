import * as fs from "fs";
import { findShortestPath, key, parseGraphFromMatrix } from "./graph";

const part1 = (fileName) => {
  const graphString = fs.readFileSync(fileName).toString();
  const matrix = graphString
    .split("\n")
    .map((a) => a.split("").map((b) => parseInt(b)));

  const graph = parseGraphFromMatrix(matrix);
  const { path: shortestPath, weight } = findShortestPath(
    graph,
    key({ row: 0, col: 0 }),
    key({ row: matrix.length - 1, col: matrix[0].length - 1 })
  );
  console.log(`Part 1 (${fileName}) weight:`, weight);
};

function createRepeatedMatrix(matrix: number[][]) {
  const realMatrix = [];
  for (let r = 0; r < matrix.length * 5; r++) {
    const row = [];
    for (let col = 0; col < matrix[0].length * 5; col++) {
      let rMod = Math.floor(r / matrix.length) % 5 > 0 ? 1 : 0;
      let cMod = Math.floor(col / matrix.length) % 5 > 0 ? 1 : 0;
      let colIndex = col % matrix[0].length;
      let rowIndex = r % matrix.length;
      let lastValue = matrix[rowIndex][colIndex];
      if (rMod !== 0 && rMod > cMod) {
        lastValue = realMatrix[r - matrix.length][colIndex];
      } else if (cMod !== 0) {
        lastValue = row[col - matrix[0].length];
      }
      let value = lastValue + Math.max(rMod, cMod);
      let maxedValue = value > 9 ? 1 : value;
      row.push(maxedValue);
    }
    realMatrix.push(row);
  }
  return realMatrix;
}

const part2 = (fileName) => {
  const graphString = fs.readFileSync(fileName).toString();
  const matrix = graphString
    .split("\n")
    .map((a) => a.split("").map((b) => parseInt(b)));

  const repeatedMatrix = createRepeatedMatrix(matrix);

  const graph = parseGraphFromMatrix(repeatedMatrix);
  const { weight } = findShortestPath(
    graph,
    key({ row: 0, col: 0 }),
    key({ row: repeatedMatrix.length - 1, col: repeatedMatrix[0].length - 1 })
  );
  console.log(`Part 2 (${fileName}) weight:`, weight);
};
const main = () => {
  part1("input.txt");
  part2("input.txt");
};

main();
