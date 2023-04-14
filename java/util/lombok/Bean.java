package util.bean.lombok;

import lombok.AccessLevel;
import lombok.Setter;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStream;


class ToSetter {
    private Long id;
    private String msg;

    public ToSetter() {
    }

    public void setId(Long id) {
        this.id = id;
    }

    public void setMsg(String msg) {
        this.msg = msg;
    }
}

class Getter {
    private Long id;
    private String msg;

    public Getter() {
    }

    public Long getId() {
        return id;
    }

    public String getMsg() {
        return msg;
    }
}

/**
 * 为类提供一个属性名 log 的 log4j 日志对象，以及一个默认构造方法
 */
class Log4j {
    private static final Logger log = LoggerFactory.getLogger(Log4j.class);

    public Log4j() {
    }
}


/**
 * {@code @Data} = @Getter + @Setter + @ToString + @EqualsAndHashCode + @RequiredArgsConstructor . <br/>
 * 默认情况下 Getter/Setter 默认是 public，可以显示设置 @Setter/@Getter 注解对 class 或属性设置访问级别
 */
class Data {
    private Long id;

    // 默认情况下 Getter/Setter 默认是 public，可以显示设置 @Setter/@Getter 注解对 class 或属性设置访问级别
    @Setter(AccessLevel.PROTECTED)
    private String msg;

    public Data() {
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getMsg() {
        return msg;
    }

//    public void setMsg(String msg) {
//        this.msg = msg;
//    }

    // equals method

    // canEqual method

    // hashCode method

    // toString method
}

/**
 * {@code  @AllArgsConstructor} 生成一个包含所有属性的构造函数，生成不再生成默认的构造函数
 */
class AllArgsConstructor {
    private final Long id;
    private final String msg;

    public AllArgsConstructor(Long id, String msg) {
        this.id = id;
        this.msg = msg;
    }
}

class NoArgsConstructor {
    private Long id;
    private String msg;

    public NoArgsConstructor() {
    }
}


/**
 * {@code @RequiredArgsConstructor} 生成一个包含了所有 {@code @NonNull} 注解还有 {@code final} 修饰的属性的构造函数
 */
class RequiredArgsConstructor {
    private final Long id;
    @lombok.NonNull
    private String name;
    private int age;

    public RequiredArgsConstructor(Long id, String name) {
        if (name == null) {
            throw new NullPointerException("name is marked non-null but is null");
        } else {
            this.id = id;
            this.name = name;
        }
    }

    public void ignore() {
        this.name = "a";
    }
}

/**
 * {@code @Cleanup} 用在变量前，用于自动关闭变量代表的资源，默认调用资源的 close() 方法，通过声明注解的 value 可以修改关闭方法
 */
class Cleanup {
    public void example() throws IOException {
        @lombok.Cleanup InputStream in = new FileInputStream("./demo.txt");
    }
}


/**
 * {@code @NonNull} 用于形参时，当该形参被赋值时，会先检测传入的值是否为 null，是则抛出 {@link NullPointerException} 异常<br/>
 * 用于属性时，会自动为其 Setter 和构造对应的形参进行 null 判断
 */
class NonNull {
    private final long id;
    // @lombok.NonNull
    private String name;

    public NonNull(String name, long id) {
        if (name == null) {
            throw new NullPointerException("name is marked non-null but is null");
        } else {
            this.name = name;
            this.id = id;
        }
    }

    public void setName(String name) {
        if (name == null) {
            throw new NullPointerException("name is marked non-null but is null");
        } else {
            this.name = name;
        }
    }

    public void say(@lombok.NonNull String word) {
        System.out.println(word);
    }

    public void sayToLombok(String word) {
        if (word == null) {
            throw new NullPointerException("word is marked non-null but is null");
        } else {
            System.out.println(word);
        }
    }
}


class ToBuilder {

}