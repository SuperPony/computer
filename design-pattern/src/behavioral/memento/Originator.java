package behavioral.memento;

public interface Originator {
    void saveState();
    void restore(Memento memento);
}
