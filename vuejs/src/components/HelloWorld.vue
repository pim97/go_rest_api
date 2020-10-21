<template>
  <div class="container">
    <div class="row">
      <div class="col-md-12">
        <label for="basic-url">Enter IMDB ID you want to look up</label>
        <div class="input-group mb-3">
          <div class="input-group-prepend">
            <span class="input-group-text" id="basic-addon3">IMDB ID tt</span>
          </div>
          <input
            type="text"
            v-model="amount.value"
            class="form-control"
            id="basic-url"
            aria-describedby="basic-addon3"
          />
        </div>

        <a href="#" @click="postMovie" class="btn btn-primary">Add to list</a>

        <table class="table">
          <thead>
            <tr>
              <th scope="col">IMDB Id</th>
              <th scope="col">Imdb score</th>
              <th scope="col">Imdb name</th>
              <th scope="col">Imdb plot</th>
              <th scope="col">Imdb year</th>
            </tr>
          </thead>

          <tbody v-for="(item, index) in amount.items" :key="index">
            <tr>
              <th scope="row">{{ item.imbd_id }}</th>
              <th scope="row">{{ item.imbd_score }}</th>
              <th scope="row">{{ item.name }}</th>
              <th scope="row">{{ item.plot }}</th>
              <th scope="row">{{ item.year }}</th>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "HelloWorld",
  data() {
    return {
      amount: {
        value: "abc",
        formatted: "",
        items: []
      }
    };
  },
  async created() {
    this.getMovies();
  },
  methods: {
    submit() {
      console.log(this.amount.value);
    },
    async getMovies() {
      this.amount.items = (
        await axios.get("http://localhost:1991/movies")
      ).data;
    },
    async postMovie() {
      const response = await axios.post("http://localhost:1991/movies", {
        name: "",
        year: 0,
        imbd_id: this.amount.value,
        imbd_score: 0
      });
      console.log(response.data);
      this.getMovies();
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
