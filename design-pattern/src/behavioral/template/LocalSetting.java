package behavioral.template;

import java.util.HashMap;
import java.util.Map;

public class LocalSetting extends Setting {

    private Map<String, String> cache = new HashMap<>();

    @Override
    protected String lookupCache(String key) {
        return this.cache.get(key);
    }

    @Override
    protected void puIntoCache(String key, String val) {
        this.cache.put(key, val);
    }
}
