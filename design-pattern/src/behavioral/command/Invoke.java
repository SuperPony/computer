package behavioral.command;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Invoke implements Invoker {
    private List<Command> commands = new ArrayList<>();

    private CommandHistory history = new History();

    public Invoke(Command... commands) {
        this.add(commands);
    }

    public void add(Command... commands) {
        this.commands.addAll(Arrays.asList(commands));
    }

    @Override
    public void execute() {
        if (this.isEmpty()) {
            return;
        }
        for (Command command :
                this.commands) {
            command.execute();
            this.history.push(command);
        }
    }


    public boolean isEmpty() {
        return this.commands.isEmpty();
    }

    @Override
    public void undo() {
        if (!this.history.isEmpty()) {
            this.history.pop().execute();
        }
    }

    @Override
    public void redo() {
        for (int i = 0; i <= this.history.size(); i++) {
            this.undo();
        }
    }
}
