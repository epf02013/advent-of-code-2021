const getSumBits = (numbers) => {
  const sumBits = numbers.reduce((acc, curr) => {
    return acc.map((x, i) => x + curr[i]);
  }, Array(numbers[0].length).fill(0));
  return sumBits;
};

const getCommonBits = (numbers, tieWinner) => {
  const sumBits = getSumBits(numbers);
  const rowCount = numbers.length;
  const commonBits = sumBits.map((columnSum) => {
    if (columnSum / rowCount > 0.5) return 1;
    if (columnSum / rowCount < 0.5) return 0;
    return tieWinner;
  });
  return commonBits;
};
const getUncommonBits = (numbers, tieWinner) => {
  const sumBits = getSumBits(numbers);
  const rowCount = numbers.length;
  const unCommonBits = sumBits.map((columnSum) => {
    if (columnSum / rowCount > 0.5) return 0;
    if (columnSum / rowCount < 0.5) return 1;
    return tieWinner;
  });
  return unCommonBits;
};

const parseAnInt = (bitArray) => parseInt(bitArray.join(""), 2);

module.exports = {
  parseAnInt,
  getUncommonBits,
  getCommonBits,
};
