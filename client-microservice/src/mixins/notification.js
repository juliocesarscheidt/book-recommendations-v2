import Vue from 'vue';

const NotificationMixin = {
  methods: {
    notifier(text, type) {
      const wordCount = text.toString().split(' ').length;
      // Duration can variable from 5 secs to 12 secs, according with the word count from text
      const duration = Math.min(Math.max(parseInt(wordCount, 10) * 1000, 5000), 12000);
      // Types: success, warn, error
      const classIcon = type === 'error' ? 'fas fa-exclamation-circle' : (type === 'warn' ? 'fas fa-info-circle' : 'fas fa-check-circle');
      this.$notify({
        group: 'notification',
        title: `${this.$t(`notification.title_${type}`)} <i class="${classIcon}" aria-hidden="true"></i>`,
        text,
        type,
        duration,
      });
    },
    notifyError(text, type = 'error') {
      this.notifier(text, type);
    },
    notifySuccess(text, type = 'success') {
      this.notifier(text, type);
    },
    notifyWarn(text, type = 'warn') {
      this.notifier(text, type);
    },
  },
};

Vue.mixin(NotificationMixin);
