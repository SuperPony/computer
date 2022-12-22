package behavioral.observer;

// Publisher
public interface Observable {
    void addSubscriber(Observer subscriber);

    void removeSubscriber(Observer subscriber);
}
