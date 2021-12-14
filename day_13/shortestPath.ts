import * as fs from "fs";
import {
  findShortestPath,
  parseUnweightedGraphFromEdgeList,
  parseWeightedGraphFromEdgeList,
} from "./graph";

const part1 = (fileName) => {
  const graphString = fs.readFileSync(fileName).toString();

  // const graph = parseUnweightedGraphFromEdgeList<string>(graphString)
  const graph = parseWeightedGraphFromEdgeList<string>({
    graphString,
    weightSeparator: ",",
    nodeSeparator: "-",
  });
  const { path: shortestPath, weight } = findShortestPath(
    graph,
    "start",
    "end"
  );
  console.log(`Part 1 (${fileName}) Shortest Path:`, shortestPath);
  console.log(`Part 1 (${fileName}) weight:`, weight);
};
const main = () => {
  part1("test-input.txt");
  // part1("input.txt");
};
