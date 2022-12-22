package creational.builder;


class CarBuilder implements Builder {
    private String color;
    private int seats;
    private CarType carType;

    @Override
    public void setCarType(CarType carType) {
        this.carType = carType;
    }

    @Override
    public void setColor(String color) {
        this.color = color;
    }

    @Override
    public void setSeats(int seats) {
        this.seats = seats;
    }

    public Car generateResult() {
        return new Car(carType, seats, color);
    }
}