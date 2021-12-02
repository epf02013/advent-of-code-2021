const fs = require("fs");

const parseCommand = (command) => {
  const distance = parseInt(command.split(" ")[1]) || 0;
  if (command.includes("forward")) {
    return {
      aim: 0,
      horizontal: distance,
    };
  }
  if (command.includes("down")) {
    return {
      aim: distance,
      horizontal: 0,
    };
  }
  return {
    aim: -distance,
    horizontal: 0,
  };
};
const instructions = fs
  .readFileSync("input.txt")
  .toString()
  .split("\n")
  .map(parseCommand);

const endPosition = instructions.reduce(
  ({ depth, horizontal, aim }, currCommand) => ({
    depth: depth + currCommand.horizontal * aim,
    horizontal: horizontal + currCommand.horizontal,
    aim: aim + currCommand.aim,
  }),
  {
    depth: 0,
    horizontal: 0,
    aim: 0,
  }
);
console.log(endPosition);
console.log(endPosition.depth * endPosition.horizontal);
