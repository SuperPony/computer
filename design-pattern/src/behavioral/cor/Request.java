package behavioral.cor;

public class Request {
    protected final int code;
    protected final String body;

    public Request(int code, String body) {
        this.code = code;
        this.body = body;
    }
}
