import { PriorityQueue } from "../priority-queue";

export interface GraphNode<T> {
  value: T;
  name: string;
  edges: { start: string; end: string; weight: number }[];
}

export type Graph<T> = { [nodeName: string]: GraphNode<T> };

export function parseUnweightedGraphFromEdgeList<T>(
  graphString: string
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
    nodes[v2].edges.push({ start: v2, end: v1, weight: 1 });
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

  // const visited: { [name: string]: boolean } = {};
  while (!queue.isEmpty()) {
    const { value, priority } = queue.dequeue();
    const { node, path } = value;
    if (node.name === endNode) {
      return { path: [...path, node], weight: priority };
    } else {
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
