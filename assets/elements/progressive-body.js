import { eventsMixin } from '../lib/mixins'
export class ProgressiveBody extends HTMLBodyElement {
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
  connectedCallback() {
    this.listen(document, 'loading-start', () => (this.loading = true))
    this.listen(document, 'loading-stop', () => (this.loading = false))
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}
Object.assign(ProgressiveBody.prototype, eventsMixin)
