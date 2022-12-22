package structural.decorator;

public class Email implements Notifier{

    @Override
    public boolean send(String message) {
        return false;
    }
}
