import CARD from "./card/index.vue";
import DATA from "./data/index.vue";
import OILCARD from "./oilcard/index.vue";
import PHONEBILL from "./phonebill/index.vue";
import POWERBILL from "./powerbill/index.vue";
import RIGHTS from "./rights/index.vue";

export const list = [
  {
    key: "phonebill",
    title: "话费",
    component: PHONEBILL
  },
  {
    key: "rights",
    title: "权益",
    component: RIGHTS
  },
  {
    key: "card",
    title: "卡券",
    component: CARD
  },
  {
    key: "oilcard",
    title: "油卡",
    component: OILCARD
  },
  {
    key: "data",
    title: "流量",
    component: DATA
  },
  {
    key: "powerbill",
    title: "电费",
    component: POWERBILL
  }
];
