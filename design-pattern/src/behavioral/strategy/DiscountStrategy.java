package behavioral.strategy;

import java.math.BigDecimal;

public interface DiscountStrategy {
    BigDecimal computedDiscount(BigDecimal total);
}
