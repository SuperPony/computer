package creational.builder;

class Director {

    public void constructCityCar(Builder builder){
        builder.setCarType(CarType.CITY_CAR);
        builder.setColor("red");
        builder.setSeats(4);
    }

    public void constructSuv(Builder builder){
        builder.setSeats(4);
        builder.setColor("green");
        builder.setCarType(CarType.SUV);
    }
}
