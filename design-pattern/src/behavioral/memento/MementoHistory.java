package behavioral.memento;

public interface MementoHistory {
    Memento pop();
    void push(Memento memento);
    boolean isEmpty();
}
