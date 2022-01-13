export class Loan {
  public id: number;
  public from: string;
  public to: string;
  public startDate: Date;
  public endDate: Date;
  public rate: number;
  public amount: number;
  private payments: Payment[];
  constructor({
    id,
    from,
    to,
    startDate,
    endDate,
    rate,
    amount,
    payments = []
  }: any) {
    this.id = id;
    this.from = from;
    this.to = to;
    this.startDate = new Date(startDate);
    this.endDate =
      endDate === "0001-01-01T00:00:00Z" ? new Date() : new Date(endDate);
    this.rate = rate;
    this.amount = amount;
    this.payments = payments;
  }

  // functions required
  // this will go through the loan and see how much is remaining. It will totalise to the end date.
  total() {
    return (
      this.amount * Math.pow(1 + this.rate / 100 / 365, this.getTimeInDays()) +
      this.payments.reduce((acc, pv) => acc + pv.total(this), 0)
    );
  }

  // addPayment will add a payment to the array
  addPayment(p: Payment) {
    this.payments.push(p);
  }
  totalPayments() {
    return this.payments.reduce((acc, pv) => {
      return acc + Math.abs(pv.amount);
    }, 0);
  }

  getTimeInDays() {
    return Math.round(
      (this.endDate.getTime() - this.startDate.getTime()) / 86400000
    );
  }
  getTimeInMonths() {
    return (
      (this.endDate.getDate() > this.startDate.getDate() ? 0 : -1) +
      this.endDate.getMonth() -
      this.startDate.getMonth() +
      12 * Math.abs(this.endDate.getFullYear() - this.startDate.getFullYear())
    );
  }
  getActive() {
    return this.endDate.toDateString() === new Date().toDateString();
  }
}

export class Payment {
  public id: string;
  public amount: number;
  public settledAt: Date;
  constructor(c: any) {
    this.id = c.id;
    this.amount = c.attributes.amount.value;
    this.settledAt = new Date(c.attributes.settledAt);
  }

  total(loan: Loan) {
    return (
      this.amount *
      Math.pow(1 + loan.rate / 100 / 365, this.getTimeInDays(loan))
    );
  }

  getTimeInDays(loan: Loan) {
    return (loan.endDate.getTime() - this.settledAt.getTime()) / 86400000;
  }
}
