package behavioral.command;

interface Command {
    void execute();

    void undo();
}

interface Invoker extends Command {
    void redo();
}

interface CommandHistory {
    void push(Command command);

    Command pop();

    boolean isEmpty();

    int size();
}

