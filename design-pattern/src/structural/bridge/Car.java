package structural.bridge;

public class Car {

    private Band band;
    private Color color;

    public Car(Band band, Color color) {
        this.band = band;
        this.color = color;
    }

    public String getBand() {
        return band.getName();
    }

    public String  getColor() {
        return color.getRgb();
    }
}
