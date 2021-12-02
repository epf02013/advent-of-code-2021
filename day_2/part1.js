const fs = require("fs");

const parseCommand = (command) => {
  const distance = parseInt(command.split(" ")[1]) || 0;
  if (command.includes("forward")) {
    return {
      depth: 0,
      horizontal: distance,
    };
  }
  if (command.includes("down")) {
    return {
      depth: distance,
      horizontal: 0,
    };
  }
  return {
    depth: -distance,
    horizontal: 0,
  };
};
const instructions = fs
  .readFileSync("input.txt")
  .toString()
  .split("\n")
  .map(parseCommand);

const endPosition = instructions.reduce(
  ({ depth, horizontal }, currCommand) => ({
    depth: depth + currCommand.depth,
    horizontal: horizontal + currCommand.horizontal,
  }),
  {
    depth: 0,
    horizontal: 0,
  }
);
console.log(endPosition);
console.log(endPosition.depth * endPosition.horizontal);
