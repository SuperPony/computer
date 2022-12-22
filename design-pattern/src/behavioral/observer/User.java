package behavioral.observer;

public class User implements Observer{
    @Override
    public void execute(Event event) {
        System.out.println(event.getEventType());
    }
}
