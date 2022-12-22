package structural.decorator;

public class Example {
    public static void main(String[] args) {
            Notifier n = new LogDecorator(new ValidationDecorator(new Email()));
            n.send("hello world");
    }
}
