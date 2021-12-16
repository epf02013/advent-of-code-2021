import { PriorityQueue } from "../priority-queue";
import { Queue } from "../queue";

export interface GraphNode<T> {
  value: T;
  name: string;
  edges: { start: string; end: string; weight: number }[];
}

export type Graph<T> = { [nodeName: string]: GraphNode<T> };

export function parseUnweightedGraphFromEdgeList<T>(
  graphString: string,
  directed: boolean = false
): Graph<T> {
  const edgeStrings = graphString.split("\n");
  let emptyObject: Graph<T> = {};
  const nodes = edgeStrings
    .flatMap((l) => l.split("-"))
    .reduce(
      (acc, curr) => ({ ...acc, [curr]: { edges: [], name: curr } }),
      emptyObject
    );
  edgeStrings.forEach((e) => {
    const [v1, v2] = e.split("-");
    nodes[v1].edges.push({ start: v1, end: v2, weight: 1 });
    if (directed) {
      nodes[v2].edges.push({ start: v2, end: v1, weight: 1 });
    }
  });
  // @ts-ignore
  return nodes;
}

export function parseWeightedGraphFromEdgeList<T>({
  graphString,
  weightSeparator = ",",
  nodeSeparator = "-",
  directed = false,
}: {
  graphString: string;
  weightSeparator?: string;
  nodeSeparator?: string;
  directed?: boolean;
}): Graph<T> {
  const edgeStrings = graphString.split("\n");
  let emptyObject: Graph<T> = {};
  const nodes = edgeStrings
    .flatMap((l) => l.split(weightSeparator)[0].split(nodeSeparator))
    .reduce(
      (acc, curr) => ({ ...acc, [curr]: { edges: [], name: curr } }),
      emptyObject
    );
  edgeStrings.forEach((e) => {
    const [vs, w] = e.split(weightSeparator);
    const [v1, v2] = vs.split(nodeSeparator);
    nodes[v1].edges.push({ start: v1, end: v2, weight: parseInt(w) });
    if (!directed) {
      nodes[v2].edges.push({ start: v2, end: v1, weight: parseInt(w) });
    }
  });
  // @ts-ignore
  return nodes;
}
export const key = ({ row, col }) => `${row},${col}`;

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

export function parseGraphFromMatrix(matrix: number[][]) {
  const graph = {};
  matrix.forEach((row, r) => {
    row.forEach((v, c) => {
      let point = { row: r, col: c };
      const neighbourPoints = getNeighbouringPoints(matrix, point);
      const node: GraphNode<{ row: number; col: number }> = {
        edges: [],
        name: key(point),
        value: point,
      };
      const edges = neighbourPoints.map((p: any) => ({
        start: key(point),
        end: key(p),
        weight: matrix[p.row][p.col],
      }));
      node.edges = edges;
      graph[key(point)] = node;
    });
  });
  return graph;
}

export function findShortestPath<T>(
  graph: Graph<T>,
  startNode: string,
  endNode: string
): { path: GraphNode<T>[]; weight: number } {
  const queue = new PriorityQueue<{
    node: GraphNode<T>;
    path: GraphNode<T>[];
  }>();
  queue.enqueue({ node: graph[startNode], path: [] }, 0);

  const visited: { [name: string]: number } = {};

  while (!queue.isEmpty()) {
    const { value, priority } = queue.dequeue();
    const { node, path } = value;

    if (node.name === endNode) {
      return { path: [...path, node], weight: priority };
    } else if (!visited[node.name] || visited[node.name] > priority) {
      visited[node.name] = priority;
      path.push(node);
      node.edges.forEach(
        (e: { start: string; end: string; weight: number }) => {
          queue.enqueue(
            { node: graph[e.end], path: [...path] },
            priority + e.weight
          );
        }
      );
    }
  }
  return { path: [], weight: 0 };
}

export function findAllConnectedNodes<T>(
  graph: Graph<T>,
  startNode: string
): GraphNode<T>[] {
  const queue = new Queue<string>();
  queue.enqueue(startNode);
  const visited: { [name: string]: GraphNode<T> } = {};
  while (!queue.isEmpty()) {
    const currNodeName = queue.dequeue();
    if (!visited[currNodeName]) {
      visited[currNodeName] = graph[currNodeName];
      graph[currNodeName].edges.forEach((e) => {
        queue.enqueue(e.end);
      });
    }
  }
  return Object.values(visited);
}
export function findConnectedComponents<T>(graph: Graph<T>): GraphNode<T>[][] {
  const connectedComponents: GraphNode<T>[][] = [];
  const foundNodes: { [key: string]: true } = {};
  Object.keys(graph).forEach((k) => {
    if (!foundNodes[k]) {
      const connectedComponent = findAllConnectedNodes(graph, graph[k].name);
      connectedComponents.push(connectedComponent);
      connectedComponent.forEach((n) => {
        foundNodes[n.name] = true;
      });
    }
  });
  return connectedComponents;
}
