import { formatPhone } from '../filters/';

const formatPhoneDirective = {
  twoWay: true,
  bind: function (el, binding, vnode) {
    el.value = formatPhone(el.value);
  },
  update: function (el, binding, vnode) {
    el.value = formatPhone(el.value);
  },
  unbind: function (el, binding, vnode) {
  },
};

export {
  formatPhoneDirective,
}
