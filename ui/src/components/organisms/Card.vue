<template>
  <div
    class="max-w-sm bg-white shadow rounded-lg overflow-hidden mb-3 mx-2 md:w-2/3 w-full p-6"
  >
    <div class="flex flex-row justify-around items-center">
      <p class="text-xl">{{ heading }}</p>
      <p class="text-l text-gray-600">${{ commaAmount }}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: "Card",
  props: {
    heading: {
      type: String
    },
    amount: {
      default: ""
    }
  },
  computed: {
    commaAmount() {
      return this.amount.toLocaleString("en");
    }
  }
};
</script>
