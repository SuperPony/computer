package creational.builder;

public class Example {
    public static void main(String[] args) {
        CarBuilder b = new CarBuilder();
        Director d = new Director();
        d.constructCityCar(b);
        Car cityCar = b.generateResult();
        d.constructSuv(b);
        Car suv = b.generateResult();
        System.out.println(cityCar.getColor()); // red
        System.out.println(suv.getColor()); // green
    }
}





