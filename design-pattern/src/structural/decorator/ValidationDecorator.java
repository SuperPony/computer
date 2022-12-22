package structural.decorator;

public class ValidationDecorator extends BaseDecorator{


    public ValidationDecorator(Notifier notifier) {
        super(notifier);
    }

    @Override
    public boolean send(String message) {
        // validation params...

        return notifier.send(message);
    }
}
