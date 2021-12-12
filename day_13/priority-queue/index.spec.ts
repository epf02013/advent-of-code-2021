import { PriorityQueue } from "./index";

describe("PriorityQueue", () => {
  describe("isEmpty", () => {
    it("should return true when empty", () => {
      const queue = new PriorityQueue();
      expect(queue.isEmpty()).toBe(true);
    });
    it("should return false when not empty", () => {
      const queue = new PriorityQueue<number>();
      queue.enqueue(3, 2);
      expect(queue.isEmpty()).toBe(false);
    });
  });
  describe("dequeue", () => {
    describe("when multiple things have been enqueued", () => {
      describe("when it is a min queue", () => {
        it("should return the ones with the lowest priority", () => {
          const queue = new PriorityQueue<string>(true);
          queue.enqueue("Number 2", 2);
          queue.enqueue("3", 3);
          queue.enqueue("4", 4);
          queue.enqueue("Number one", 1);
          queue.enqueue("3", 3);
          queue.enqueue("3", 3);
          expect([
            queue.dequeue(),
            queue.dequeue(),
            queue.dequeue(),
            queue.dequeue(),
            queue.dequeue(),
            queue.dequeue(),
          ]).toEqual(["Number one", "Number 2", "3", "3", "3", "4"]);
        });
      });
      describe("when it is a max queue", () => {
        it("should return the ones with the highest priority", () => {
          const queue = new PriorityQueue<string>(false);
          queue.enqueue("2", 2);
          queue.enqueue("3", 3);
          queue.enqueue("4", 4);
          queue.enqueue("1", 1);
          queue.enqueue("3", 3);
          queue.enqueue("3", 3);
          expect([
            queue.dequeue(),
            queue.dequeue(),
            queue.dequeue(),
            queue.dequeue(),
            queue.dequeue(),
            queue.dequeue(),
          ]).toEqual(["4", "3", "3", "3", "2", "1"]);
        });
      });
    });
  });
});
