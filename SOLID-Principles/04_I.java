/*
 * Created on 14 April 2024
 * 
 * @author Sai Sumanth
 */

/*
 * 📌 Core idea:
 * Split fat interfaces into smaller, more specific ones.
 */

interface BankService {
	void deposit();

	void withdraw();

	void checkBalance();

	void takeLoan();

	void payLoan();
}

class LoanAccountWithIssues implements BankService {
	public void deposit() {
		// Not supported
	}

	public void withdraw() {
		// Not supported
	}

	public void checkBalance() {
		System.out.println("Loan balance shown");
	}

	public void takeLoan() {
		System.out.println("Loan taken");
	}

	public void payLoan() {
		System.out.println("Loan paid");
	}
}

/*
 * ✅ After implementing Interface Seggregation Principle
 */

interface DepositService {
	void deposit();
}

interface WithdrawService {
	void withdraw();
}

interface LoanService {
	void takeLoan();

	void payLoan();
}

class SavingsAccount implements DepositService, WithdrawService {
	public void deposit() {
		System.out.println("Depositing in savings");
	}

	public void withdraw() {
		System.out.println("Withdrawing from savings");
	}
}

class LoanAccount implements LoanService {
	public void takeLoan() {
		System.out.println("Loan taken");
	}

	public void payLoan() {
		System.out.println("Loan paid");
	}
}
