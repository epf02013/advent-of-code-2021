const fs = require("fs")

const items = fs.readFileSync("input.txt").toString().split("\n")

const {count} =
    items.reduce(
        (acc,curr)=>({prev: curr, count: acc.count + 1}),
        {prev: -1, count: 0}
    )

console.log(count)

