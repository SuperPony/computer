package structural.adapter;

public class SquarePeg {

    private  double width;

    public SquarePeg(double width) {
        this.width = width;
    }

    public double getSquare(){
        return Math.pow(this.width, 2);
    }

    public double getWidth() {
        return width;
    }

}
