package behavioral.visitor;

import java.io.File;

public class Example {
    public static void main(String[] args) {
        FileStructure fileStructure =new FileStructure(new File("."));
        fileStructure.handle(new JavaFileVisitor());
    }
}
