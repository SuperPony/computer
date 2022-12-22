package behavioral.strategy;

import java.math.BigDecimal;

public class Order {
    private BigDecimal price;
    private DiscountStrategy strategy;

    public Order(BigDecimal price) {
        this.price = price;
    }

    public void setStrategy(DiscountStrategy strategy) {
        this.strategy = strategy;
    }

    public BigDecimal computedTotal() {
        if (this.strategy == null) {
            return this.price;
        }
        this.price = this.strategy.computedDiscount(this.price);

        return this.price;
    }
}
