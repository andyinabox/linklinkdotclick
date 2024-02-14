import { eventsMixin } from '../lib/mixins'
import { createLink } from '../lib/api'
import { ProgressiveForm } from './progressive-form'
class AddLinkForm extends ProgressiveForm {
  constructor() {
    super()
    this.inputs.url.remove()
    this.querySelector('label').remove()
  }

  async createLink() {
    let url = prompt('Enter a website or feed URL')
    if (!url) return

    try {
      this.broadcast('loading-start')
      const link = await createLink(url)
      this.broadcast('link-add-success', { link })
    } catch (err) {
      handleError(err)
    } finally {
      this.broadcast('loading-stop')
    }
  }

  connectedCallback() {
    this.listen(this.buttons['btn-link-add'], 'click', this.createLink)
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}
Object.assign(AddLinkForm.prototype, eventsMixin)
customElements.define('add-link-form', AddLinkForm, { extends: 'form' })
