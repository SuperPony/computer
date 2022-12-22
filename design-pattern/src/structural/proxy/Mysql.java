package structural.proxy;

public class Mysql implements Store {
    protected String user;
    protected String pwd;

    public Mysql(String user, String pwd) {
        this.user = user;
        this.pwd = pwd;
    }

    public boolean connect(){
        System.out.println("Mysql connect");
        return  true;
    }
    // Other method...
    public boolean close(){
        System.out.println("Mysql close");
        return  false;
    }
}
