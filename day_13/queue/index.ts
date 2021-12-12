interface MyNode<T> {
  value: T;
  next?: MyNode<T>;
}
export class Queue<T> {
  head?: MyNode<T>;
  tail?: MyNode<T>;
  enqueue(value: T) {
    if (!this.head) {
      this.head = {
        value,
        next: null,
      };
    }
    if (!this.tail) {
      this.tail = this.head;
    } else {
      const newTail = {
        value,
        next: null,
      };
      this.tail.next = newTail;
      this.tail = newTail;
    }
  }
  dequeue() {
    const head = this.head;
    this.head = head.next;
    return head.value;
  }
  isEmpty() {
    return !this.head;
  }
}
