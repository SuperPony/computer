package structural.composite;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class ElementDom implements Dom {
    public List<Dom> children = new ArrayList<>();
    public final String name;

    public ElementDom(String name, Dom... domes) {
        this.name = name;
        this.addAll(domes);
    }

    public ElementDom(String name) {
        this.name = name;
    }

    public boolean add(Dom dom) {
        return this.children.add(dom);
    }

    public boolean addAll(Dom... domes) {
        return this.children.addAll(Arrays.asList(domes));
    }

    public boolean remove(Dom dom) {
        return this.children.remove(dom);
    }

    @Override
    public String getName() {
        return name;
    }

    public List<Dom> getChildren() {
        return children;
    }
}
