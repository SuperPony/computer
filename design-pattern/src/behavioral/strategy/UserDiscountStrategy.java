package behavioral.strategy;

import java.math.BigDecimal;
import java.math.RoundingMode;

public class UserDiscountStrategy implements DiscountStrategy {

    /**
     * 计算满减后的价格
     * @param total 价格
     * @return 计算后的价格
     */
    @Override
    public BigDecimal computedDiscount(BigDecimal total) {
        BigDecimal magnification = total.divide(BigDecimal.valueOf(100)).setScale(0, RoundingMode.DOWN);
        return total.subtract(magnification.multiply(BigDecimal.valueOf(30)));
    }
}
