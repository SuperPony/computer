package behavioral.cor;

public class StartHandle implements Handler {
    private Handler next;

    @Override
    public void setNext(Handler next) {
        this.next = next;
    }

    @Override
    public void handle(Request req) {
        System.out.println("StartHandle");
        // Do something...
        if (this.next != null) {
            this.next.handle(req);
        }
    }
}
