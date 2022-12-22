package behavioral.command;

public class AppendCommand implements Command {
    protected Editor receiver;

    public AppendCommand(Editor receiver) {
        this.receiver = receiver;
    }

    @Override
    public void execute() {
        this.receiver.append("hello world");
    }

    @Override
    public void undo() {
        // Do something....
    }

}
