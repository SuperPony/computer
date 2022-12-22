package behavioral.visitor;

import java.io.File;

public class ClassFileVisitor implements Visitor {
    @Override
    public void visitDir(File dir) {
        // ...
    }

    @Override
    public void visitFile(File file) {
        if (file.getName().endsWith(".class")) {
            System.out.println("Will clean class file: " + file);
        }
    }
}
