<template>
  <main id="app">
    <Header/>
    <b-container>
      <transition name="slide" mode="out-in">
        <router-view></router-view>
      </transition>
      <notifications position="top right" group="notification" class="my-4 mx-4"/>
    </b-container>
  </main>
</template>

<script>
import { mapState } from 'vuex';
import Header from './components/common/Header';

export default {
  components: {
    Header,
  },
  computed: {
    ...mapState(['language']),
  },
  watch: {
    'language': {
      handler() {
        if (this.language && this.$i18n) {
          this.$i18n.locale = this.language;
        }
      },
    },
  },
}
</script>

<style>
#app {
  width: auto;
  height: auto;
  background-color: #fff !important;
  overflow: hidden;
}

@keyframes slide-in {
  from {
    transform: translateX(-25px);
    opacity: 0;
  }
  to {
    transform: translateX(0px);
    opacity: 1;
  }
}

@keyframes slide-out {
  from {
    transform: translateX(0px);
    opacity: 1;
  }
  to {
    transform: translateX(-25px);
    opacity: 0;
  }
}

.justify-content-60 {
  width: 60%;
  text-align: justify;
}

.slide-enter-active {
  animation: slide-in 0.25s ease;
}

.slide-leave-active {
  animation: slide-out 0.25s ease;
}

.full-height {
  height: 100%;
  width: auto;
}

.flex {
  display: flex;
}

.flex-justify-center {
  justify-content: center;
}

.flex-justify-start {
  justify-content: flex-start;
}

.flex-justify-end {
  justify-content: flex-end;
}

.flex-align-center {
  align-items: center;
}

.flex-align-start {
  align-items: flex-start;
}

.flex-align-end {
  align-items: flex-end;
}

.flex-space-between {
  justify-content: space-between;
}

.flex-space-around {
  justify-content: space-around;
}

.flex-start {
  justify-content: flex-start;
}

.flex-end {
  justify-content: flex-end;
}

.flex-row {
  flex-direction: row;
}

.flex-column {
  flex-direction: column;
}

@media only screen and (max-width: 991px) {
  .justify-content-60 {
    width: 100%;
  }
}
</style>

