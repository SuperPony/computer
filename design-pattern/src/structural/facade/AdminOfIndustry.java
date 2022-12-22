package structural.facade;

public class AdminOfIndustry {

    public Company register(String name) {
        return new Company(name);
    }
}
