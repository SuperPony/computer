package behavioral.template;

public abstract class Setting {

    public final String getSetting(String key) {
        String val = this.lookupCache(key);
        if (val == null) {
            val = this.readFromDatabase(key);
            this.puIntoCache(key, val);
        }

        return val;
    }

    private String readFromDatabase(String key) {
        // read data form database;
        return "val";
    }

    protected abstract String lookupCache(String key);

    protected abstract void puIntoCache(String key, String val);

}
