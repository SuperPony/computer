package behavioral.strategy;

import java.math.BigDecimal;

public class Example {
    public static void main(String[] args) {
        Order order = new Order(BigDecimal.valueOf(150));
        System.out.println(order.computedTotal()); // 150
        order.setStrategy(new UserDiscountStrategy());
        System.out.println(order.computedTotal()); // 120
        order.setStrategy(new VipUserDiscountStrategy());
        System.out.println(order.computedTotal()); // 84
    }
}
