package structural.bridge;

public class Example {
    public static void main(String[] args) {
        Car car = new Car(
                new Band("Benz"),
                new Color("#fff")
        );
        System.out.println(car.getBand()); // Benz
    }
}
