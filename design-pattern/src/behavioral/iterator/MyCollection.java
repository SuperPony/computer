package behavioral.iterator;

import java.util.Arrays;

public class MyCollection<T> implements ConcreteCollection<T> {
    private T[] array;

    public MyCollection(T... array) {
        this.array = Arrays.copyOfRange(array, 0, array.length);
    }

    @Override
    public Iterator<T> iterator() {
        return new MyIterator();
    }

    class MyIterator implements Iterator<T> {
        int index;

        @Override
        public int getIndex() {
            return this.index;
        }

        @Override
        public T getNext() {
            T res = array[index];
            this.index++;
            return res;
        }

        @Override
        public boolean hasNext() {
            return this.index < MyCollection.this.array.length;
        }
    }
}
