package behavioral.visitor;

import java.io.File;
import java.util.Objects;

public class FileStructure {
    private File path;

    public FileStructure(File path) {
        this.path = path;
    }

    public void handle(Visitor visitor) {
        this.scan(this.path, visitor);
    }

    private void scan(File file, Visitor visitor) {
        if (file.isDirectory()) {
            visitor.visitDir(file);
            for (File sub :
                    Objects.requireNonNull(file.listFiles())) {
                this.scan(sub, visitor);
            }
        } else if (file.isFile()) {
            visitor.visitFile(file);
        }
    }
}
