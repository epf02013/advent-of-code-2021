import { Queue } from "./";

describe("queue", () => {
  describe("isEmpty", () => {
    it("should return false when not empty", () => {
      const queue = new Queue<number>();
      queue.enqueue(3);
      expect(queue.isEmpty()).toBe(false);
    });
    it("should return true when empty", () => {
      const queue = new Queue<number>();
      expect(queue.isEmpty()).toBe(true);
    });
  });
  describe("when multiple items have been enqueued", () => {
    it("should return them in fifo order", () => {
      const queue = new Queue<number>();
      queue.enqueue(1);
      queue.enqueue(2);
      const first = queue.dequeue();
      const second = queue.dequeue();
      queue.enqueue(3);
      expect(queue.isEmpty()).toBe(false);
      expect([first, second, queue.dequeue()]).toEqual([1, 2, 3]);
    });
  });
});
