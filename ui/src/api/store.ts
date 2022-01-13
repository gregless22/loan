import { Loan, Payment } from "@/models/loan";

async function getLoans(url: string) {
  let data = [];
  try {
    console.log("url", url);
    const response = await fetch(url, {
      credentials: "omit"
    });
    console.log(response);
    data = await response.json();
    console.log(data);
  } catch (err) {
    console.log("err", err);
  }
  // .then(response => console.log("response", response.json()))
  // .then(response => {
  //   //   console.log("ok", response.ok);
  //   //   console.log("response", response);
  //   response.json();
  // })
  // // .then(data => {
  // //   console.log("Success:", data);
  // // })
  // // .then(data => (loans = data.map((e: any) => new Loan(e))))
  // .then(data => console.log(data))
  // .catch(err => console.log("Errorr", err));

  // return loans;
  return data.map((e: any) => new Loan(e));
}

async function getPayments(url: string, options: any) {
  console.log("url payments", url);
  let payments: Payment[] = [];
  await fetch(url, options)
    .then(response => response.json())
    .then(data => (payments = data.data.map((e: any) => new Payment(e))))
    .catch(err => console.log(err));
  return payments;
}

class API {
  // all of the constants required
  private loansURL = "http://api.loan.server/loans/";
  private upURL =
    "https://api.up.com.au/api/v1/transactions?filter%5Btag%5D=LoanPayment";
  private upHeaders = {
    headers: new Headers({
      Authorization:
        "Bearer up:yeah:NhSasHMass5wnO7RnGM7FxM9DOjSpdrQzvnQxDNO5jM7GcTwVq4b7MwLq2Hely3iSGuIKrEFtGPa60v356Z597pGInWGyNrtSquaiUolCxZtng5s6pUO77dRr2qWH6xy"
    })
  };
  public loans: Loan[] = [];
  constructor() {
    // do nothing
  }

  async getLoans() {
    const loans = getLoans(this.loansURL);
    const payments = getPayments(this.upURL, this.upHeaders);
    const values = await Promise.all([payments, loans]);
    console.log("values", values);
    //combine the values
    values[0].forEach((e: Payment) => {
      values[1].forEach((f: Loan) => {
        // if between date and if the loan is the most
        if (
          e.settledAt.getTime() > f.startDate.getTime() &&
          e.settledAt.getTime() < f.endDate.getTime()
        ) {
          f.addPayment(e);
        }
      });
    });
    return values;
  }
}

export const store: API = new API();
