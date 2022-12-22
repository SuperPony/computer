package behavioral.command;

import java.util.Iterator;
import java.util.Stack;

public class History implements CommandHistory {
    private Stack<Command> history = new Stack<>();

    public boolean isEmpty() {
        return this.history.isEmpty();
    }

    @Override
    public int size() {
        return this.history.size();
    }

    @Override
    public Command pop() {
        return this.history.pop();
    }

    @Override
    public void push(Command command) {
        this.history.push(command);
    }

    public Iterator<Command> iterator() {
        return this.history.stream().iterator();
    }
}
