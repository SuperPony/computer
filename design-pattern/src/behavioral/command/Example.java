package behavioral.command;

public class Example {

    public static void main(String[] args) {
        Editor editor = new Editor();
        Invoker invoker = new Invoke(
                new AppendCommand(editor),
                new CopyCommand(editor)
        );

        invoker.execute();
//        invoker.undo();
//        invoker.undo();
        invoker.redo();
    }
}
