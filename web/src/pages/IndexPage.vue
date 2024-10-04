<template>
  <div class="row">
    <div class="col" style="min-width: 15rem; max-width: 15rem;">
      <div class="q-pa-md">
        <q-card class="my-card">
          <q-card-section>
            <div class="text-h6">Region: </div>

            <div class="q-pa-md">
              <div class="q-gutter-sm">
                <q-radio v-model="region" val="Krasnodar" label="Krasnodar" color="blue-10" />
                <q-radio v-model="region" val="Nizhny Novgorod" label="Nizhny Novgorod" color="blue-10" />
                <q-radio v-model="region" val="Minsk" label="Minsk" color="blue-10" />
                <q-radio v-model="region" val="Italy" label="Italy" color="blue-10" />
                <q-radio v-model="region" val="Istanbul" label="Istanbul" color="blue-10" />
                <q-radio v-model="region" val="Yerevan" label="Yerevan" color="blue-10" />
              </div>

            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>
    <div class="col">
      <div class="q-pa-md">
        <div class="q-gutter-md row items-start">
          <q-date v-model="days" mask="DD-MM-YYYY" range color="blue-10" />
        </div>
        <div class="q-pa-md q-gutter-sm">
          <q-btn class="brand-elem-not-accent" v-on:click="onReset(days)" label="Reset" />
          <q-btn class="brand-elem-accent" v-on:click="onRequest(days)" label="Request" />
        </div>
      </div>
    </div>
  </div>

  <div class="q-pa-md">
    <q-table title="Prices" :rows="rows" :columns="columns" row-key="title" :loading="loading" />
  </div>
</template>

<script>
import { ref } from "vue";

const columns = [
  {
    name: "title",
    required: true,
    field: "title",
    label: "Hotel",
    align: "left",
    field: (row) => row.title,
    format: (val) => `${val}`,
    sortable: true,
  },
  { name: "current_price", label: "Price", field: "current_price", align: "left" },
  { name: "yandex_price", label: "Competitor's price (Yandex)", field: "yandex_price", align: "left" },
  { name: "yandex_discount", label: "Competitor's discount (Yandex)", field: "yandex_discount", align: "left" },
  { name: "icon", label: "", field: "icon", align: "right" },
  { name: "recommended_price", label: "Recommended price", field: "recommended_price", align: "left" },
];

function getIcon(current_price, recommended_price) {
  const curPrice = +current_price
  const recPrice = +recommended_price


  const pr = curPrice < recPrice ? recPrice / curPrice : curPrice / recPrice;
  if (pr >= 0.05) {
    return curPrice > recPrice ? "↓" : "↑"
  } else {
    return "="
  }
}

async function getHotels(region, from, to) {
  try {
    let response = await fetch("http://localhost:8088/hotels", {
      method: "POST",
      body: JSON.stringify({
        region: region,
        checkin: from,
        checkout: to,
      }),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    });

    let hotelsJSON = await response.json()
    let hotelsRows = hotelsJSON.map(hotel => {
      hotel.icon = getIcon(hotel.current_price, hotel.recommended_price)
      return hotel
    }

    )
    return hotelsRows
  } catch (e) {
    console.error(`Error: ${e}`);
  }

}


export default {
  setup() {
    return {
      columns,
      rows: ref([]),
      loading: ref(false),
      region: ref('Krasnodar'),
      days: ref([{ from: new Date(), to: new Date() }]),
      async onRequest(arg) {
        console.log("Request", arg);
        this.loading = true
        let rows = await getHotels(this.region, this.days.from, this.days.to);
        this.rows = ref(rows)
        this.loading = false
        console.log("===== ROWS", this.rows)
      },
      onReset(arg) {
        arg.from = '';
        arg.to = '';
      },
    };
  },
};
</script>
