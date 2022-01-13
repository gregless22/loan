<template>
  <header></header>
  <nav></nav>
  <main>
    <div class="container mx-auto p-4 flex justify-around flex-col md:flex-row">
      <g-card heading="Initial Amount:" :amount="intial"></g-card>
      <g-card heading="Current Amount:" :amount="amount"></g-card>
      <g-card heading="Total Paid:" :amount="paid"></g-card>
      <!-- <g-card heading="Date Paid By:"></g-card> -->
    </div>
    <div class="flex justify-center"></div>

    <div class="container mx-auto grid grid-cols-1">
      <div class="block text-xl text-gray-900 ml-4 pb-5">
        Loans:
      </div>
      <div v-for="loan in loans" :key="loan.id" class="">
        <g-card-loan :loan="loan"></g-card-loan>

        <div v-for="(payment, i) in loan.payments" :key="i">
          <div class="block text-xl text-gray-900 ml-4 pb-5" v-if="i === 0">
            Repayments:
          </div>
          <g-card-payment :payment="payment"></g-card-payment>
        </div>
      </div>
    </div>
  </main>
</template>

<script>
import { store } from "@/api/store";
export default {
  name: "Layout",
  data() {
    return {
      loans: []
    };
  },
  created() {
    store.getLoans().then(val => (this.loans = val[1]));
    // store.getLoans().then(data => (this.loans = data));
  },
  computed: {
    amount() {
      return this.loans[this.loans.length - 1]?.total();
    },
    intial() {
      return this.loans[this.loans.length - 1]?.amount;
    },
    paid() {
      return this.loans.reduce((acc, pv) => acc + pv.totalPayments(), 0);
    }
  }
};
</script>
