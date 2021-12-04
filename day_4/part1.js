const fs = require("fs");
const produce = require("immer").produce;

// type Board  {
//     items: {
//         [value: string]:{row: int;col: int; marked: true; value: number}
//     }
//     colCounts: { [value: number]: number }
//     rowCounts: { [value: number]: number }
// }

const parseBoard = (boardString) => {
  const noNewLines = boardString.replace(/\n/gm, " ");
  let items = noNewLines
    .split(" ")
    .map((a) => a.trim())
    .filter((a) => a)
    .reduce((acc, number, index) => {
      return {
        ...acc,
        [number]: {
          row: Math.floor(index / 5),
          col: index % 5,
          value: parseInt(number),
          marked: false,
        },
      };
    }, {});
  const board = {
    items,
    colCounts: {
      0: 0,
      1: 0,
      2: 0,
      3: 0,
      4: 0,
    },
    rowCounts: {
      0: 0,
      1: 0,
      2: 0,
      3: 0,
      4: 0,
    },
  };
  return board;
};

const markABoard = (board, pickedNumber) => {
  if (!board.items[pickedNumber]) {
    return board;
  }
  const markedBoard = produce(board, (draft) => {
    const item = draft.items[pickedNumber];
    item.marked = true;
    draft.colCounts[item.col] += 1;
    draft.rowCounts[item.row] += 1;
  });
  return markedBoard;
};

const boardIsInWinningState = (board) => {
  return (
    Math.max(
      ...Object.values(board.rowCounts),
      ...Object.values(board.colCounts)
    ) === 5
  );
};

const getWinningBoardAndNumber = (boards, pickedNumbers) => {
  let markedBoards = boards;
  for (let i = 0; i < pickedNumbers.length; i++) {
    markedBoards = markedBoards.map((b) => markABoard(b, pickedNumbers[i]));
    let winningBoard = markedBoards.find(boardIsInWinningState);
    if (winningBoard) {
      return { winningBoard, winningNumber: pickedNumbers[i] };
    }
  }
};
const calculateScore = (board, winningNumber) => {
  const unmarkedSum = Object.values(board.items)
    .filter((i) => !i.marked)
    .reduce((sum, curr) => sum + curr.value, 0);
  return unmarkedSum * winningNumber;
};
const main = () => {
  const parts = fs.readFileSync("input.txt").toString().split("\n\n");

  const pickedNumbers = parts[0].split(",");

  const boardStrings = parts.slice(1);
  const boards = boardStrings.map(parseBoard);
  const { winningBoard, winningNumber } = getWinningBoardAndNumber(
    boards,
    pickedNumbers
  );
  const score = calculateScore(winningBoard, winningNumber);
  console.log(score);
};
main();
