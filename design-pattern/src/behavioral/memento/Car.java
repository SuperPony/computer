package behavioral.memento;

public class Car implements Originator {
    private String band;
    private String color;

    private MementoHistory mementoHistory = new History();

    public Car(String band, String color) {
        this.band = band;
        this.color = color;
    }

    public void setBand(String band) {
        this.saveState();
        this.band = band;
    }

    public void setColor(String color) {
        this.saveState();
        this.color = color;
    }

    @Override
    public void saveState() {
        Memento memento = new CarMemento(this.band, this.color);
        this.mementoHistory.push(memento);

    }

    public void undo() {
        Memento carMemento = this.mementoHistory.pop().getState();
        this.restore(carMemento);
    }

    @Override
    public void restore(Memento memento) {
        CarMemento carMemento = (CarMemento) memento;
        this.band = carMemento.band;
        this.color = carMemento.color;
    }

    @Override
    public String toString() {
        return "Car{" +
                "band='" + band + '\'' +
                ", color='" + color + '\'' +
                '}';
    }

    class CarMemento implements Memento {
        private final String band;
        private final String color;

        public CarMemento(String band, String color) {
            this.band = band;
            this.color = color;
        }

        @Override
        public Memento getState() {
            return this;
        }
    }
}
