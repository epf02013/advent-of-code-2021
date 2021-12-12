import * as fs from "fs";

type Node = {
  neighbours: string[];
  name: string;
};

type Graph = { [key: string]: Node };

const parseGraph = (graphString: string): Graph => {
  const edgeStrings = graphString.split("\n");
  let emptyObject: Graph = {};
  const nodes = edgeStrings.flatMap(l => l.split("-")).reduce((acc, curr) => ({...acc, [curr]: { neighbours: [], name: curr}}),emptyObject)
  edgeStrings.forEach(e => {
    const [v1,v2] = e.split("-")
    nodes[v1].neighbours.push(v2)
    nodes[v2].neighbours.push(v1)
  })
  return nodes
};

const part1 = (graph: Graph) => {
  const stack:{name: string, path: string[]}[] = [{name:'start', path: []}]
  const finalPaths: string[][] = []
  while (stack.length > 0) {
    const node = stack.pop()!
    const isALargeCave = node.name !== node.name.toLowerCase();
    const visited = node.path.includes(node.name);
    if(node.name === "end") {
      finalPaths.push(node.path)
    } else if(isALargeCave || !visited ) {
      node.path.push(node.name)
      graph[node.name].neighbours.forEach(neighbour => {
        stack.push({name: neighbour, path: [...node.path]})
      })
    }
  }
  console.log("Part 1", finalPaths.length)
}

const graphString = fs.readFileSync("input.txt").toString()
const graph = parseGraph(graphString)
part1(graph)

