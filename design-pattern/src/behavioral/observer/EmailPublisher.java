package behavioral.observer;

import java.util.ArrayList;
import java.util.List;

public class EmailPublisher implements Observable {
    private List<Observer> subscribers = new ArrayList<>();

    private Event event;

    public EmailPublisher(Event event) {
        this.event = event;
    }

    @Override
    public void addSubscriber(Observer observer) {
        this.subscribers.add(observer);
    }

    @Override
    public void removeSubscriber(Observer observer) {
        this.subscribers.remove(observer);
    }

    public void notifySubscribers() {
        for (Observer s :
                this.subscribers) {
            s.execute(this.event);
        }
    }
}
