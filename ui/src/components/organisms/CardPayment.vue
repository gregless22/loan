<template>
  <div
    class="bg-gray-100 shadow rounded-lg overflow-hidden mb-3 mx-2 w-full px-4 py-1"
  >
    <div class="grid grid-cols-3">
      <p>Amount: ${{ Math.abs(payment.amount) }}</p>
      <p>Date Payed: {{ payment.settledAt.toUTCString().substring(4, 16) }}</p>
    </div>
  </div>
</template>

<script>
import { Payment } from "@/models/loan";
export default {
  name: "LoanCard",
  props: {
    payment: {
      type: Payment
    }
  },
  computed: {}
};
</script>
