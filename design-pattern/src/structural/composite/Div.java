package structural.composite;

import java.util.List;

public class Div implements Dom{

    public List<Dom> children;

    public final String name;

    public Div(String name) {
        this.name = name;
    }


    @Override
    public String getName() {
        return name;
    }

    @Override
    public List<Dom> getChildren() {
        return children;
    }
}
