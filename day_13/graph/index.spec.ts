import {
  findShortestPath,
  Graph,
  parseUnweightedGraphFromEdgeList,
  parseWeightedGraphFromEdgeList,
} from "./index";

describe("parseUnweightedGraphFromEdgeList", () => {
  it("it should create a graph", () => {
    const graphString =
      "start-A\n" +
      "start-b\n" +
      "A-c\n" +
      "A-b\n" +
      "b-d\n" +
      "A-end\n" +
      "b-end";
    const graph: Graph<string> = parseUnweightedGraphFromEdgeList(graphString);
    expect(graph["A"].edges.some((e) => e.end === "c")).toBe(true);
    expect(graph["A"].edges.some((e) => e.end === "b")).toBe(true);
    expect(graph["A"].edges.some((e) => e.end === "end")).toBe(true);
  });
});

describe("parseWeightedGraphFromEdgeList", () => {
  describe("when it is undirected", () => {
    it("it should create edges going both ways", () => {
      const graphString =
        "start-A,3\n" +
        "start-b,3\n" +
        "A-end,4\n" +
        "A-b,2\n" +
        "b-d,4\n" +
        "A-end,5\n" +
        "b-end,5";
      const graph: Graph<string> = parseWeightedGraphFromEdgeList({
        graphString: graphString,
      });
      expect(
        graph["b"].edges.some((e) => e.end === "A" && e.weight === 2)
      ).toBe(true);
      expect(
        graph["A"].edges.some((e) => e.end === "b" && e.weight === 2)
      ).toBe(true);
      expect(
        graph["A"].edges.some((e) => e.end === "end" && e.weight === 5)
      ).toBe(true);
    });
  });
  describe("when it is directed", () => {
    it("it should create edges going both ways", () => {
      const graphString =
        "start-A,3\n" +
        "start-b,3\n" +
        "A-end,4\n" +
        "A-b,2\n" +
        "b-d,4\n" +
        "A-end,5\n" +
        "b-end,5";
      const graph: Graph<string> = parseWeightedGraphFromEdgeList({
        graphString: graphString,
        directed: true,
      });
      expect(graph["b"].edges.some((e) => e.end === "A")).toBe(false);
      expect(
        graph["A"].edges.some((e) => e.end === "b" && e.weight === 2)
      ).toBe(true);
    });
  });
});

describe("shortestPath", () => {
  it("should return the shortest path and weight", () => {
    const graphString = "start-A,3\n" + "start-b,2\n" + "A-end,4\n" + "b-end,9";
    const graph: Graph<string> = parseWeightedGraphFromEdgeList({
      graphString: graphString,
      directed: false,
    });
    const { path: shortestPath, weight } = findShortestPath<string>(
      graph,
      "start",
      "end"
    );
    expect(shortestPath.map((n) => n.name)).toEqual(["start", "A", "end"]);
    expect(weight).toEqual(7);
  });
});

describe("allPathsWithSpecialConsideration", () => {
  it("should return the shortest path and weight", () => {
    const graphString = "start-A,3\n" + "start-b,2\n" + "A-end,4\n" + "b-end,9";
    const graph: Graph<string> = parseWeightedGraphFromEdgeList({
      graphString: graphString,
      directed: false,
    });
    const { path: shortestPath, weight } = findShortestPath<string>(
      graph,
      "start",
      "end"
    );
    expect(shortestPath.map((n) => n.name)).toEqual(["start", "A", "end"]);
    expect(weight).toEqual(7);
  });
});
