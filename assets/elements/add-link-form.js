import { eventsMixin } from '../lib/mixins'
import { createLink } from '../lib/api'
import { ProgressiveForm } from './progressive-form'
class AddLinkForm extends ProgressiveForm {
  constructor() {
    super()
    this.querySelector('input[type="text"]').remove()
    this.querySelector('label').remove()
    this.addLinkBtn = this.querySelector('#btn-link-add')
  }

  async handleCreateLink() {
    let url = prompt('Enter a website or feed URL')
    if (!url) return

    try {
      this.broadcast('loading-start')
      const link = await createLink(url)
      console.log('created link', link)
      this.broadcast('link-add-success', { link })
    } catch (err) {
      handleError(err)
    } finally {
      this.broadcast('loading-stop')
    }
  }

  connectedCallback() {
    this.listen(this.addLinkBtn, 'click', this.handleCreateLink)
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}
Object.assign(AddLinkForm.prototype, eventsMixin)
customElements.define('add-link-form', AddLinkForm, { extends: 'form' })
