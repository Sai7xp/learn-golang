public class LiskovExample {
    public static void main(String[] args) {
        BankAccount acc = new SavingsAccount(); // This is valid in Java
        acc.deposit(100);
        acc.deposit(20);

        BankAccount fdAcc = new FDAccount();
        fdAcc.deposit(100);
        fdAcc.withdraw(100); // 🔴 Exception, normal withdraw not allowed in FD.
        // so subclass failed to replace the superclass object

        /*
         * So we have to create new class WithdrawableAccount and move withdraw method
         * to this new class, and BankAccount contains only deposit method
         * 
         * class WithdrawableAccount extends BankAccount{}
         * 
         * Now, class SavingsAccount extends WithdrawableAccount{}
         * class FDAccount extends BankAccount{}
         */
    }
}

class BankAccount {
    double balance;

    void deposit(double amount) {
        this.balance += amount;
        System.out.println("Deposit success ✅");
    }

    void withdraw(double amount) {
        this.balance -= amount;
        System.out.println("Withdrawl success ✅");
    }
}

class SavingsAccount extends BankAccount {
}

class FDAccount extends BankAccount {
}