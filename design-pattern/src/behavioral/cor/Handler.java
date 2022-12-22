package behavioral.cor;

interface Handler {
    void setNext(Handler handle);

    void handle(Request req);
}

interface Handler2 {
    void handle(Request req);
    void addHandle(Handler2... handles);
}
