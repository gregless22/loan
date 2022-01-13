import { createApp } from "vue";
import App from "./App.vue";
import "./assets/tailwind.css";

// global component registration
import Card from "@/components/organisms/Card.vue";
import Graph from "@/components/organisms/Graph.vue";
import CardLoan from "@/components/organisms/CardLoan.vue";
import CardPayment from "@/components/organisms/CardPayment.vue";

const app = createApp(App);

app.component("g-card", Card);
app.component("g-graph", Graph);
app.component("g-card-loan", CardLoan);
app.component("g-card-payment", CardPayment);

app.mount("#app");
