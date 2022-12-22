package structural.composite;

public class Example {

    public static void main(String[] args) {
        Dom tree = new ElementDom(
                "root",
                new Div("div"),
                new Div("div")
        );
    }
}
