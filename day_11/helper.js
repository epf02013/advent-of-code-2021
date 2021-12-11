const key = ({ row, col }) => `${row},${col}`;

const neighbours =
  (matrix) =>
  ({ row, col }) => {
    const rowMin = Math.max(row - 1, 0);
    const rowMax = Math.min(matrix.length - 1, row + 1);
    const colMin = Math.max(col - 1, 0);
    const colMax = Math.min(matrix[0].length - 1, col + 1);

    const points = [
      { row: rowMin, col },
      { row: rowMax, col },
      { row, col: colMin },
      { row, col: colMax },
      { row: rowMax, col: colMax },
      { row: rowMin, col: colMin },
      { row: rowMin, col: colMax },
      { row: rowMax, col: colMin },
    ];
    const uniquePoints = Object.values(
      points.reduce((acc, curr) => ({ ...acc, [key(curr)]: curr }), {})
    ).filter((a) => a.row != row || a.col != col);
    return uniquePoints.map((p) => matrix[p.row][p.col]);
  };
module.exports = {
  getNeighboursForMatrix: neighbours,
  key,
};
