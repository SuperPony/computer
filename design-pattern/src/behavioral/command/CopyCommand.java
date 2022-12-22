package behavioral.command;

public class CopyCommand implements Command {
    protected Editor receiver;

    public CopyCommand(Editor receiver) {
        this.receiver = receiver;
    }

    @Override
    public void execute() {
        this.receiver.copy();
    }

    @Override
    public void undo() {
        // Do something....
    }
}
