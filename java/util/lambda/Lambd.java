package util.bean.lombok;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class LombokBean {
    public final boolean dead = false;
    public String name;
    public int age;
}
