package behavioral.memento;

import java.util.Stack;

public class History implements MementoHistory {
    private Stack<Memento> history = new Stack<>();

    public Memento get(int index) {
        return this.history.get(index);
    }

    @Override
    public Memento pop() {
        return this.history.pop();
    }

    @Override
    public void push(Memento memento) {
        if (this.history.size() >= 10) {
            this.pop();
        }
        this.history.push(memento);
    }

    @Override
    public boolean isEmpty() {
        return this.history.isEmpty();
    }
}
