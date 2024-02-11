import { slotsMixin, listenMixin } from './mixins'
export class Component extends HTMLElement {
  constructor() {
    super()
    this.registerSlotsMixin()
    this.registerListenMixin()
  }
  // automatically render when data is set
  set data(d) {
    this._data = d
    this.render()
  }

  // create a clone of data when retrieving
  // (we only want this changed when using setter)
  get data() {
    return { ...this._data }
  }
  set loading(bool) {
    if (bool) {
      this.classList.add('loading')
    } else {
      this.classList.remove('loading')
    }
  }
  get loading() {
    return this.classList.contains('loading')
  }
  render() {}
  disconnectedCallback() {
    this.unlisten()
  }
}

Object.assign(Component.prototype, slotsMixin, listenMixin)
