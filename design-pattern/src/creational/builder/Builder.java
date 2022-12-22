package creational.builder;

interface Builder {
    void setCarType(CarType carType);

    void setColor(String color);

    void setSeats(int seats);
}