package behavioral.cor;

public class EndHandle implements Handler {
    private Handler next;

    @Override
    public void setNext(Handler handle) {
        this.next = handle;
    }

    @Override
    public void handle(Request req) {
        System.out.println("EndHandle");
        // Do something...
        if (this.next != null) {
            this.next.handle(req);
        }
    }
}
