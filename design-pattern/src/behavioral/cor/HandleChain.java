package behavioral.cor;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class HandleChain implements Handler2 {
    private List<Handler2> handlers = new ArrayList<>();

    public HandleChain() {
    }

    public HandleChain(Handler2... handlers) {
        this.addHandle(handlers);
    }

    public void addHandle(Handler2... handlers) {
        this.handlers.addAll(Arrays.asList(handlers));
    }

    public void handle(Request request) {
        if (this.handlers.isEmpty()) {
            return;
        }
        for (Handler2 h :
                this.handlers) {
            if (this.check(request)) {
                h.handle(request);
            } else {
                System.out.println("stop");
            }
        }
    }

    private boolean check(Request req) {
        return req.code == 200;
    }
}
