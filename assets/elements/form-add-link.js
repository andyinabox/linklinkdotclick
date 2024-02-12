import { eventsMixin } from '../lib/mixins'
import { addLink } from '../lib/api'
import { ProgressiveForm } from '../lib/progressive-form'
class FormAddLink extends ProgressiveForm {
  constructor() {
    super()
    this.querySelector('input[type="text"]').remove()
  }

  async handleCreateLink() {
    let url = prompt('Enter a website or feed URL')
    if (!url) return
    this.broadcast('link-add-request', { url })
  }

  connectedCallback() {
    const submit = this.querySelector('input[type="submit"]')
    this.listen(submit, 'click', this.handleCreateLink)
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}
Object.assign(FormAddLink.prototype, eventsMixin)
customElements.define('linky-add-link', FormAddLink, { extends: 'form' })
