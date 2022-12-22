package structural.decorator;

public class LogDecorator extends BaseDecorator {
    public LogDecorator(Notifier notifier) {
        super(notifier);
    }

    @Override
    public boolean send(String message) {
        boolean res = notifier.send(message);
        // write log...
        return res;
    }
}
