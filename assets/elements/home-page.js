import { ProgressiveBody } from './progressive-body'
class HomePage extends ProgressiveBody {
  set loading(v) {
    if (v) {
      this.classList.add('loading')
    } else {
      this.classList.remove('loading')
    }
  }
  get loading() {
    this.classList.contains('loading')
  }

  set editing(v) {
    if (v) {
      this.classList.add('editing')
    } else {
      this.classList.remove('editing')
    }
  }
  get editing() {
    this.classList.contains('editing')
  }

  connectedCallback() {
    this.listen(this, 'edit-mode-start', () => (this.editing = true))
    this.listen(this, 'edit-mode-stop', () => (this.editing = false))
    this.listen(this, 'loading-start', () => (this.loading = true))
    this.listen(this, 'loading-stop', () => (this.loading = false))
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}
customElements.define('home-page', HomePage, { extends: 'body' })
