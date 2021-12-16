export class PriorityQueue<T> {
  private isMinQueue: boolean;
  constructor(isMinQueue = true) {
    this.isMinQueue = isMinQueue;
  }
  isEmpty(): Boolean {
    return this.size === 0;
  }

  items: { value: T; priority: number }[] = new Array(2);
  size: number = 0;

  enqueue(value: T, priority: number): void {
    this.items[this.size] = { value, priority };
    let index = this.size;
    this.size += 1;
    while (index > 0) {
      const parentIndex = Math.floor((index - 1) / 2);
      if (this.getIndexOfItemWithPriority(parentIndex, index) === parentIndex) {
        break;
      }
      const temp = this.items[index];
      this.items[index] = this.items[parentIndex];
      this.items[parentIndex] = temp;
      index = parentIndex;
    }
  }
  getIndexOfItemWithPriority(i1: number, i2: number): number {
    if (i1 > this.size - 1) return i2;
    if (i2 > this.size - 1) return i1;
    if (this.isMinQueue) {
      const indexOfPriorityItem =
        this.items[i1].priority < this.items[i2].priority ? i1 : i2;
      return indexOfPriorityItem;
    }
    const indexOfPriorityItem =
      this.items[i1].priority > this.items[i2].priority ? i1 : i2;
    return indexOfPriorityItem;
  }
  dequeue(): { value: T; priority: number } {
    const returnValue = this.items[0];
    this.size -= 1;
    this.items[0] = this.items[this.size];
    this.items[this.size] = null;
    let index = 0;
    while (index < this.size) {
      const priorityChildIndex = this.getIndexOfItemWithPriority(
        index * 2 + 1,
        index * 2 + 2
      );
      if (
        this.getIndexOfItemWithPriority(priorityChildIndex, index) === index
      ) {
        break;
      }
      const temp = this.items[index];
      this.items[index] = this.items[priorityChildIndex];
      this.items[priorityChildIndex] = temp;
      index = priorityChildIndex;
    }
    return { value: returnValue.value, priority: returnValue.priority };
  }
}
