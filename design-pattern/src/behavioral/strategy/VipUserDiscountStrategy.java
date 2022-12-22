package behavioral.strategy;

import java.math.BigDecimal;

public class VipUserDiscountStrategy implements DiscountStrategy {
    /**
     * VIP 用户在满减基础上，再7折.
     * @param total 价格.
     * @return 计算后的 VIP 用户价格.
     */
    @Override
    public BigDecimal computedDiscount(BigDecimal total) {
        return total.multiply(BigDecimal.valueOf(0.7));
    }
}
