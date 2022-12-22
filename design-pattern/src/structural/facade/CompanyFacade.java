package structural.facade;

public class CompanyFacade {
    protected AdminOfIndustry admin = new AdminOfIndustry();
    protected Bank bank = new Bank();
    protected Taxation taxation = new Taxation();

    protected Company company;

    public Company openCompany(String name) {
        Company c = this.admin.register(name);
        String bankAccount = this.bank.openAccount(c.getCompanyId());
        c.setBankAccount(bankAccount);
        String taxCode = this.taxation.applyTaxCode(c.getCompanyId());
        c.setTaxCode(taxCode);

        return c;
    }
}
