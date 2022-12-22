package behavioral.memento;

public class Example {
    public static void main(String[] args) {
        Car car = new Car("Benz", "red");
        System.out.println(car); // Car{band='Benz', color='red'}
        car.setBand("Audi");
        System.out.println(car); // Car{band='Audi', color='red'}
        car.undo();
        System.out.println(car); // Car{band='Benz', color='red'}
    }
}
