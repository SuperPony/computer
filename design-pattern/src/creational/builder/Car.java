package creational.builder;

class Car {
    private final CarType carType;
    private final int seats;
    private final String color;

    public Car(CarType carType, int seats, String color) {
        this.carType = carType;
        this.seats = seats;
        this.color = color;
    }

    public CarType getCarType() {
        return carType;
    }

    public int getSeats() {
        return seats;
    }

    public String getColor() {
        return color;
    }
}
