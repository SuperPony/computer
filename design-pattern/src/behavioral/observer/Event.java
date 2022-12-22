package behavioral.observer;

public class Event {
    private final String eventType;
    private final String data;

    public Event(String eventType, String data) {
        this.eventType = eventType;
        this.data = data;
    }

    public String getEventType() {
        return eventType;
    }

    public String getData() {
        return data;
    }
}
