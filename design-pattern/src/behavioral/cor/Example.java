package behavioral.cor;

public class Example {
    public static void main(String[] args) {
        Handler startHandle = new StartHandle();
        Handler endHandle = new EndHandle();
        startHandle.setNext(endHandle);
        startHandle.handle(new Request(200, "hello world \n"));
    }
}
