import { createApp } from "vue";
import { createPinia } from "pinia";
import "./index.css";
import router from "./utils/router";
import App from "./App.vue";

const pinia = createPinia();
const app = createApp(App);

app.use(pinia);
app.use(router);

app.mount("#app");
