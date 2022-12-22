package behavioral.iterator;

public interface Iterator<T> {
    int getIndex();

    T getNext();

    boolean hasNext();
}
