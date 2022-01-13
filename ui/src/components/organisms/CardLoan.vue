<template>
  <div :class="active">
    <div class="grid grid-cols-4 items-center">
      <p class="text-l text-orange-600 flex flex-row">
        {{ loan.from }}
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
          class="w-6 mx-3"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M17 8l4 4m0 0l-4 4m4-4H3"
          />
        </svg>
        {{ loan.to }}
      </p>
      <!-- <p class="text-xl">{{ loan }}</p> -->
      <p class="text-l text-gray-600">Amount: ${{ commaAmount }}</p>
      <p class="text-l text-gray-600">Rate: {{ loan.rate }} %</p>
      <div>
        <p class="text-l text-gray-600">
          Date Started: {{ loan.startDate.toUTCString().substring(4, 16) }}
        </p>
        <p class="text-l text-gray-600">
          Date Ended: {{ loan.endDate.toUTCString().substring(4, 16) }}
        </p>
        <p>
          Total Days: <b>{{ loan.getTimeInDays() }}</b>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import { Loan } from "@/models/loan";
export default {
  name: "LoanCard",
  props: {
    loan: {
      type: Loan
    }
  },
  computed: {
    commaAmount() {
      return this.loan.amount.toLocaleString("en");
    },
    description() {
      return `${this.loan.from}`;
    },
    active() {
      return this.loan.getActive()
        ? "bg-green-100 shadow rounded-lg overflow-hidden mb-3 mx-2 w-full p-6"
        : "bg-red-100 shadow rounded-lg overflow-hidden mb-3 mx-2 w-full p-6";
    }
  }
};
</script>
