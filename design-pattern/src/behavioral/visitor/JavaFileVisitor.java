package behavioral.visitor;

import java.io.File;

public class JavaFileVisitor implements Visitor {
    @Override
    public void visitDir(File dir) {
        // ....
    }

    @Override
    public void visitFile(File file) {
        if (file.getName().endsWith(".java")) {
            System.out.println("Found java file: " + file);
        }
    }
}
