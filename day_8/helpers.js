const shared = (a, b) => {
  const map = {};
  a.split("").forEach((s) => {
    map[s] = false;
  });
  b.split("").forEach((s) => {
    map[s] = map[s] === false;
  });
  return Object.entries(map)
    .filter(([_, v]) => v)
    .map(([k]) => k);
};

const ck = (key) => {
  const values = {
    a: 1,
    b: 10,
    c: 100,
    d: 1000,
    e: 10000,
    f: 100000,
    g: 1000000,
  };
  return key.split("").reduce((sum, curr) => (sum += values[curr]), 0);
};

module.exports = {
  shared,
  ck,
};
